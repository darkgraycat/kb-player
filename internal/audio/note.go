package audio

import "math"

const C3Freq = 130.81

const (
	// Octave 3
	NoteC3 = iota
	NoteCs3
	NoteD3
	NoteDs3
	NoteE3
	NoteF3
	NoteFs3
	NoteG3
	NoteGs3
	NoteA3
	NoteAs3
	NoteB3
	// Octave 4
	NoteC4
	NoteCs4
	NoteD4
	NoteDs4
	NoteE4
	NoteF4
	NoteFs4
	NoteG4
	NoteGs4
	NoteA4
	NoteAs4
	NoteB4
	// Octave 5
	NoteC5
	NoteCs5
	NoteD5
	NoteDs5
	NoteE5
	NoteF5
	NoteFs5
	NoteG5
	NoteGs5
	NoteA5
	NoteAs5
	NoteB5
)

func NoteFreq(note int) float64 {
	// f = f0 * 2**(n/12)
	return C3Freq * math.Pow(2, float64(note)/12)
}
