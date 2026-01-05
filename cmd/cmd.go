package cmd

import (
	"kbplayer/internal/audio"
	"kbplayer/internal/tui"
	"os"
	"os/exec"
)

func Execute(cfg *Config) error {
	keys := setupKeys(cfg)
	// TODO: next setup synth

	tui.WithRaw(int(os.Stdin.Fd()), func() (any, error) {
		buf := make([]byte, 1)

		tui.Render()
		for {
			if _, err := os.Stdin.Read(buf); err != nil {
				return nil, err
			}

			ch := buf[0]

			switch ch {
			case 'Q':
				return nil, nil
			}

			// TODO: simulate streaming
			if freq, ok := keys[ch]; ok {
				go play(freq, cfg)
			}
		}
	})

	tui.ClearScreen()
	return nil
}

func play(freq float64, cfg *Config) {
	// TODO: move into separate Player struct
	f, _ := os.CreateTemp("", "note*.wav")
	defer os.Remove(f.Name())
	defer f.Close()

	// TODO: cache all notes from available keys (keys config first)
	w := audio.NewWAV(cfg.Audio.SampleRate, cfg.Audio.Channels)
	w.AddTone(freq, cfg.Audio.Duration)
	w.WriteFull(f)

	args := append([]string{f.Name()}, cfg.Output.Args...)
	exec.Command(cfg.Output.Command, args...).Run()
}

func setupKeys(cfg *Config) map[byte]float64 {
	keys := make(map[byte]float64, len(cfg.Keys))
	for key, code := range cfg.Keys {
		keys[key[0]] = audio.NoteFreq(code)
	}
	return keys
}
