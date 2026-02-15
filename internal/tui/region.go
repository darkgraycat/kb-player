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

func (r *Region) Write(line string, a ...any) *Region {
	r.Focus()
	n, _ := fmt.Printf(line, a...)
	r.cx += n
	return r
}

func (r *Region) WriteLine(line string, a ...any) *Region {
	if r.cy >= r.Height-2 {
		return r
	}
	r.Write(line, a...)
	r.cx = 0
	r.cy++
	return r
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

func (r *Region) Focus() *Region {
	Move(r.X+r.cx+1, r.Y+r.cy+1)
	return r
}

func (r *Region) Move(x, y int) *Region {
	r.cx, r.cy = x, y
	return r.Focus()
}

