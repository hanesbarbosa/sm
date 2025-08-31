package funnel

import "errors"

var (
	ErrorAlphabetLambdaSize      = errors.New("lambda does not have 4 bases")
	ErrorAlphabetEmptyOmega      = errors.New("omega alphabet is empty")
	ErrorAlphabetCommonCharacter = errors.New("lambda and omega alphabets cannot have common characters")
	ErrorAlphabetOmegaSize       = errors.New("omega alphabet exceeds maximum number of characters")
)
