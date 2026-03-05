package cmd

import (
	"context"
	"kbplayer/internal/audio"
	"kbplayer/internal/tui"
	"os"
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
	defer close(wavChan)

	cmdMap := setupCommandMap(cfg)
	wavMap := setupWavMap(cfg)

	tui.WithRaw(int(os.Stdin.Fd()), func() (any, error) {
		// controls
		buf := make([]byte, 1)
		ctx := context.Background()
		var stopAudioLoop context.CancelFunc

		// ui
		ui := InitTui(cfg, 80)

		for {
			ch, err := tui.ReadBuf(os.Stdin, buf)
			if err != nil {
				return nil, err
			}

			// handle commands
			if cmd, ok := cmdMap[ch]; ok {
				switch cmd {
				case CommandQuit:
					return nil, nil
				case CommandPlay:
					if stopAudioLoop == nil {
						playCtx, cancel := context.WithCancel(ctx)
						stopAudioLoop = cancel
						go audioLoop(playCtx, output, wavChan)
						ui.ChangeMode(ModeRecord)
					} else {
						stopAudioLoop()
						stopAudioLoop = nil
						ui.ChangeMode(ModeNormal)
					}
				}
			}

			// play notes
			if stopAudioLoop != nil {
				if wav, ok := wavMap[ch]; ok {
					wavChan <- wav
					ui.RecordPressedChar(ch)
				}
			}

			tui.Move(0, 0) // reset cursor
		}
	})

	tui.Clear()
	return nil
}

func audioLoop(ctx context.Context, output audio.Output, wavChan chan *audio.WAV) {
	for {
		select {
		case wav := <-wavChan:
			go output.Play(wav)
		case <-ctx.Done():
			return
		}
	}
}

func setupWavMap(cfg *Config) map[byte]*audio.WAV {
	notes := make(map[byte]*audio.WAV, len(cfg.Notes))
	for key, note := range cfg.Notes {
		w := audio.NewWAV(cfg.Audio.SampleRate, cfg.Audio.Channels)
		w.AddTone(note.Freq(), cfg.Audio.Duration)
		notes[key[0]] = w
	}
	return notes
}

func setupCommandMap(cfg *Config) map[byte]Command {
	return map[byte]Command{
		byte(cfg.Keymap.Quit): CommandQuit,
		byte(cfg.Keymap.Play): CommandPlay,
	}
}
