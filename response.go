package sova

import "time"

var _ error = (*ResponseError)(nil)

type ResponseError struct {
	Message string `json:"message"`
	Success *bool  `json:"success,omitempty"` //only appears in /store/verify-player
}

func (e ResponseError) Error() string {
	return e.Message
}

// player/{id}

type Timestamp string

func (t Timestamp) Parse() (time.Time, error) {
	return time.Parse(time.RFC3339Nano, string(t))
}

type Duration int

func (d Duration) D() time.Duration {
	return time.Second * time.Duration(d)
}

type Rank struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	ColoredName string `json:"colored_name"`
}

type Statistics struct {
	Kills      int      `json:"kills"`
	Deaths     int      `json:"deaths"`
	Wins       int      `json:"wins"`
	Losses     int      `json:"losses"`
	KillStreak int      `json:"kill_streak"`
	WinStreak  int      `json:"win_streak"`
	PlayTime   Duration `json:"play_time"`
}

type Punishment struct {
	Reason    string     `json:"reason"`
	ExpiresAt *Timestamp `json:"expires_at,omitempty"`
}

type Player struct {
	ID         int         `json:"id"`
	Nickname   string      `json:"nickname"`
	CreatedAt  Timestamp   `json:"created_at"` //TODO: джавид гондон
	Rank       Rank        `json:"rank"`
	Statistics Statistics  `json:"statistics"`
	Ban        *Punishment `json:"ban,omitempty"`
	Mute       *Punishment `json:"mute,omitempty"`
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
