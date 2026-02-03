package cmd

import (
	"fmt"
	"kbplayer/internal/tui"
	"os"
)

func TestExecute(cfg *Config) error {
	tui.Clear()

	buf := make([]byte, 1)

	// TODO: bug - with x:0 and y:0 - border appeared broken
	r1 := tui.NewRegion(1, 1, 20, 6)
	r1.DrawBorder(tui.ClrRed)
	r1.DrawTitle("Region 1", tui.ClrMagenta)
	r1.AppendLine("Im line nr 1 of %v", 1)
	r1.AppendLine("Im line nr 2 of %v", 1)

	r2 := tui.NewRegion(1, 7, 8, 4)
	r2.DrawBorder(tui.ClrBlue)
	r2.DrawTitle("Region 2", tui.ClrCyan)
	r2.AppendLine("Im line nr 1 of %v", 2)
	r2.AppendLine("Im line nr 2 of %v", 2)
	r2.AppendLine("Im line nr 3 of %v", 2)

	r1.MoveInside()
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
