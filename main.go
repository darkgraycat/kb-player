package main

import (
	"fmt"
	"kbplayer/cmd"
	"os"
)

var defaultConfigPath = "config.toml"

func main() {
	configPath := defaultConfigPath

	args := os.Args
	if len(args) > 1 {
		configPath = args[1]
	}

	cfg, err := cmd.LoadConfig(configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading config")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = cmd.Execute(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error during execution")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
