package cmd

import "kbplayer/internal/tui"

type Ui struct {
	Mode          Mode
	Header        *tui.Region
	Body          *tui.Region
	Footer        *tui.Region
	width         int
	recordedChars []byte // TODO save state somehow
}

func InitTui(cfg *Config, width int) Ui {
	tui.Clear()
	ui := Ui{
		width:         width,
		recordedChars: make([]byte, width*8),
	}

	ui.Header = tui.NewRegion(1, 1, width, 4).
		DrawBorder(tui.ClrRed).
		DrawTitle("KB Player v0.0", 0).
		WriteLine("Sample rate: %d\tDuration: %f\tMode: %s", cfg.Audio.SampleRate, cfg.Audio.Duration, cfg.Output.Mode).
		WriteLine("Quit: %v, Play: %v", cfg.Keymap.Quit, cfg.Keymap.Play)
	_, hy, _, hh := ui.Header.GetDimensions()

	ui.Body = tui.NewRegion(1, hy+hh, width, 16).
		DrawBorder(tui.ClrCyan).
		DrawTitle("Main", 0).
		WriteLine("Characters pressed will appear below:")
	_, by, _, bh := ui.Body.GetDimensions()

	ui.Footer = tui.NewRegion(1, by+bh, width, 3).
		DrawBorder(tui.ClrGreen).
		DrawTitle(ModeNormal.String(), 0)

	tui.Move(0, 0)
	return ui
}

func (ui Ui) ChangeMode(mode Mode) {
	ui.Mode = mode
	ui.Footer.DrawTitle(ui.Mode.String(), 0)
}

func (ui Ui) RecordPressedChar(ch byte) {
	bcx, bcy := ui.Body.GetCursor()
	_, _, bw, _ := ui.Body.GetDimensions()
	if bcx > bw-3 { // 2 border, 1 to avoid overlap
		ui.Body.Move(0, bcy+1)
	}
	ui.Body.Write("%c ", ch)
}
