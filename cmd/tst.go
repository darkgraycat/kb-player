package cmd

import (
	"fmt"
	"kbplayer/internal/tui"
	"os"
)

func TestExecute(cfg *Config) error {
	tui.Clear()

	buf := make([]byte, 1)

	r1 := tui.NewRegion(1, 1, 20 + 2, 4 + 2)
	r1.DrawBorder(tui.ClrRed)
	r1.DrawTitle("Region 1", tui.ClrMagenta)
	r1.WriteLine("Im line nr 1 of %v", 1)
	r1.WriteLine("Im line nr 2 of %v", 1)
	r1.WriteLine("01234567890123456789")

	r2 := tui.NewRegion(1, 7, 20 + 2, 3 + 2)
	r2.DrawBorder(tui.ClrBlue)
	r2.DrawTitle("Region 2", tui.ClrCyan)
	r2.Write("Word")
	r2.Write("AnotherLol")

	r1.Focus()
	fmt.Print("Using inside")

	


	tui.WithRaw(int(os.Stdin.Fd()), func() (any, error) {
		for {
			ch, err := tui.ReadBuf(os.Stdin, buf)
			if err != nil {
				return nil, err
			}
			if ch == 3 {
				return nil, nil
			}

		}
	})

	tui.Clear()
	return nil
}
