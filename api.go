// Package sova is a wrapper for sovamc.com API.
// You can create API instance with sova.New(), just like:
//
//	api := sova.New()
//
// Or using the Config:
//
//	api := sova.Config{...}.New()
//
// Then you can call methods easily:
//
//	resp, err := api.PlayerSearch(ctx, "джавид")
//	if err != nil {
//	}
//	_ = resp
//
// You can see more examples in "/example" directory
package sova

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/k4ties/sovapi/internal"
)

const RootURL = "https://api.sovamc.com/api/"

func NewAPI() *API {
	var conf Config
	return conf.New()
}

type API struct {
	conf Config
}

type ErrBadStatus struct {
	StatusCode int
	Status     string
	Body       []byte
}

func (e ErrBadStatus) Error() string {
	if e.Status == "" {
		return fmt.Sprintf("bad status code %d", e.StatusCode)
	}
	return fmt.Sprintf("bad status code %d: %s", e.StatusCode, e.Status)
}

func (api *API) get(parent context.Context, endpoint string) ([]byte, error) {
	path, err := url.JoinPath(RootURL, endpoint)
	if err != nil {
		return nil, fmt.Errorf("join path: %w", err)
	}
	ctx, cancel := context.WithTimeout(parent, api.conf.RequestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	res, err := api.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer res.Body.Close() //nolint:errcheck

	data, err := io.ReadAll(io.LimitReader(res.Body, int64(api.conf.MaxBodySize)))
	if err != nil {
		//if res.StatusCode != http.StatusOK {
		//	// it is expected to not have body
		//	return nil, ErrBadStatus{StatusCode: res.StatusCode, Status: res.Status}
		//}
		return nil, fmt.Errorf("read response body: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		// try to parse existing body as ResponseError and return the matching
		// error
		if err := unmarshalAndMatchResponseError(data); err != nil {
			return nil, err
		}
		// if failed return default ErrBadStatus
		return nil, ErrBadStatus{StatusCode: res.StatusCode, Status: res.Status, Body: data}
	}
	return data, nil
}

func (api *API) doRequest(req *http.Request) (*http.Response, error) {
	resp, err := api.conf.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp == nil { //nilaway
		return nil, errors.New("no response")
	}
	return resp, nil
}

var ErrResponseDoesntMatchSignature = errors.New("response doesnt match deafult signature (missing Data field)")

func unmarshalResponseAsData[T any](data []byte) (T, bool, error) {
	var zero T
	type resp struct {
		Data    *T    `json:"data,omitempty"`
		Success *bool `json:"success,omitempty"`
	}
	res, respErr, err := unmarshalResponseAsStruct[resp](data)
	if err != nil {
		return zero, respErr, err
	}
	if res.Data == nil {
		return zero, false, ErrResponseDoesntMatchSignature
	}
	return *res.Data, true, nil
}

var ErrUnhandledResponseType = errors.New("unhandled (invalid) response type")

func unmarshalResponseAsStruct[T any](data []byte) (res T, respErr bool, err error) {
	var resp T
	if err := json.Unmarshal(data, &resp); err == nil {
		return resp, false, nil
	}
	return res, false, ErrUnhandledResponseType
}

type ErrUnmarshalResponse struct {
	Parent error
}

func (e ErrUnmarshalResponse) Error() string {
	return fmt.Sprintf("unmarshal response: %v", e.Parent)
}

func getAndUnmarshal[T any](api *API, ctx context.Context, endpoint string, asStruct ...bool) (zero T, err error) {
	data, err := api.get(ctx, endpoint)
	if err != nil {
		return zero, err
	}
	var (
		resp    T
		respErr bool
	)
	if internal.HasTrueOption(asStruct) {
		resp, respErr, err = unmarshalResponseAsStruct[T](data)
	} else {
		resp, respErr, err = unmarshalResponseAsData[T](data)
	}
	if err == nil {
		// success
		return resp, nil
	}
	if respErr {
		// return error directly if it is extracted from ResponseError
		return zero, err
	}
	return zero, ErrUnmarshalResponse{Parent: err}
}

// f is shorter alias for fmt.Sprintf.
func f(f string, a ...any) string {
	return fmt.Sprintf(f, a...)
}

// TODO account api .................
