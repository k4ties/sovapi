package sova

import (
	"context"
	"fmt"
)

// StoreVerifyPlayer checks if player with this nickname exists.
//
// /api/store/verify-player/{nickname}
func (api *API) StoreVerifyPlayer(ctx context.Context, nickname string) (*StoreVerifyPlayerResponse, error) {
	return getAndUnmarshalf[StoreVerifyPlayerResponse](api, ctx, "store/verify-player/%s", nickname)
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
func (api *API) StoreRanks(ctx context.Context, nickname string) (*StoreRanksResponse, error) {
	resp, err := getAndUnmarshalf[StoreRanksResponse](api, ctx, "store/ranks/%s", nickname)
	if err != nil {
		return nil, err
	}
	if len(resp.Data) == 0 {
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
func (api *API) StoreItems(ctx context.Context, nickname string) (*StoreItemsResponse, error) {
	resp, err := getAndUnmarshalf[StoreItemsResponse](api, ctx, "store/items/%s", nickname)
	if err != nil {
		return nil, err
	}
	if len(resp.Data) == 0 {
		return nil, ErrNoItemsAvailable{Player: nickname}
	}
	return resp, nil
}
