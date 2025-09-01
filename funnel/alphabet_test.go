package funnel

import (
	"testing"
)

func TestAlphabet(t *testing.T) {
	// Case: lambda without 4 bases returns an error.
	// Covers the case for "empty lambda returns an error".
	// Lambda and omega alphabets.
	l, o := "ACG", "BDHKMNRSUVWY-"
	// New Alphabet.
	_, err := NewAlphabet(l, o)
	// Check error.
	if err == nil || err != ErrorAlphabetLambdaSize {
		t.Errorf("lambda without %d bases should return the \"%s\" error", TotalSizeLambdaAlphabet, ErrorAlphabetLambdaSize.Error())
	}

	// Case: empty omega returns an error.
	// Lambda and omega alphabets.
	l, o = "ACGT", ""
	// New Alphabet.
	_, err = NewAlphabet(l, o)
	// Check error.
	if err == nil || err != ErrorAlphabetEmptyOmega {
		t.Errorf("empty omega should return the \"%s\" error", ErrorAlphabetEmptyOmega.Error())
	}

	// Case: lambda or omega with repeated characters returns an error.
	// Covers the cases for
	// 1 - Lambda with an omega character returns an error.
	// Lambda and omega alphabets.
	l, o = "ACG-", "BDHKMNRSUVWY-"
	// New Alphabet.
	_, err = NewAlphabet(l, o)
	// Check error.
	if err == nil || err != ErrorAlphabetCommonCharacter {
		t.Errorf("omega characters in lambda should return the \"%s\" error", ErrorAlphabetCommonCharacter.Error())
	}
	// 2 - Omega with an lambda character returns an error.
	// Lambda and omega alphabets.
	l, o = "ACGT", "BDHKMNRSUVWY-A"
	// New Alphabet.
	_, err = NewAlphabet(l, o)
	// Check error.
	if err == nil || err != ErrorAlphabetCommonCharacter {
		t.Errorf("lambda characters in omega should return the \"%s\" error", ErrorAlphabetCommonCharacter.Error())
	}

	// Case: when lambda and omega together exceed the maximum number of characters it returns an error.
	// Lambda and omega alphabets.
	l, o = "ACGT", "!#$%&'()*+,-./0123456789:;<=>?@0123456789BDEFHIJKLMNOPQRSUVWXYZ[\\]^`"
	// New Alphabet.
	_, err = NewAlphabet(l, o)
	// Check error.
	if err == nil || err != ErrorAlphabetOmegaSize {
		t.Errorf("alphabet with more than %d characters should return the \"%s\" error", TotalSizeAlphabet, ErrorAlphabetOmegaSize.Error())
	}
}
