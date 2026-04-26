package errmatch

import "regexp"

type Entry struct {
	Message string

	Regex     *regexp.Regexp
	FindMatch func([]string) (string, bool) //optional

	Ret func(matched string) error
}

func (e Entry) Match(to string) error {
	if e.Message != "" && to == e.Message {
		return e.Ret("")
	}
	if e.Regex == nil {
		return nil
	}
	m, ok := e.findMatchRegex(to)
	if ok {
		return e.Ret(m)
	}
	return nil
}

func (e Entry) findMatchRegex(msg string) (string, bool) {
	if e.Regex == nil {
		return "", false
	}
	x := e.Regex.FindStringSubmatch(msg)
	if e.FindMatch != nil {
		return e.FindMatch(x)
	}
	if len(x) > 1 {
		return x[1], true
	}
	return "", false
}

// Ret is a helper function that you can use to directly pass error in Entry.
func Ret(err error) func(string) error {
	return func(string) error {
		return err
	}
}
