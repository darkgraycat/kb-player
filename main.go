package main

import (
	"fmt"
	"kbplayer/cmd"
	"os"
)

var defaultConfigPath = "config.toml"

// TODO: incude CLI arguments to override default config path
func main() {

	cfg, err := cmd.LoadConfig(defaultConfigPath)
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
