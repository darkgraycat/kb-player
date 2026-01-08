package cmd

import (
	"kbplayer/internal/audio"
	"kbplayer/internal/tui"
	"os"
)

func Execute(cfg *Config) error {
	// TODO: next setup synth

	var output audio.Output
	if cfg.Output.Mode == "stream" {
		output = audio.NewStreamOutput(cfg.Output.Command, cfg.Output.Args...)
	} else {
		output = audio.NewFileOutput(cfg.Output.Command, cfg.Output.Args...)
	}

	wavMap := setupWavMap(cfg)

	tui.WithRaw(int(os.Stdin.Fd()), func() (any, error) {
		buf := make([]byte, 1)
		tui.Render()
		for {
			if _, err := os.Stdin.Read(buf); err != nil {
				return nil, err
			}
			ch := buf[0]
			switch ch {
			case byte(cfg.Keymap.Quit):
				return nil, nil
			}

			if wav, ok := wavMap[ch]; ok {
				go output.Play(wav)
			}

		}
	})

	tui.ClearScreen()
	return nil
}

func setupWavMap(cfg *Config) map[byte]*audio.WAV {
	notes := make(map[byte]*audio.WAV, len(cfg.Notes))
	for key, code := range cfg.Notes {
		w := audio.NewWAV(cfg.Audio.SampleRate, cfg.Audio.Channels)
		w.AddTone(audio.NoteFreq(code), cfg.Audio.Duration)
		notes[key[0]] = w
	}
	return notes
}
