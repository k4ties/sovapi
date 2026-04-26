package main

import (
	"context"
	"fmt"
	"log/slog"
	"os/signal"
	"strings"
	"syscall"

	sova "github.com/k4ties/sovapi"
	"github.com/kr/pretty"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	log := slog.Default()
	api := sova.NewAPI()

	p, err := searchExactPlayer(ctx, api, "lunarelly")
	if err != nil {
		panic(fmt.Errorf("search exact player: %w", err))
	}
	log.Info("found player with exact match", "id", p.ID, "nickname", p.Nickname)
	log.Info("trying to fetch player ranked statistics...")

	resp, err := api.PracticeStatisticsElo(ctx, p.ID)
	if err != nil {
		panic(fmt.Errorf("call /practice/statistics/elo: %w", err))
	}
	_, _ = pretty.Println(resp)
}

func searchExactPlayer(ctx context.Context, api *sova.API, name string) (sova.Player, error) {
	// since their current API doesn't support searching by name, we will use
	// /api/player/search/ and check for exact name match (case-insensitive)
	resp, err := api.PlayerSearch(ctx, name)
	if err != nil {
		return sova.Player{}, fmt.Errorf("call /api/player/search: %w", err)
	}
	for _, p := range resp.Data {
		if strings.EqualFold(p.Nickname, name) {
			return p, nil
		}
	}
	return sova.Player{}, sova.ErrCannotFindPlayer{}
}
