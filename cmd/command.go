package cmd

import "fmt"

type Command uint8

const (
	CommandStart Command = iota
	CommandStop
	CommandExit
)

type Key byte

func (k *Key) UnmarshalTOML(data any) error {
	switch v := data.(type) {
	case string:
		if len(v) != 1 {
			return fmt.Errorf("command key must be a single character string")
		}
		*k = Key(v[0])
	case int64:
		if v < 0 || v > 255 {
			return fmt.Errorf("command key int must be 0-255, got %d", v)
		}
		*k = Key(v)
	default:
		return fmt.Errorf("command key must be a string or int, got %T", data)
	}

	return nil
}
