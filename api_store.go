package sova

import (
	"context"
	"errors"
	"fmt"
)

var ErrNoSuccessFieldFound = errors.New("no success field found")

// StoreVerifyPlayer checks if player with this nickname exists.
//
// /api/store/verify-player/{nickname}
func (api *API) StoreVerifyPlayer(ctx context.Context, nickname string) (*StoreVerifyPlayerResponse, error) {
	resp, err := getAndUnmarshal[struct {
		//Data []interface{} `json:"data"`
		Success *StoreVerifyPlayerResponse `json:"success,omitempty"`
	}](
		api, ctx,
		f("store/verify-player/%s", nickname),
		true,
	)
	if err != nil {
		return nil, err
	}
	if resp.Success == nil {
		return nil, ErrNoSuccessFieldFound
	}
	return resp.Success, nil
}

// StoreVerifyPlayerDirect ...
func (api *API) StoreVerifyPlayerDirect(ctx context.Context, nickname string) bool {
	res, err := api.StoreVerifyPlayer(ctx, nickname)
	if err != nil {
		return false
	}
	return *res
}

type ErrNoAvailableRanks struct {
	Player string
}

func (e ErrNoAvailableRanks) Error() string {
	if e.Player == "" {
		return "no available ranks"
	}
	return fmt.Sprintf("no available ranks for %s", e.Player)
}

// StoreRanks returns list of ranks for a specific player (nickname).
//
// /api/store/ranks/{nickname}
func (api *API) StoreRanks(ctx context.Context, nickname string) (StoreRanksResponse, error) {
	resp, err := getAndUnmarshal[StoreRanksResponse](
		api, ctx,
		f("store/ranks/%s", nickname),
	)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNoAvailableRanks{Player: nickname}
	}
	return resp, nil
}

type ErrNoItemsAvailable struct {
	Player string
}

func (e ErrNoItemsAvailable) Error() string {
	if e.Player == "" {
		return "no items available"
	}
	return fmt.Sprintf("no items available for %s", e.Player)
}

// StoreItems ...
//
// /api/store/items/{nickname}
func (api *API) StoreItems(ctx context.Context, nickname string) (StoreItemsResponse, error) {
	resp, err := getAndUnmarshal[StoreItemsResponse](
		api, ctx,
		f("store/items/%s", nickname),
	)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNoItemsAvailable{Player: nickname}
	}
	return resp, nil
}
