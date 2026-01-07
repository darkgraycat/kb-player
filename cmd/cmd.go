package cmd

import (
	"kbplayer/internal/audio"
	"kbplayer/internal/tui"
	"os"
)

func Execute(cfg *Config) error {
	keys := setupKeys(cfg)
	// TODO: next setup synth

	var output audio.Output
	if cfg.Output.Mode == "stream" {
		output = audio.NewStreamOutput(cfg.Output.Command, cfg.Output.Args...)
	} else {
		output = audio.NewFileOutput(cfg.Output.Command, cfg.Output.Args...)
	}

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

			// TODO: simulate streaming
			if freq, ok := keys[ch]; ok {
				go play(
					freq,
					cfg.Audio.Duration,
					cfg.Audio.SampleRate,
					cfg.Audio.Channels,
					output,
				)
			}
		}
	})

	tui.ClearScreen()
	return nil
}

func play(freq, dur float64, sr, ch int, o audio.Output) error {
	// TODO: cache all notes from available keys (keys config first)
	w := audio.NewWAV(sr, ch)
	w.AddTone(freq, dur)
	return o.Play(w)
}

func setupKeys(cfg *Config) map[byte]float64 {
	keys := make(map[byte]float64, len(cfg.Notes))
	for key, code := range cfg.Notes {
		keys[key[0]] = audio.NoteFreq(code)
	}
	return keys
}

