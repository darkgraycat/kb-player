package cmd

import (
	"kbplayer/internal/audio"
	"kbplayer/internal/tui"
	"os"
	"time"
)

func Execute(cfg *Config) error {
	var output audio.Output
	if cfg.Output.Mode == "stream" {
		output = audio.NewStreamOutput(cfg.Output.Command, cfg.Output.Args...)
		panic("Stream mode is not implemented yet")
	} else {
		output = audio.NewFileOutput(cfg.Output.Command, cfg.Output.Args...)
	}

	wavChan := make(chan *audio.WAV, 10)
	ctlChan := make(chan Command, 1)
	defer close(wavChan)
	defer close(ctlChan)

	go audioLoop(output, wavChan, ctlChan)

	rawModeLoop(cfg, wavChan, ctlChan)

	return nil
}

func rawModeLoop(cfg *Config, wavChan chan *audio.WAV, ctlChan chan Command) {
	cmdMap := setupCommandMap(cfg)
	wavMap := setupWavMap(cfg)

	tui.WithRaw(int(os.Stdin.Fd()), func() (any, error) {
		buf := make([]byte, 1)
		tui.Render()
		for {
			if _, err := os.Stdin.Read(buf); err != nil {
				return nil, err
			}
			ch := buf[0]

			if cmd, ok := cmdMap[ch]; ok {
				ctlChan <- cmd
				return nil, nil
			}

			if wav, ok := wavMap[ch]; ok {
				wavChan <- wav
			}
		}
	})

	tui.ClearScreen()
}

func audioLoop(output audio.Output, wavChan chan *audio.WAV, ctlChan chan Command) {
	// TODO: simulate streaming using ticker
	// or even move it into new audio.Output implementation
	ticker := time.NewTicker(120 * time.Millisecond)
	defer ticker.Stop()

	running := true
	for running {
		select {
		case wav := <-wavChan:
			go output.Play(wav)
		case <-ticker.C:
		case cmd := <-ctlChan:
			switch cmd {
			case CommandStop:
				running = false
			}
		}
	}
}

func setupWavMap(cfg *Config) map[byte]*audio.WAV {
	notes := make(map[byte]*audio.WAV, len(cfg.Notes))
	for key, note := range cfg.Notes {
		w := audio.NewWAV(cfg.Audio.SampleRate, cfg.Audio.Channels)
		// w.AddTone(audio.NoteFreq(note), cfg.Audio.Duration)
		w.AddTone(note.Freq(), cfg.Audio.Duration)
		notes[key[0]] = w
	}
	return notes
}

func setupCommandMap(cfg *Config) map[byte]Command {
	return map[byte]Command{
		byte(cfg.Keymap.Quit): CommandExit,
		byte(cfg.Keymap.Stop): CommandStop,
	}
}
