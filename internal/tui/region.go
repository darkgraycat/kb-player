package tui

import "fmt"

type Region struct {
	X, Y          int
	Width, Height int
	currentLine   int
}

func NewRegion(x, y, w, h int) *Region {
	return &Region{
		X: x, Y: y,
		Width: w, Height: h,
		currentLine: y + 1,
	}
}

func (r *Region) AddContent(lines ...string) {
	for i, line := range lines {
		if i >= r.Height {
			break
		}
		Move(r.X+1, r.currentLine)
		r.currentLine++
		fmt.Print(line)
	}
}

func (r *Region) ClearContent() {
	ClearRect(r.X+1, r.Y+1, r.Width-2, r.Height-2)
	r.currentLine = r.Y + 1
}

func (r *Region) DrawBorder(color int) *Region {
	WithColor(color)
	DrawBorder(r.X, r.Y, r.Width, r.Height)
	WithColor(ClrReset)
	return r
}

func (r *Region) DrawTitle(title string, color int) *Region {
	Move(r.X+2, r.Y)
	WithColor(color)
	fmt.Printf("[%s]", title)
	WithColor(ClrReset)
	return r
}

func (r *Region) MoveInside() *Region {
	Move(r.X+1, r.Y+1)
	return r
}
