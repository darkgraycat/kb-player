package tui

import "fmt"

type Region struct {
	X, Y          int
	Width, Height int
	cx, cy        int
}

func NewRegion(x, y, w, h int) *Region {
	return &Region{
		X: x, Y: y,
		Width: w, Height: h,
	}
}

func (r *Region) AppendLine(line string, a ...any) *Region {
	if r.cy >= r.Height-2 {
		return r
	}
	Move(r.X+1, r.Y+1+r.cy)
	n, _ := fmt.Printf(line, a...)
	r.cx = n
	r.cy++
	return r
}

func (r *Region) ReplaceLines(lines ...string) {
	r.Clear()
	for i, line := range lines {
		if i >= r.Height {
			break
		}
		r.AppendLine("%s", line)
	}
}

func (r *Region) Clear() *Region {
	ClearRect(r.X+1, r.Y+1, r.Width-2, r.Height-2)
	r.cx, r.cy = 0, 0
	return r
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
