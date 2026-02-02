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
var cursorX int = 1
// Current terminal cursor position Y
var cursorY int = 1

// Clears whole terminal screen
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// Prints spaces over specified rect
func ClearRect(x, y, w, h int) {
	for i := range h {
		Move(x, y+i)
		fmt.Print(strings.Repeat(" ", w))
	}
}

// Get terminal cursor position
func GetCursor() (int, int) {
	return cursorX, cursorY
}

// Move terminal cursor
func Move(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
	cursorX = x
	cursorY = y
}

// Draw border inside specified rectangle
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

// Draw menu with multiple options
func DrawMenu(options []string, current int) {
	for i, opt := range options {
		prefix := "  "
		if i == current {
			prefix = "\033[33m>\033[0m "
		}
		fmt.Printf("%s%s\r\n", prefix, opt)
		Move(cursorX, cursorY+1)
	}
}

// Set current terminal color
func SetColor(code int) {
	fmt.Printf("\033[%dm", code)
}
