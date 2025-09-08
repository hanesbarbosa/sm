package smc

import (
	"testing"
)

func TestNewGlossary(t *testing.T) {
	// Case: a new glossary has codes for each Σ¹ character.
	// Lambda and Omega alphabets.
	l, o := "ACGT", "BDHKMNRSUVWY->\n"
	// New Alphabet.
	a, err := NewAlphabet(l, o)
	if err != nil {
		t.Error(err.Error())
	}
	// Check the existence of Σ¹ characters.
	s := a.sigma()
	// New glossary.
	g, err := NewGlossary(a)
	if err != nil {
		t.Error(err.Error())
	}
	// Initial value for special codes.
	// Special codes start at 11000000 = 192.
	sc := uint8(0b11000000)
	// Check if there is a special code for each letter.
	for _, c := range s {
		if g.Code(string(c)) < sc {
			t.Errorf("special code for character %d must be equal or higher than %d", c, sc)
		}
	}
}

func TestTripletsToCode(t *testing.T) {
	// Case: three letters from Lambda (i.e., triplets) are converted into the right codes.
	// Lambda and Omega alphabets.
	l, o := "ACGT", "BDHKMNRSUVWY->\n"
	// New Alphabet.
	a, err := NewAlphabet(l, o)
	if err != nil {
		t.Error(err.Error())
	}
	// New glossary.
	g, err := NewGlossary(a)
	if err != nil {
		t.Error(err.Error())
	}
	// Character tokens to be converted into codes.
	// A = 00, C = 01, G = 10, T = 11.
	s := []string{"AAA", "ACG", "AGT", "CGT", "TTA"}
	// Expected codes.
	ec := []uint8{0b00000000, 0b00000110, 0b00001011, 0b00011011, 0b00111100}
	// Check parsing.
	for i := 0; i < len(s); i++ {
		c, err := g.tripletsToCode(s[i])
		if err != nil {
			t.Error(err.Error())
		}
		if c != ec[i] {
			t.Errorf("expected code %08b but got %08b", ec[i], c)
		}
	}

	// Case: only triplets in Λ³ are accepted.
	s = []string{"", "A", "AA", "AZT", "WER", "ACGT"}
	// Check parsing.
	for i := 0; i < len(s); i++ {
		_, err = g.tripletsToCode(s[i])
		// Check error.
		if err == nil || err != ErrorNotATriplet {
			t.Errorf("string \"%s\" with at leas one symbol not in Lambda should return the \"%s\" error", s[i], ErrorNotATriplet.Error())
		}
	}
}

func TestCodeToTriplets(t *testing.T) {
	// Case: codes from 0 to 191 are a result of Lambda combinations of 3 letters.
	t.Skip()
}

// TODO:

// Case: a new glossary has unique special codes for tokens not in Λ³.
// Possible combinations can be:
// 1 - 3 characters not in Λ³: Ω³, Λ¹Ω², Ω²Λ¹, Λ²Ω¹, Ω¹Λ², Ω¹Λ¹Ω¹, Λ¹Ω¹Λ¹.
// 2 - 2 characters: Ω², Λ², Λ¹Ω¹, Ω¹Λ¹.
// Special codes start at 11000000 = 192 and cannot pass 11111111 = 255.

// Case: trying to store a code higher than 11111111 = 255 raises an error.
// This case is intended for checking the user input when trying to fill the
// Priority Statistical Model (PSM).
