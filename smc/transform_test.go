package smc

import (
	"testing"
)

func TestIsNotSpecialCode(t *testing.T) {
	// Case: composites have leading bits representing the instructions.
	// Composites 00000000 = 00, 01000000 = 64, 10000000 = 128, 11000000 = 192.
	c := []uint8{0b00000000, 0b01000000, 0b10000000, 0b11000000}
	// Expected results.
	e := []bool{true, true, true, false}
	// Check results.
	for i := 0; i < len(c); i++ {
		if b := isNotSpecialCode(c[i]); b != e[i] {
			t.Errorf("expected %t for input %d but got %t", e[i], c[i], b)
		}
	}
}

func TestSubstring(t *testing.T) {
	// Case: isolate from the composite instruction and the {c}_{i} substring given by i.
	// We can only have 4 substrings from indexes 0, 1, 2, 3.
	// Composite.
	c := uint8(0b00011011)
	// Expected results (i.e., 0, 1, 2, 3).
	// The first is the instruction.
	e := []uint8{0b00000000, 0b00000001, 0b00000010, 0b00000011}
	// Check results.
	for i := 0; i < len(e); i++ {
		if s := substring(c, i); s != e[i] {
			t.Errorf("expected %d for input %d for index %d but got %d", e[i], c, i, s)
		}
	}
}

func TestTransform(t *testing.T) {
	// Case: complete cycle of transformations for composite "CAT".
	// Expected results: "CAT" -> "GTT" -> "CTA" -> "CAT" (original form).
	e := []uint8{0b00010011, 0b01101111, 0b10011100, 0b00010011}
	// Check results.
	for i := 0; i < len(e)-1; i++ {
		if c := transform(e[i]); c != e[i+1] {
			t.Errorf("expected %08b as transformation of %08b but got %08b", e[i+1], e[i], c)
		}
	}

	// Case: special code (i.e., leading {11}_{2} instruction) when transformed should return unaltered.
	// Special composite.
	c := uint8(0b11010110)
	// Check result.
	if ct := transform(c); ct != c {
		t.Errorf("transformation of %08b should be %08b", c, ct)
	}
}

func TestInverseTransform(t *testing.T) {
	// Case: complete cycle of detransformations for composites "GTT", "CTA" and "CAT" (original form).
	// Composites with equivalent instructions.
	c := []uint8{0b01101111, 0b10011100, 0b00010011}
	// Expected result is "CAT" (original form).
	e := uint8(0b00010011)
	// Check results.
	for i := 0; i < len(c); i++ {
		if ct := inverseTransform(c[i]); ct != e {
			t.Errorf("expected %08b as transformation of %08b but got %08b", e, c[i], ct)
		}
	}

	// Case: special code (i.e., leading {11}_{2} instruction) when detransformed should return unaltered.
	// Special composite.
	sc := uint8(0b11010110)
	// Check result.
	if ct := transform(sc); ct != sc {
		t.Errorf("transformation of %08b should be %08b", sc, ct)
	}
}

func TestMap(t *testing.T) {
	// Case: map a set of 3 characters into a vector of 3 transforms and 2 special codes.
	// Composites.
	c := []uint8{0b00011011, 0b11000111, 0b11010101}
	// Expected result [[⟨ESC⟩, m, ¶], [Ç], [Õ]].
	// The first vector is the set of all transformations for 00011011.
	// The second and third vectors are special codes that will not be transformed.
	e := [][]uint8{{0b00011011, 0b01101101, 0b10110110},
		{0b11000111, 0b11000111, 0b11000111},
		{0b11010101, 0b11010101, 0b11010101}}
	// Check results.
	for i := 0; i < len(c); i++ {
		// For each value in the composite vector there is an expected vector result.
		v := Map(c[i])
		for j := 0; j < len(e[i]); j++ {
			if v[j] != e[i][j] {
				t.Errorf("expected %08b as a map for value %08b but got %08b", e[i][j], c[i], v[j])
			}
		}
	}
}
