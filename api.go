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
//	 _ = resp
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
	"unicode/utf8"
)

const RootURL = "https://api.sovamc.com/api/"

func NewAPI() *API {
	var conf APIConfig
	return conf.New()
}

type API struct {
	conf APIConfig
}

// Player ...
//
// /api/player/{id}
func (api *API) Player(ctx context.Context, id int) (PlayerResponse, error) {
	return getAndUnmarshalf[PlayerResponse](api, ctx, "player/%d", id)
}

// PlayerSearch searches for a player with a specific query.
//
// /api/player/search/{query}
func (api *API) PlayerSearch(ctx context.Context, query string) (PlayerSearchResponse, error) {
	if utf8.RuneCount(([]byte)(query)) < 2 {
		return PlayerSearchResponse{}, ErrNicknameMustBeTwoChars
	}
	resp, err := getAndUnmarshalf[PlayerSearchResponse](api, ctx, "player/search/%s", query)
	if err != nil {
		return PlayerSearchResponse{}, err
	}
	if len(resp.Data) == 0 {
		return PlayerSearchResponse{}, ErrCannotFindPlayer
	}
	return resp, nil
}

// PracticeMode returns a list of all available practice modes, including
// ranked.
//
// /api/practice/mode/
func (api *API) PracticeMode(ctx context.Context) (PracticeModeResponse, error) {
	return getAndUnmarshal[PracticeModeResponse](api, ctx, "practice/mode")
}

// PracticeModeRanked returns a list of all available ranked practice modes.
//
// /api/practice/mode/ranked/
func (api *API) PracticeModeRanked(ctx context.Context) (PracticeModeResponse, error) {
	return getAndUnmarshal[PracticeModeResponse](api, ctx, "practice/mode/ranked")
}

// PracticeStatisticsElo tries to get player statistics for all ranked modes.
//
// /api/practice/statistics/elo/{player_id}/
func (api *API) PracticeStatisticsElo(ctx context.Context, id int) (PracticeStatisticsEloResponse, error) {
	return getAndUnmarshalf[PracticeStatisticsEloResponse](api, ctx, "practice/statistics/elo/%d", id)
}

// PracticeStatisticsLeaderboardElo fetches leaderboard entries (ranked players
// statistics) for a specific mode by its id.
//
// /api/practice/statistics/leaderboard/elo/{mode_id}/
func (api *API) PracticeStatisticsLeaderboardElo(ctx context.Context, modeID int) (StatisticsEloLeaderboardResponse, error) {
	//  это угар
	return getAndUnmarshalf[StatisticsEloLeaderboardResponse](api, ctx, "practice/statistics/leaderboard/elo/%d", modeID)
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
	if err, ok := parseError(resp.Message); ok {
		return err, true
	}
	return resp, true
}

// the following methods should be API private methods, but they can't only
// because of Go doesn't support generics that way

func unmarshalResponse[T any](data []byte) (zero T, respErr bool, err error) {
	if err, ok := unmarshalResponseError(data); ok {
		return zero, true, err
	}
	var resp T
	if err := json.Unmarshal(data, &resp); err == nil {
		return resp, false, nil
	}
	return zero, false, errors.New("invalid (unhandled) response type")
}

func getAndUnmarshalf[T any](api *API, ctx context.Context, f string, a ...any) (T, error) {
	return getAndUnmarshal[T](api, ctx, fmt.Sprintf(f, a...))
}

func getAndUnmarshal[T any](api *API, ctx context.Context, path string) (zero T, err error) {
	data, err := api.get(ctx, path)
	if err != nil {
		return zero, err
	}
	resp, respErr, err := unmarshalResponse[T](data)
	if err == nil {
		return resp, nil
	}
	if respErr {
		// return error directly if it is extracted from ResponseError
		return zero, err
	}
	return zero, fmt.Errorf("unmarshal response: %w", err)
}
