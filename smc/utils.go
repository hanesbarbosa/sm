package smc

func notEqual(b1, b2 []uint8) bool {
	// Check size.
	if len(b1) != len(b2) {
		return true
	}
	// Check members.
	for i, v := range b1 {
		if v != b2[i] {
			return true
		}
	}

	return false
}

// TODO: use a constant for the size.
func tripletsCombination(g *Glossary) [192]string {
	// Output string for triplets.
	s := [192]string{}
	// Iterate for all possible triplets.
	for i := 0; i < 192; i++ {
		s[i] = g.codeToTriplets(uint8(i))
	}

	return s
}
