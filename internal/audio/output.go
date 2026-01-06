package audio

import (
	"os"
	"os/exec"
)

type Output interface {
	Play(wav *WAV) error
}

type FileOutput struct {
	Command string
	Args    []string
}

func NewFileOutput(command string, args ...string) *FileOutput {
	return &FileOutput{
		Command: command,
		Args:    args,
	}
}

func (o *FileOutput) Play(wav *WAV) error {
	f, err := os.CreateTemp("", "note*.wav")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	defer f.Close()

	wav.WriteFull(f)
	args := append([]string{f.Name()}, o.Args...)

	return exec.Command(o.Command, args...).Run()
}

type StreamOutput struct {
	Command string
	Args    []string
}

func NewStreamOutput(command string, args ...string) *StreamOutput {
	return &StreamOutput{
		Command: command,
		Args:    args,
	}
}

func (o *StreamOutput) Play(wav *WAV) error {
	// TODO: implement in future using ffplay
	return nil
}
