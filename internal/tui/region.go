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

func (r *Region) AddLine(lines ...string) {
	for i, line := range lines {
		if i >= r.Height {
			break
		}
		Move(r.X+1, r.currentLine)
		r.currentLine++
		fmt.Print(line)
	}
}

func (r *Region) Clear() {
	ClearRect(r.X+1, r.Y+1, r.Width-2, r.Height-2)
	r.currentLine = r.Y + 1
}

func (r *Region) DrawBorder(color int) *Region {
	SetColor(color)
	DrawBorder(r.X, r.Y, r.Width, r.Height)
	SetColor(ClrReset)
	return r
}

func (r *Region) DrawTitle(title string, color int) *Region {
	Move(r.X+2, r.Y)
	SetColor(color)
	fmt.Printf("[%s]", title)
	SetColor(ClrReset)
	return r
}

func (r *Region) MoveInside() *Region {
	Move(r.X+1, r.Y+1)
	return r
}
