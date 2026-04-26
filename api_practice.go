package sova

import (
	"context"
	"fmt"
	"strconv"
)

type ErrNoModesAvailable struct {
	Ranked bool
}

func (e ErrNoModesAvailable) Error() string {
	if e.Ranked {
		return "no ranked modes available"
	}
	return "no modes available"
}

// PracticeMode returns a list of all available practice modes, including
// ranked.
//
// /api/practice/mode/
func (api *API) PracticeMode(ctx context.Context) (*PracticeModeResponse, error) {
	resp, err := getAndUnmarshal[PracticeModeResponse](api, ctx, "practice/mode")
	if err != nil {
		return nil, err
	}
	if len(resp.Data) == 0 {
		return nil, ErrNoModesAvailable{}
	}
	return resp, nil
}

// PracticeModeRanked returns a list of all available ranked practice modes.
//
// /api/practice/mode/ranked/
func (api *API) PracticeModeRanked(ctx context.Context) (*PracticeModeResponse, error) {
	resp, err := getAndUnmarshal[PracticeModeResponse](api, ctx, "practice/mode/ranked")
	if err != nil {
		return nil, err
	}
	if len(resp.Data) == 0 {
		return nil, ErrNoModesAvailable{Ranked: true}
	}
	return resp, nil
}

type RankedStatisticsType int

const (
	RankedStatisticsTypePlayer RankedStatisticsType = iota + 1
	RankedStatisticsTypeMode
)

type ErrNoStatisticsAvailable struct {
	Type RankedStatisticsType
	For  string
}

func (e ErrNoStatisticsAvailable) Error() string {
	if e.For == "" {
		return "no statistics available"
	}
	return fmt.Sprintf("no statistics available for %s", e.For)
}

// PracticeStatisticsElo tries to get player statistics for all ranked modes.
//
// /api/practice/statistics/elo/{player_id}/
func (api *API) PracticeStatisticsElo(ctx context.Context, playerID int) (*PracticeStatisticsEloResponse, error) {
	resp, err := getAndUnmarshalf[PracticeStatisticsEloResponse](api, ctx, "practice/statistics/elo/%d", playerID)
	if err != nil {
		return nil, err
	}
	if len(resp.Data) == 0 {
		return nil, ErrNoStatisticsAvailable{Type: RankedStatisticsTypePlayer, For: strconv.Itoa(playerID)}
	}
	return resp, nil
}

// PracticeStatisticsLeaderboardElo fetches leaderboard entries (ranked players
// statistics) for a specific mode by its id.
//
// /api/practice/statistics/leaderboard/elo/{mode_id}/
func (api *API) PracticeStatisticsLeaderboardElo(ctx context.Context, modeID int) (*StatisticsEloLeaderboardResponse, error) {
	resp, err := getAndUnmarshalf[StatisticsEloLeaderboardResponse](api, ctx, "practice/statistics/leaderboard/elo/%d", modeID)
	if err != nil {
		return nil, err
	}
	if len(resp.Data) == 0 {
		return nil, ErrNoStatisticsAvailable{Type: RankedStatisticsTypeMode, For: strconv.Itoa(modeID)}
	}
	return resp, nil
}
