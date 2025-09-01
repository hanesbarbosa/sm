package funnel

import (
	"testing"
)

func TestAlphabet(t *testing.T) {
	// Case: Lambda without 4 bases returns an error.
	// Covers the case for "empty Lambda returns an error".
	// Lambda and Omega alphabets.
	l, o := "ACG", "BDHKMNRSUVWY-"
	// New Alphabet.
	_, err := NewAlphabet(l, o)
	// Check error.
	if err == nil || err != ErrorAlphabetLambdaSize {
		t.Errorf("Lambda without %d bases should return the \"%s\" error", TotalSizeLambdaAlphabet, ErrorAlphabetLambdaSize.Error())
	}

	// Case: empty Omega returns an error.
	// Lambda and Omega alphabets.
	l, o = "ACGT", ""
	// New Alphabet.
	_, err = NewAlphabet(l, o)
	// Check error.
	if err == nil || err != ErrorAlphabetEmptyOmega {
		t.Errorf("empty Omega should return the \"%s\" error", ErrorAlphabetEmptyOmega.Error())
	}

	// Case: Lambda or Omega with repeated characters returns an error.
	// Covers the cases for:
	// 1 - Lambda with an Omega character returns an error.
	// Lambda and Omega alphabets.
	l, o = "ACG-", "BDHKMNRSUVWY-"
	// New Alphabet.
	_, err = NewAlphabet(l, o)
	// Check error.
	if err == nil || err != ErrorAlphabetCommonCharacter {
		t.Errorf("Omega characters in Lambda should return the \"%s\" error", ErrorAlphabetCommonCharacter.Error())
	}
	// 2 - Omega with an Lambda character returns an error.
	// Lambda and Omega alphabets.
	l, o = "ACGT", "BDHKMNRSUVWY-A"
	// New Alphabet.
	_, err = NewAlphabet(l, o)
	// Check error.
	if err == nil || err != ErrorAlphabetCommonCharacter {
		t.Errorf("Lambda characters in Omega should return the \"%s\" error", ErrorAlphabetCommonCharacter.Error())
	}

	// Case: when Lambda and Omega together exceed the maximum number of characters it returns an error.
	// Lambda and Omega alphabets.
	l, o = "ACGT", "!#$%&'()*+,-./0123456789:;<=>?@0123456789BDEFHIJKLMNOPQRSUVWXYZ[\\]^`"
	// New Alphabet.
	_, err = NewAlphabet(l, o)
	// Check error.
	if err == nil || err != ErrorAlphabetOmegaSize {
		t.Errorf("alphabet with more than %d characters should return the \"%s\" error", TotalSizeAlphabet, ErrorAlphabetOmegaSize.Error())
	}

	// Case: Sigma is comprised by both Lambda and Omega.
	// Lambda and Omega alphabets.
	l, o = "ACGT", "BDHKMNRSUVWY-"
	// New Alphabet.
	a, err := NewAlphabet(l, o)
	if err != nil {
		t.Error(err.Error())
	}
	// Check if |Σ| = |Λ| + |Ω|.
	// Sigma.
	s := a.sigma()
	// Expected length.
	el := len(l + o)
	if sl := len(s); sl != el {
		t.Errorf("expected length %d for Sigma but got %d", el, sl)
	}
	// Check if Σ = Λ U Ω.
	es := []byte(l)
	es = append(es, []byte(o)...)
	// Check each element from Sigma.
	if notEqual(es, s) {
		t.Error("Sigma alphabet is not comprised by Lambda and Omega")
	}
}
