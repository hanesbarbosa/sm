package smc

func isNotSpecialCode(c uint8) bool {
	// Isolate first bit.
	b1 := c >> 7
	// Isolate second bit.
	b2 := c << 1
	b2 >>= 7
	// NAND.
	return !(b1 == 1 && b2 == 1)
}

func substring(c uint8, i int) uint8 {
	// Left offset.
	c <<= (i % BASES) * 2
	// Right offset.
	return c >> 6
}

func transform(c uint8) uint8 {
	// Only transform special codes.
	if isNotSpecialCode(c) {
		// Instruction.
		i := substring(c, 0)
		// Substrings.
		c1, c2, c3 := substring(c, 1), substring(c, 2), substring(c, 3)
		// Transformations.
		c1Prime, c2Prime, c3Prime := c1^c3, c3, c2^c3
		// Iterate instruction.
		i = (i + 1) % (BASES - 1)
		// Pack results into a composite.
		return (i << 6) ^ (c1Prime << 4) ^ (c2Prime << 2) ^ c3Prime
	} else {
		return c
	}
}

func inverseTransform(c uint8) uint8 {
	// Only detransform special codes.
	if isNotSpecialCode(c) {
		// Instruction.
		for substring(c, 0) != 0 {
			c = transform(c)
		}
	}
	// Return final composite.
	return c
}

// Map generates all possible transformations for the given composite.
// If the composite is a special code, then a set of repeated characters will be given.
func Map(c uint8) []uint8 {
	// Output vector with initial state.
	v := []uint8{c}
	// Create next 2 transformations based on the initial composite.
	for i := 0; i < 2; i++ {
		v = append(v, transform(v[i]))
	}
	return v
}
