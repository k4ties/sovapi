package sova

import (
	"errors"
	"fmt"
)

var (
	ErrNicknameMustBeTwoChars = errors.New("nickname must be at least 2 characters long")
	ErrServerError            = errors.New("server error")
)

type ErrRouteNotFound struct {
	Route string
}

func (e ErrRouteNotFound) Error() string {
	if e.Route == "" {
		return "route not found"
	}
	return fmt.Sprintf("route not found: %s", e.Route)
}

type ErrNoSuchMode struct {
	Mode string
}

func (e ErrNoSuchMode) Error() string {
	if e.Mode == "" {
		return "no such mode"
	}
	return fmt.Sprintf("no such mode: %s", e.Mode)
}

type ErrCannotFindPlayer struct {
	Player string
}

func (e ErrCannotFindPlayer) Error() string {
	if e.Player == "" {
		return "cannot find player"
	}
	return fmt.Sprintf("cannot find player: %s", e.Player)
}

type ErrUnmarshalResponse struct {
	Parent error
}

func (e ErrUnmarshalResponse) Error() string {
	return fmt.Sprintf("unmarshal response: %v", e.Parent)
}

func (e ErrUnmarshalResponse) Unwrap() error {
	return e.Parent
}
