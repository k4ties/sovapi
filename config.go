package sova

import (
	"net/http"
	"time"
)

const DefaultMaxBodySize uint64 = 8 * 1024 * 1024 // 8 mb

type Config struct {
	Client         *http.Client
	RequestTimeout time.Duration
	MaxBodySize    uint64
}

func (conf Config) New() *API {
	if conf.Client == nil {
		conf.Client = http.DefaultClient
	}
	if conf.RequestTimeout <= 0 {
		conf.RequestTimeout = time.Second * 5
	}
	if conf.MaxBodySize == 0 {
		conf.MaxBodySize = DefaultMaxBodySize
	}
	return &API{conf: conf}
}
