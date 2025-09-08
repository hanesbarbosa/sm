package smc

import (
	"bytes"
)

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
	// Define the next available index for special codes starting at 0b11000000 = 192.
	g.nextSpecialCode = uint8(size - (1 << BitsPerCode))
	// Create all codes for Λ³ (i.e., DNA most frequent bases with instructions 00, 01 and 10).
	g.setLambdaCodes()
	// Create all codes for Σ¹ (i.e., individual characters for final truncation).
	g.setSigmaCodes()
	// Generate special codes based on user input
	// to populate the Priority Statistical Model (PSM).
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

func (g *Glossary) setLambdaCodes() {
	// Define all Lambda triplets transforms.
	tc := tripletsCombination(g)
	// Generate the code mapping.
	for i := 0; i < len(tc); i++ {
		g.code[tc[i]] = uint8(i)
	}
}

// Code returns the ASCII code for the given string.
func (g *Glossary) Code(t string) uint8 {
	return g.code[t]
}

func (g *Glossary) tripletsToCode(t string) (uint8, error) {
	// Check number of characters.
	if g.isNotATriplet(t) {
		return 0, ErrorNotATriplet
	}
	// Find the indexes for each base in the token.
	b1 := uint8(bytes.IndexByte(g.alphabet.lambda, t[0]))
	b1 <<= 4
	b2 := uint8(bytes.IndexByte(g.alphabet.lambda, t[1]))
	b2 <<= 2
	b3 := uint8(bytes.IndexByte(g.alphabet.lambda, t[2]))
	// Build the code wiht the initial instruction and 3 bases.
	c := uint8(0)
	return c ^ b1 ^ b2 ^ b3, nil
}

func (g *Glossary) isNotATriplet(t string) bool {
	// TODO: use a constant.
	if len(t) != 3 {
		return true
	}
	// Check allowed symbols.
	for _, v := range t {
		// Countdown for bases.
		count := BASES
		for i := 0; i < len(g.alphabet.lambda); i++ {
			if uint8(v) == g.alphabet.lambda[i] {
				break
			}
			count -= 1
		}
		// Check if symbol was found.
		if count <= 0 {
			return true
		}
	}

	return false
}

func (g *Glossary) codeToTriplets(c uint8) string {
	// Get codes for bases besides the initial instruction.
	c1, c2, c3 := substring(c, 1), substring(c, 2), substring(c, 3)
	// Build the resulting triplet.
	return string(g.alphabet.lambda[c1]) + string(g.alphabet.lambda[c2]) + string(g.alphabet.lambda[c3])
}
