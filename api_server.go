package sova

import (
	"context"
	"errors"
)

var ErrUnsuccessResponse = errors.New("unsuccess response")

// ServerOnline ...
//
// /api/server/online/
func (api *API) ServerOnline(ctx context.Context) (ServerOnline, error) {
	resp, err := getAndUnmarshal[struct {
		Success bool         `json:"success"`
		Data    ServerOnline `json:"data"`
	}](
		api, ctx,
		"server/online",
		true,
	)
	if err != nil {
		return ServerOnline{}, err
	}
	if !resp.Success {
		return ServerOnline{}, ErrUnsuccessResponse
	}
	return resp.Data, nil
}

// ServerOnlineDirect ...
func (api *API) ServerOnlineDirect(ctx context.Context) int {
	resp, err := api.ServerOnline(ctx)
	if err != nil {
		return 0
	}
	return resp.Online
}
