package funnel

import (
	"testing"
)

func TestGlossary(t *testing.T) {
	// Case: a new glossary has unique codes for each Σ¹ character.
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

	// Case: a new glossary has unique codes for all Λ³ tokens.

	// Case: a new glossary has unique special codes for tokens not in Λ³.
	// Possible combinations can be:
	// 1 - 3 characters not in Λ³: Ω³, Λ¹Ω², Ω²Λ¹, Λ²Ω¹, Ω¹Λ², Ω¹Λ¹Ω¹, Λ¹Ω¹Λ¹.
	// 2 - 2 characters: Ω², Λ², Λ¹Ω¹, Ω¹Λ¹.
	// Special codes start at 11000000 = 192 and cannot pass 11111111 = 255.
}
