package audio

import (
	"fmt"
	"kbplayer/internal/utils"
	"math"
)

const C3Freq = 130.81

type Note int

const (
	// Octave 3
	NoteC3 Note = iota
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

func NoteFreq(note Note) float64 {
	// f = f0 * 2**(n/12)
	return C3Freq * math.Pow(2, float64(note)/12)
}

func StrToNote(s string) (Note, error) {
	var octave, code int
	switch len(s) {
	case 2:
		octave = int(s[1] - '0')
	case 3:
		octave = int(s[2] - '0')
		code = utils.Match(s[1], map[byte]int{'#': 1, 'b': -1}, 0)
	default:
		return 0, fmt.Errorf("string should have 2 or 3 characters")
	}
	code += utils.Match(s[0], map[byte]int{
		'C': 0,
		'D': 2,
		'E': 4,
		'F': 5,
		'G': 7,
		'A': 9,
		'B': 11,
	}, 0)
	return Note((octave-3)*12 + code), nil
}

func (n *Note) UnmarshalTOML(data any) error {
	switch v := data.(type) {
	case string:
		note, err := StrToNote(v)
		if err != nil {
			return err
		}
		*n = note
	default:
		return fmt.Errorf("key must be a string, got %T", data)
	}
	return nil
}
