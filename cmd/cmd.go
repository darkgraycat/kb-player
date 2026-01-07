package cmd

import (
	"kbplayer/internal/audio"
	"kbplayer/internal/tui"
	"os"
)

func Execute(cfg *Config) error {
	notes := setupNotes(cfg)
	// TODO: next setup synth

	var output audio.Output
	if cfg.Output.Mode == "stream" {
		output = audio.NewStreamOutput(cfg.Output.Command, cfg.Output.Args...)
	} else {
		output = audio.NewFileOutput(cfg.Output.Command, cfg.Output.Args...)
	}

	playbackLoop(output, cfg, notes)

	tui.ClearScreen()
	return nil
}

func play(freq, dur float64, sr, ch int, o audio.Output) error {
	// TODO: cache all notes from available keys (keys config first)
	w := audio.NewWAV(sr, ch)
	w.AddTone(freq, dur)
	return o.Play(w)
}

func setupNotes(cfg *Config) map[byte]float64 {
	notes := make(map[byte]float64, len(cfg.Notes))
	for key, code := range cfg.Notes {
		notes[key[0]] = audio.NoteFreq(code)
	}
	return notes
}

func playbackLoop(output audio.Output, cfg *Config, notes map[byte]float64) {
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

			// TODO: add echo support - keep note playing till next
			// TODO: simulate streaming
			if freq, ok := notes[ch]; ok {
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
}
