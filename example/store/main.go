package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	sova "github.com/k4ties/sovapi"
	"github.com/kr/pretty"
)

const player = "lunarelly"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	api := sova.NewAPI()

	logStoreRanks(ctx, api)
	logStoreItems(ctx, api)
}

func logStoreRanks(ctx context.Context, api *sova.API) {
	resp, err := api.StoreRanks(ctx, player)
	if err != nil {
		fmt.Printf("error: call /api/store/ranks/: %v\n", err)
		return
	}
	_, _ = pretty.Println(resp)
}

func logStoreItems(ctx context.Context, api *sova.API) {
	resp, err := api.StoreItems(ctx, player)
	if err != nil {
		fmt.Printf("error: call /api/store/items/: %v\n", err)
		return
	}
	_, _ = pretty.Println(resp)
}
