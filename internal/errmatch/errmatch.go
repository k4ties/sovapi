// Package errmatch represents error matching utilities.
//
// This may be useful to match errors from the third party server (in case of
// you know the pattern).
package errmatch

func Match(e []Entry, source string) error {
	for _, entry := range e {
		if err := entry.Match(source); err != nil {
			return err
		}
	}
	return nil
}
