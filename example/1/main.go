package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	sova "github.com/k4ties/sovapi"
	"github.com/kr/pretty"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	api := sova.NewAPI()

	resp, err := api.PracticeMode(ctx)
	if err != nil {
		panic(fmt.Errorf("call /practice/mode: %w", err))
	}
	_, _ = pretty.Println(resp)
}
