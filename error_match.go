package sova

import "regexp"

type errorMatch struct {
	msg   string
	regex *regexp.Regexp

	ret func(match string) error
}

func retDirect(err error) func(string) error {
	return func(string) error {
		return err
	}
}

var errMatches = []errorMatch{
	{
		msg: "Nickname must be at least 2 characters long",
		ret: retDirect(ErrNicknameMustBeTwoChars),
	},
	{
		msg: "Server Error",
		ret: retDirect(ErrServerError),
	},
	{
		regex: regexp.MustCompile(`No query results for model \[App\\Models\\Player\\Player] (.*)`),
		ret:   func(m string) error { return ErrCannotFindPlayer{Player: m} },
	},
	{
		regex: regexp.MustCompile(`The route (.*) could not be found\.`),
		ret:   func(m string) error { return ErrRouteNotFound{Route: m} },
	},
	{
		regex: regexp.MustCompile(`No query results for model \[App\\Models\\Practice\\PracticeMode] (.*)`),
		ret:   func(m string) error { return ErrNoSuchMode{Mode: m} },
	},
}

func parseError(message string) error {
	for _, match := range errMatches {
		if err := matchError(match, message); err != nil {
			return err
		}
	}
	return nil
}

func matchError(m errorMatch, msg string) error {
	if m.msg != "" {
		if msg == m.msg {
			return m.ret(msg)
		}
		// nothing left to compare
		return nil
	}
	x := m.regex.FindStringSubmatch(msg)
	if len(x) > 1 {
		return m.ret(x[1])
	}
	return nil
}
