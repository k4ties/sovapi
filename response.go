package sova

var _ error = (*ResponseError)(nil)

type ResponseError struct {
	Message string `json:"message"`
}

func (e ResponseError) Error() string {
	return e.Message
}

// player/{id}

type Player struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

type PlayerResponse struct {
	Data *Player `json:"data"`
}

// player/search

type PlayerSearchResponse struct {
	Data []Player `json:"data"`
}

// practice/mode; practice/mode/ranked

type PracticeMode struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Ranked      bool   `json:"ranked"`
}

type PracticeModeResponse struct {
	Data []PracticeMode `json:"data"`
}

// practice/statistics/elo/{player_id}

type RankedModeStatistic struct {
	ModeID   int    `json:"mode_id"`
	ModeName string `json:"mode_name"`
	Elo      int    `json:"amount"`
}

type PracticeStatisticsEloResponse struct {
	Data []RankedModeStatistic `json:"data"`
}

// practice/statistics/leaderboard/elo/{mode_id}

type RankedPlayerStatistic struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Elo      int    `json:"amount"`
}

type StatisticsEloLeaderboardResponse struct {
	Data []RankedPlayerStatistic `json:"data"`
}

// TODO store api
