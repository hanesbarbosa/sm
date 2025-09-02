package funnel

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
	// Initiate maps with the size for the whole set of codes (i.e., 2^8 = 256).
	l := 1 << (BitsPerInstruction + BitsPerCode)
	g.code = make(map[string]uint8, l)
	g.token = make(map[uint8]string, l)
	// Store given alphabet.
	g.alphabet = a
	// Define remaining number of special codes.
	g.specialCodeCount = MaximumSizeOmegaAlphabet
	// Define the next available index for special codes.
	g.nextSpecialCode = 0b11000000
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
