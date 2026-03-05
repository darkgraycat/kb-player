package cmd

import "kbplayer/internal/tui"

type Ui struct {
	mode   Mode
	header *tui.Region
	body   *tui.Region
	footer *tui.Region
}

func SetupTui(cfg *Config, width int) Ui {
	tui.Clear()
	ui := Ui{}

	ui.header = tui.NewRegion(1, 1, width, 4).
		DrawBorder(tui.ClrRed).
		DrawTitle("KB Player v0.0", 0).
		WriteLine("Sample rate: %d\tDuration: %f\tMode: %s", cfg.Audio.SampleRate, cfg.Audio.Duration, cfg.Output.Mode).
		WriteLine("Quit: %v, Play: %v", cfg.Keymap.Quit, cfg.Keymap.Play)
	ui.body = tui.NewRegion(1, ui.header.Y+ui.header.Height, width, 16).
		DrawBorder(tui.ClrCyan).
		DrawTitle("Main", 0).
		WriteLine("Characters pressed will appear below:")
	ui.footer = tui.NewRegion(1, ui.body.Y+ui.body.Height, width, 3).
		DrawBorder(tui.ClrGreen).
		DrawTitle(ModeNormal.String(), 0)

	tui.Move(0, 0)
	return ui
}
