# Terminal based keyboard player

## FAQ

### Why?
Because

### Really, why?
Because I want to learn Go

### But why keyboard player
Because its fun and kinda crazy

## How to run
Check `config.toml` file

For now it requires `afplay` to be installed.
But you can change to whatever WAV player you prefer.

Config sections:
- audio - better leave as it is, or try playing with duration
- output - here you can specify CLI command to play WAV files
- notes - bindings. supports notes only, like: C3 Cb3 C#3
- keymap - application control

## How it works
Please check `cmd/cmd.go` file.
It grabs data from `[notes]` in config, generates temp WAV files and then listens for input. Thats it.

