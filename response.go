package sova

var _ error = (*ResponseError)(nil)

type ResponseError struct {
	Message string `json:"message"`
	Success *bool  `json:"success,omitempty"` //only appears in /store/verify-player
}

func (e ResponseError) Error() string {
	return e.Message
}

// player/{id}

type Player struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

type PlayerResponse = Player

// player/search

type PlayerSearchResponse = []Player

// practice/mode

type PracticeMode struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Ranked      bool   `json:"ranked"`
}

type PracticeModeResponse = []PracticeMode

// practice/mode/ranked

type PracticeModeRankedResponse = []PracticeMode

// practice/statistics/elo/{player_id}

type RankedModeStatistic struct {
	ModeID   int    `json:"mode_id"`
	ModeName string `json:"mode_name"`
	Elo      int    `json:"amount"`
}

type PracticeStatisticsEloResponse = []RankedModeStatistic

// practice/statistics/leaderboard/elo/{mode_id}

type RankedPlayerStatistic struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Elo      int    `json:"amount"`
}

type StatisticsEloLeaderboardResponse = []RankedPlayerStatistic

// store/verify-player

type StoreVerifyPlayerResponse = bool

// store/ranks

type StoreRank struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Purchasable bool   `json:"purchasable"`
	Price       int    `json:"price"`
}

type StoreRanksResponse = []StoreRank

// store/items

type StoreItem struct { // услуга (unmute/unban)
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Price       int    `json:"price"`
}

type StoreItemsResponse = []StoreItem
