package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	sova "github.com/k4ties/sovapi"
)

// нодебаф модерн ранкед = 1
// нодебаф лоу = 2
// нодебаф хайгх = 3
// гэпл ранкед = 4
// мидфайт ранкед = 5

const (
	NoDebuffModernRanked = iota + 1
	_
	_
	GAppleRanked
	MidFightRanked
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	api := sova.NewAPI()

	resp, err := api.PracticeStatisticsLeaderboardElo(ctx, NoDebuffModernRanked)
	if err != nil {
		panic(fmt.Errorf("do player search: %w", err))
	}
	for i, p := range resp.Data {
		fmt.Printf("%d) %s (%d)\n", i+1, p.Nickname, p.Elo)
	}
}
