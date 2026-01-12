package cmd

import (
	"fmt"
	"kbplayer/internal/audio"

	"github.com/BurntSushi/toml"
)

type Key byte

func (k *Key) UnmarshalTOML(data any) error {
	switch v := data.(type) {
	case string:
		if len(v) != 1 {
			return fmt.Errorf("key must be a single character string")
		}
		*k = Key(v[0])
	case int64:
		if v < 0 || v > 255 {
			return fmt.Errorf("key int must be 0-255, got %d", v)
		}
		*k = Key(v)
	default:
		return fmt.Errorf("key must be string or int, got %T", data)
	}
	return nil
}

type Config struct {
	Audio struct {
		SampleRate int     `toml:"sample_rate"`
		Channels   int     `toml:"channels"`
		Duration   float64 `toml:"duration"`
	} `toml:"audio"`
	Output struct {
		Mode    string
		Command string
		Args    []string
	} `toml:"output"`
	Notes  map[string]audio.Note `toml:"notes"`
	Keymap struct {
		Quit Key `toml:"quit"`
	} `toml:"keymap"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
