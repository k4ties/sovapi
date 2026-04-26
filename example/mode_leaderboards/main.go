package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"os/signal"
	"syscall"

	sova "github.com/k4ties/sovapi"
	"github.com/kr/pretty"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	api := sova.NewAPI()

	mode := randomRankedMode(api, ctx)
	fmt.Printf("mode_id=%d ; mode_name=%q\n", mode.ID, mode.Name)

	resp, err := api.PracticeStatisticsLeaderboardElo(ctx, mode.ID)
	if err != nil {
		panic(fmt.Errorf("call /practice/statistics/leaderboard/elo: %w", err))
	}
	_, _ = pretty.Println(resp)
}

func randomRankedMode(api *sova.API, ctx context.Context) sova.PracticeMode {
	resp, err := api.PracticeModeRanked(ctx)
	if err != nil {
		panic(fmt.Errorf("call /practice/mode: %w", err))
	}
	if len(resp.Data) == 0 {
		panic(errors.New("/api/practice/mode/: 0 modes in response"))
	}
	return resp.Data[rand.IntN(len(resp.Data))]
}
