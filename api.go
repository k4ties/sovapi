// Package sova is a wrapper for sovamc.com API.
// You can create API instance with sova.New(), just like:
//
//	api := sova.New()
//
// Or using the config APIConfig:
//
//	api := sova.APIConfig{...}.New()
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
	"strings"

	"github.com/k4ties/sovapi/internal/errmatch"
)

const RootURL = "https://api.sovamc.com/api/"

func NewAPI() *API {
	var conf Config
	return conf.New()
}

type API struct {
	conf Config
}

func (api *API) get(parent context.Context, path string) ([]byte, error) {
	path = RootURL + strings.TrimPrefix(path, "/") //maybe url.JoinPath would be usable here

	ctx, cancel := context.WithTimeout(parent, api.conf.RequestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	res, err := doRequest(api.conf.Client, req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer res.Body.Close() //nolint:errcheck

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	return data, nil
}

func doRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp == nil { //nilaway
		return nil, errors.New("no response")
	}
	return resp, nil
}

func unmarshalResponseError(data []byte) (error, bool) {
	var resp ResponseError
	if err := json.Unmarshal(data, &resp); err != nil || resp.Message == "" {
		return nil, false
	}
	if err := errmatch.Match(errMatches, resp.Message); err != nil {
		return err, true
	}
	// ResponseError implements error, so we can simply return it
	return resp, true
}

// the following methods should be API private methods, but they can't only
// because of Go doesn't support generics that way

func unmarshalResponse[T any](data []byte) (result *T, respErr bool, err error) {
	if err, ok := unmarshalResponseError(data); ok {
		return nil, true, err
	}
	var resp *T
	if err := json.Unmarshal(data, &resp); err == nil {
		return resp, false, nil
	}
	return nil, false, errors.New("invalid (unhandled) response type")
}

func getAndUnmarshalf[T any](api *API, ctx context.Context, f string, a ...any) (*T, error) {
	return getAndUnmarshal[T](api, ctx, fmt.Sprintf(f, a...))
}

type ErrUnmarshalResponse struct {
	Parent error
}

func (e ErrUnmarshalResponse) Error() string {
	return fmt.Sprintf("unmarshal response: %v", e.Parent)
}

func (e ErrUnmarshalResponse) Unwrap() error {
	return e.Parent
}

func getAndUnmarshal[T any](api *API, ctx context.Context, path string) (result *T, err error) {
	data, err := api.get(ctx, path)
	if err != nil {
		return nil, err
	}
	resp, respErr, err := unmarshalResponse[T](data)
	if err == nil {
		// success
		return resp, nil
	}
	if respErr {
		// return error directly if it is extracted from ResponseError
		return nil, err
	}
	return nil, ErrUnmarshalResponse{Parent: err}
}

// TODO account api .................
