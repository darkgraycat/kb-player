package tui

import (
	"fmt"
	"strings"
)

const (
	BorderTL = '╭'
	BorderTR = '╮'
	BorderBL = '╰'
	BorderBR = '╯'
	BorderH  = '─'
	BorderV  = '│'
)

const (
	ClrBlack   = 30
	ClrRed     = 31
	ClrGreen   = 32
	ClrYellow  = 33
	ClrBlue    = 34
	ClrMagenta = 35
	ClrCyan    = 36
	ClrWhite   = 37
	ClrReset   = 0
)

// Current terminal cursor position X
var CursorX int = 1
// Current terminal cursor position Y
var CursorY int = 1

// TODO: move it somewhere
func DrawMenu(options []string, current int) {
	for i, opt := range options {
		prefix := "  "
		if i == current {
			prefix = "\033[33m>\033[0m "
		}
		fmt.Printf("%s%s\r\n", prefix, opt)
		Move(CursorX, CursorY+1)
	}
}

// TODO: move it somewhere
func Render() {
	w, h := 80, 24
	ClearScreen()
	title := NewRegion(1, 1, w, 4).
		DrawBorder(ClrRed).
		DrawTitle("KB Player v0.0", 0)
	title.AddContent(
		"This is CLI application to play some tunes",
		"Made in order to learn Golang, and just for fun",
	)
	keymap := NewRegion(1, 5, 24, 16).
		DrawBorder(ClrBlue).
		DrawTitle("Keymap", 0)
	separator := " " + string(BorderV) + " "
	keymap.AddContent(
		// TODO: made dynamic (depends on configs)
		"Ctrl+c - to quit",
		"a - C"+separator+"s - D",
		"d - E"+separator+"f - F",
		"g - G"+separator+"h - A",
		"j - B"+separator,
		"z - C"+separator+"x - D",
		"c - E"+separator+"v - F",
		"b - G"+separator+"n - A",
		"m - B"+separator,
		"Shift - play #",
	)
	main := NewRegion(25, 5, w-24, 16).
		DrawBorder(ClrCyan).
		DrawTitle("Main", 0)
	main.AddContent("-a-s-d--g-d-s-a")
	main.AddContent("-a-s-d--g-d-s-a")
	main.ClearContent()
	options := make([]string, 0)
	options = append(options, "Start", "Menu", "Options", "Other", "Exit")
	current := 0
	main.MoveInside()
	DrawMenu(options, current)

	status := NewRegion(1, h-3, w, 3).
		DrawBorder(ClrYellow).
		DrawTitle("Status", 0)
	status.AddContent("Sample rate: 44100\tlistening for input")

	Move(2, 6) // highlight Q - to quit
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func ClearRect(x, y, w, h int) {
	for i := range h {
		Move(x, y+i)
		fmt.Print(strings.Repeat(" ", w))
	}
}

func Move(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
	CursorX = x
	CursorY = y
}

func DrawBorder(x, y, w, h int) {
	hor := string(BorderH)
	ver := string(BorderV)

	Move(x, y) // top
	fmt.Print(string(BorderTL), strings.Repeat(hor, w-2), string(BorderTR))
	// sides
	for i := 1; i < h-1; i++ {
		Move(x, y+i)
		fmt.Print(ver)
		Move(x+w-1, y+i)
		fmt.Print(ver)
	}
	Move(x, y+h-1) // bottom
	fmt.Print(string(BorderBL), strings.Repeat(hor, w-2), string(BorderBR))
}

func WithColor(code int) {
	fmt.Printf("\033[%dm", code)
}
