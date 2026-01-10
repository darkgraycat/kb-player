package cmd

import (
	"kbplayer/internal/audio"
	"kbplayer/internal/tui"
	"os"
	"time"
)

type Command int

const (
	CmdStart Command = iota
	CmdStop
	CmdExit
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

	go audioLoop(output, wavChan, ctlChan)

	rawModeLoop(cfg, wavChan, ctlChan)

	tui.ClearScreen()
	return nil
}

func rawModeLoop(cfg *Config, wavChan chan *audio.WAV, ctlChan chan Command) {
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
				ctlChan <- CmdStop
				return nil, nil
			}

			if wav, ok := wavMap[ch]; ok {
				wavChan <- wav
			}
		}
	})
}

func audioLoop(output audio.Output, wavChan chan *audio.WAV, ctlChan chan Command) {
	// TODO: simulate streaming using ticker
	// or even move it into new audio.Output implementation
	ticker := time.NewTicker(60 * time.Millisecond)
	defer ticker.Stop()

	running := true
	for running {
		select {
		case wav := <-wavChan:
			go output.Play(wav)
		case <-ticker.C:
		case cmd := <-ctlChan:
			switch cmd {
			case CmdStop:
				running = false
			}
		}
	}
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
