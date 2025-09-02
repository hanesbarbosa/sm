package smc

import (
	"strings"
)

// Alphabet is the structure that stores all allowed bases.
// Lambda is the set of the 4 more frequent bases.
// Omega is the set of the ambiguous bases, errors, etc.
type Alphabet struct {
	lambda []byte
	omega  []byte
}

// NewAlphabet instantiates a new Alphabet structure.
func NewAlphabet(lambda, omega string) (*Alphabet, error) {
	// New struct for the alphabet characters.
	a := new(Alphabet)
	// Uppercase the input alphabets.
	l, o := strings.ToUpper(lambda), strings.ToUpper(omega)
	// Store.
	a.lambda = []byte(l)
	a.omega = []byte(o)
	// Validate fields.
	err := a.validate()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Alphabet) validate() error {
	// Return variable
	var err error
	// Validate Lambda.
	err = a.validateLambda()
	if err != nil {
		return err
	}
	// Validate Omega.
	err = a.validateOmega()
	if err != nil {
		return err
	}

	return err
}

func (a *Alphabet) validateLambda() error {
	// Check maximum number of bases.
	if len(a.lambda) != TotalSizeLambdaAlphabet {
		return ErrorAlphabetLambdaSize
	}
	// Should not have Omega bases.
	if a.haveCommonCharacters() {
		return ErrorAlphabetCommonCharacter
	}

	return nil
}

func (a *Alphabet) validateOmega() error {
	// Cannot be empty.
	if len(a.omega) == 0 {
		return ErrorAlphabetEmptyOmega
	}
	// Check maximum number of bases.
	if len(a.omega) > MaximumSizeOmegaAlphabet {
		return ErrorAlphabetOmegaSize
	}

	return nil
}

func (a *Alphabet) haveCommonCharacters() bool {
	// For each character in the first alphabet check the existence in the second.
	for i := 0; i < len(a.lambda); i++ {
		for j := 0; j < len(a.omega); j++ {
			if a.lambda[i] == a.omega[j] {
				return true
			}
		}
	}

	return false
}

func (a *Alphabet) sigma() []byte {
	return append(a.lambda, a.omega...)
}
