package sova

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/k4ties/sovapi/internal/errmatch"
)

var ErrServerError = errors.New("server error")

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

var errMatches = []errmatch.Entry{
	{
		Message: "Nickname must be at least 2 characters long",
		Ret:     errmatch.Ret(ErrNicknameMustBeTwoChars),
	},
	{
		Message: "Server Error",
		Ret:     errmatch.Ret(ErrServerError),
	},
	{
		Message: `No query results for model [App\Models\Player\Player].`, // it can also return like this
		Regex:   regexp.MustCompile(`No query results for model \[App\\Models\\Player\\Player] (.*)`),
		Ret:     func(m string) error { return ErrCannotFindPlayer{Player: m} },
	},
	{
		Regex: regexp.MustCompile(`The route (.*) could not be found\.`),
		Ret:   func(m string) error { return ErrRouteNotFound{Route: m} },
	},
	{
		Regex: regexp.MustCompile(`No query results for model \[App\\Models\\Practice\\PracticeMode] (.*)`),
		Ret:   func(m string) error { return ErrNoSuchMode{Mode: m} },
	},
}
