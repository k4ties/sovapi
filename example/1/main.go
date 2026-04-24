package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	sova "github.com/k4ties/sovapi"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	api := sova.NewAPI()

	resp, err := api.PracticeMode(ctx)
	if err != nil {
		panic(fmt.Errorf("do player search: %w", err))
	}
	for i, p := range resp.Data {
		fmt.Printf("%d) %s (%d); ranked=%t\n", i+1, p.DisplayName, p.ID, p.Ranked)
	}
}
