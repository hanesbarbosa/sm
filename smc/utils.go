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
