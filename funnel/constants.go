package funnel

const (
	BitsPerInstruction  = 2
	BitsPerCode         = 6
	TotalAlphabet       = 1 << BitsPerCode
	TotalLambdaAlphabet = 1 << BitsPerInstruction
	TotalOmegaAlphabet  = TotalAlphabet - TotalLambdaAlphabet
)
