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

	resp, err := api.PlayerSearch(ctx, "liss")
	if err != nil {
		panic(fmt.Errorf("do player search: %w", err))
	}
	fmt.Printf("%#v\n", resp.Data)
}
