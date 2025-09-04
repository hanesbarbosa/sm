package smc

import "bytes"

// Glossary is a structure that holds all the allowed symbols and
// the mapping between ASCII characters and codes.
// Tokens will map one or more characters into codes.
// Codes will map integers back into the original characters.
type Glossary struct {
	alphabet         *Alphabet
	code             map[string]uint8
	token            map[uint8]string
	specialCodeCount uint8 // Number of special codes available.
	nextSpecialCode  uint8 // Next index available.
}

func NewGlossary(a *Alphabet) (*Glossary, error) {
	// New empty dictionary.
	g := new(Glossary)
	// Initiate maps with the size for the whole set of codes (i.e., 2^(2+6) = 256).
	size := 1 << (BitsPerInstruction + BitsPerCode)
	g.code = make(map[string]uint8, size)
	g.token = make(map[uint8]string, size)
	// Store given alphabet.
	g.alphabet = a
	// Define remaining number of special codes.
	g.specialCodeCount = MaximumSizeOmegaAlphabet
	// Define the next available index for special codes (i.e., 0b11000000 = 192).
	g.nextSpecialCode = uint8(size - (1 << BitsPerCode))
	// Create all codes for Σ¹.
	g.setSigmaCodes()
	// Create all codes for Λ³.

	// Generate special codes based on user input.
	// Special codes start at 11000000 = 192 and cannot pass 11111111 = 255.

	return g, nil
}

func (g *Glossary) setSigmaCodes() {
	// Define a code for each individual character.
	for _, c := range g.alphabet.sigma() {
		g.code[string(c)] = g.nextSpecialCode
		// Increase index.
		g.nextSpecialCode += 1
	}
}

// Code returns the ASCII code for the given string.
func (g *Glossary) Code(t string) uint8 {
	return g.code[t]
}

func (g *Glossary) parse(t string) uint8 {
	// Get the 4 bases from Lambda.
	l := g.alphabet.lambda
	// Find the indexes for each base in the token.
	b1 := uint8(bytes.IndexByte(l, t[0]))
	b1 <<= 4
	b2 := uint8(bytes.IndexByte(l, t[1]))
	b2 <<= 2
	b3 := uint8(bytes.IndexByte(l, t[2]))
	// Build the code wiht the instruction and 3 bases.
	c := uint8(0)
	return c ^ b1 ^ b2 ^ b3
}
