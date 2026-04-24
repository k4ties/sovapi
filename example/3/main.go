package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	sova "github.com/k4ties/sovapi"
)

// джавид = 3
// максон = 2

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	api := sova.NewAPI()

	resp, err := api.PracticeStatisticsElo(ctx, 2)
	if err != nil {
		panic(fmt.Errorf("do player search: %w", err))
	}
	for i, p := range resp.Data {
		fmt.Printf("%d) %s (%d)\n", i+1, p.ModeName, p.Elo)
	}
}
