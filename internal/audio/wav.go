package audio

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

const AMP = 30000.0

type WAV struct {
	SampleRate int
	Channels   int
	Header     []byte
	Samples    []int16
}

func NewWAV(sr, channels int) *WAV {
	return &WAV{
		SampleRate: sr,
		Channels:   channels,
		Header:     makeWAVHeader(sr, channels, 2),
		Samples:    []int16{},
	}
}

func (wav *WAV) AddSamples(samples []int16) {
	wav.Samples = append(wav.Samples, samples...)
}

func (wav *WAV) AddTone(freq, dur float64) {
	sr := float64(wav.SampleRate)
	numSamples := int(sr * dur)

	start := len(wav.Samples)
	wav.AddSamples(make([]int16, numSamples))

	phase := 0.0
	step := 2 * math.Pi * freq / sr

	attack := int(sr * 0.01)
	release := int(sr * 0.02)

	attackStep := 1.0 / float64(attack)
	releaseStep := 1.0 / float64(release)

	for i := range numSamples {
		env := 1.0

		if i < attack {
			env = attackStep * float64(i)
		} else if i >= numSamples-release {
			env = releaseStep * float64(numSamples-i)
		}

		wav.Samples[start+i] = int16(math.Sin(phase) * env * AMP)
		phase += step
	}
}

func (wav *WAV) AddChord(freqs []float64, dur float64) {
	sr := float64(wav.SampleRate)
	numSamples := int(sr * dur)

	start := len(wav.Samples)
	wav.AddSamples(make([]int16, numSamples))

	phases := make([]float64, len(freqs))
	steps := make([]float64, len(freqs))
	for i, freq  := range freqs {
		steps[i] = 2 * math.Pi * freq / sr
	}

	attack := int(sr * 0.01)
	release := int(sr * 0.02)

	attackStep := 1.0 / float64(attack)
	releaseStep := 1.0 / float64(release)

	numFreqs := float64(len(freqs))

	for i := range numSamples {
		env := 1.0

		if i < attack {
			env = attackStep * float64(i)
		} else if i >= numSamples-release {
			env = releaseStep * float64(numSamples-i)
		}

		mixed := 0.0
		for j := range freqs {
			mixed += math.Sin(phases[j])
			phases[j] += steps[j]
		}

		wav.Samples[start+i] = int16((mixed / numFreqs) * env * AMP)
	}
}

func (wav *WAV) WriteFull(w io.Writer) error {
	dataSize := uint32(len(wav.Samples) * 2)
	fileSize := uint32(36 + dataSize)

	binary.LittleEndian.PutUint32(wav.Header[4:8], fileSize)
	binary.LittleEndian.PutUint32(wav.Header[40:44], dataSize)

	if _, err := w.Write(wav.Header); err != nil {
		return err
	}

	return wav.WriteSamples(w)
}

func (wav *WAV) WriteSamples(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, wav.Samples)
}

func makeWAVHeader(sr, channels, bps int) []byte {
	buf := new(bytes.Buffer)
	buf.Write([]byte("RIFF"))
	buf.Write(make([]byte, 4)) // filesize placeholder
	buf.Write([]byte("WAVEfmt "))

	binary.Write(buf, binary.LittleEndian, uint32(16))              // fmt chunk size
	binary.Write(buf, binary.LittleEndian, uint16(1))               // audio format (1 PCM)
	binary.Write(buf, binary.LittleEndian, uint16(channels))        // number of channels
	binary.Write(buf, binary.LittleEndian, uint32(sr))              // sample rate
	binary.Write(buf, binary.LittleEndian, uint32(sr*channels*bps)) // byte rate
	binary.Write(buf, binary.LittleEndian, uint16(channels*bps))    // block align
	binary.Write(buf, binary.LittleEndian, uint16(bps*8))           // bits per sample

	buf.Write([]byte("data"))
	buf.Write(make([]byte, 4)) // datasize placeholder

	return buf.Bytes()
}

