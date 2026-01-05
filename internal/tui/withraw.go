package tui

import (
	"golang.org/x/term"
)

func WithRaw[T any](fd int, fn func() (T, error)) (T, error) {
	var none T
	state, err := term.MakeRaw(fd)
	defer term.Restore(fd, state)
	if err != nil {
		return none, err
	}
	return fn()
}
