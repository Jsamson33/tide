package domain

import (
	"crypto/sha256"
	"encoding/binary"
	"math/rand"
	"time"
)

// Line represents a single line in a hexagram (0 for broken, 1 for solid).
type Line int

const (
	Broken Line = 0
	Solid  Line = 1
)

// Hexagram represents a stack of 6 lines.
type Hexagram struct {
	Lines [6]Line
}

// Trigram represents a stack of 3 lines.
type Trigram struct {
	Lines [3]Line
}

// GetLowerTrigram returns the bottom 3 lines.
func (h Hexagram) GetLowerTrigram() Trigram {
	return Trigram{Lines: [3]Line{h.Lines[0], h.Lines[1], h.Lines[2]}}
}

// GetUpperTrigram returns the top 3 lines.
func (h Hexagram) GetUpperTrigram() Trigram {
	return Trigram{Lines: [3]Line{h.Lines[3], h.Lines[4], h.Lines[5]}}
}

// BinaryString returns the string representation of a Trigram.
func (t Trigram) BinaryString() string {
	s := ""
	for _, l := range t.Lines {
		if l == Solid {
			s += "1"
		} else {
			s += "0"
		}
	}
	return s
}

// DrawDaily generates a deterministic hexagram for the given date.
func DrawDaily(t time.Time) Hexagram {
	dateStr := t.Format("2006-01-02")
	hash := sha256.Sum256([]byte(dateStr))

	// Use the first 8 bytes of the hash as a seed for a local PRNG
	seed := int64(binary.BigEndian.Uint64(hash[:8]))
	r := rand.New(rand.NewSource(seed))

	var h Hexagram
	for i := 0; i < 6; i++ {
		if r.Intn(2) == 1 {
			h.Lines[i] = Solid
		} else {
			h.Lines[i] = Broken
		}
	}
	return h
}

// DrawLucky generates a random hexagram.
func DrawLucky() Hexagram {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var h Hexagram
	for i := 0; i < 6; i++ {
		if r.Intn(2) == 1 {
			h.Lines[i] = Solid
		} else {
			h.Lines[i] = Broken
		}
	}
	return h
}
