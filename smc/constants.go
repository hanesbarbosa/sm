package smc

const (
	BitsPerInstruction       = 2
	BitsPerCode              = 6
	BASES                    = 1 << BitsPerInstruction
	TotalSizeAlphabet        = 1 << BitsPerCode
	TotalSizeLambdaAlphabet  = 1 << BitsPerInstruction
	MaximumSizeOmegaAlphabet = TotalSizeAlphabet - TotalSizeLambdaAlphabet
)
