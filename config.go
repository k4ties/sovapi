package sova

import (
	"net/http"
	"time"
)

type Config struct {
	Client         *http.Client
	RequestTimeout time.Duration
}

func (conf Config) New() *API {
	if conf.Client == nil {
		conf.Client = http.DefaultClient
	}
	if conf.RequestTimeout <= 0 {
		conf.RequestTimeout = time.Second * 5
	}
	return &API{conf: conf}
}
