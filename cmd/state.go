package cmd

type Mode int

const (
	ModeNormal Mode = iota
	ModeRecord
)

func (m Mode) String() string {
	switch m {
	case ModeNormal:
		return "Normal"
	case ModeRecord:
		return "Record"
	default:
		return "Unknown"
	}
}
