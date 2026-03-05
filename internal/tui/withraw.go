package tui

import (
	"io"

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

func ReadBuf(r io.Reader, buf []byte) (byte, error) {
	var none byte
	if _, err := r.Read(buf); err != nil {
		return none, err
	}
	return buf[0], nil
}
