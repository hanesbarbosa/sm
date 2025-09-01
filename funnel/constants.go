package funnel

const (
	BitsPerInstruction       = 2
	BitsPerCode              = 6
	TotalSizeAlphabet        = 1 << BitsPerCode
	TotalSizeLambdaAlphabet  = 1 << BitsPerInstruction
	MaximumSizeOmegaAlphabet = TotalSizeAlphabet - TotalSizeLambdaAlphabet
)
