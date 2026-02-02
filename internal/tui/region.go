package tui

import "fmt"

type Region struct {
	X, Y          int
	Width, Height int
	currentLine   int
}

// TODO: create internal state for regions to be able to 
// append data by char not only by lines
func NewRegion(x, y, w, h int) *Region {
	return &Region{
		X: x, Y: y,
		Width: w, Height: h,
		currentLine: y + 1,
	}
}

func (r *Region) AppendLine(line string, a ...any) *Region {
	// TODO: its weird that r.currentLine represents global line
	if r.currentLine >= r.Y+r.Height {
		return r
	}
	Move(r.X+1, r.currentLine)
	fmt.Printf(line, a...)
	r.currentLine++
	return r
}

func (r *Region) ReplaceLines(lines ...string) {
	r.Clear()
	for i, line := range lines {
		if i >= r.Height { break }
		Move(r.X+1, r.currentLine)
		r.currentLine++
		fmt.Print(line)
	}
}

func (r *Region) Clear() *Region {
	ClearRect(r.X+1, r.Y+1, r.Width-2, r.Height-2)
	r.currentLine = r.Y + 1
	return r
}

// TODO: do border as "required" stuff
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
