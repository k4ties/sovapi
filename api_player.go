package sova

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"
)

type ErrCannotFindPlayer struct {
	AsID   bool
	Player string
}

func (e ErrCannotFindPlayer) Error() string {
	if e.Player == "" {
		return "cannot find player"
	}
	return fmt.Sprintf("cannot find player: %s", e.Player)
}

// Player ...
//
// /api/player/{id}
func (api *API) Player(ctx context.Context, id int) (*PlayerResponse, error) {
	resp, err := getAndUnmarshalf[PlayerResponse](api, ctx, "player/%d", id)
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, ErrCannotFindPlayer{AsID: true, Player: strconv.Itoa(id)}
	}
	return resp, nil
}

var ErrNicknameMustBeTwoChars = errors.New("nickname must be at least 2 characters long")

// PlayerSearch searches for a player with a specific query.
//
// /api/player/search/{query}
func (api *API) PlayerSearch(ctx context.Context, query string) (*PlayerSearchResponse, error) {
	if utf8.RuneCountInString(query) < 2 {
		return nil, ErrNicknameMustBeTwoChars
	}
	resp, err := getAndUnmarshalf[PlayerSearchResponse](api, ctx, "player/search/%s", query)
	if err != nil {
		return nil, err
	}
	if len(resp.Data) == 0 {
		return nil, ErrCannotFindPlayer{Player: query}
	}
	return resp, nil
}
