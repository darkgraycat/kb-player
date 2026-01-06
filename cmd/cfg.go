package cmd

import "github.com/BurntSushi/toml"

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
	Notes  map[string]int `toml:"notes"`
	Keymap struct {
		Quit string `toml:"quit"`
	} `toml:"keymap"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
