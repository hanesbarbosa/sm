package smc

const (
	BitsPerInstruction       = 2
	BitsPerCode              = 6
	BASES                    = 1 << BitsPerInstruction
	TotalSpecialCodes        = 1 << BitsPerCode
	TotalSizeLambdaAlphabet  = 1 << BitsPerInstruction
	MaximumSizeOmegaAlphabet = TotalSpecialCodes - TotalSizeLambdaAlphabet
	TotalTripletsCodes       = 1 << BitsPerCode
)
