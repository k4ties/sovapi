package sova

import (
	"errors"
	"regexp"
)

var _ error = (*ResponseError)(nil)

type ResponseError struct {
	Message string `json:"message"`
}

func (e ResponseError) Error() string {
	return e.Message
}

var (
	ErrNicknameMustBeTwoChars = errors.New("nickname must be at least 2 characters long")
	ErrServerError            = errors.New("server error")
	ErrCannotFindPlayer       = errors.New("cannot find player")
	ErrRouteNotFound          = errors.New("route not found")
	ErrNoSuchMode             = errors.New("no such mode")
)

var (
	regPlayerNotFound = regexp.MustCompile(`No query results for model \[App\\Models\\Player\\Player] .*`)
	regRouteNotFound  = regexp.MustCompile(`The route .* could not be found\.`)
	regModeNotFound   = regexp.MustCompile(`No query results for model \[App\\Models\\Practice\\PracticeMode] .*`)
)

func parseError(message string) error {
	switch {
	case message == "Nickname must be at least 2 characters long":
		return ErrNicknameMustBeTwoChars
	case message == "Server Error":
		return ErrServerError
	case regPlayerNotFound.MatchString(message):
		return ErrCannotFindPlayer //todo: pass id
	case regModeNotFound.MatchString(message):
		return ErrNoSuchMode // ...
	case regRouteNotFound.MatchString(message):
		return ErrRouteNotFound // ...
	}
	return nil
}
