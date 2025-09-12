package smc

// PSM is the structure representing the Priority Statistical Model (PSM).
//
// It has a list of characters showing which transformation
// should be picked for a given triplet. For instance, the original triplet
// "AAA" = 00000000 has 3 transformations: 00000000 = 0, 01000000 = 64 and 10000000 = 128.
// In this case the PSM will have only one of these values at the position 0 of the vector,
// meaning that this is the transformation chosen for "AAA". Position 1 of the vector will
// have the choice for transformations of "AAC" = 00000001 and so on.
//
// The preferred codes can be chosen by statistical analysis. For instance, the frequency
// of characters from comments' lines can build a model. Another example could be methods for checking
// the distance between characters and chosing the closest ones. All in all, The choice for the criteria
// depends of which decision will be made about the second-step compression.
type PSM struct {
	code [TotalTripletsCodes]uint8
}

// Instantiate a new model with a populated hierarchy of symbols.
func NewPSM() *PSM {
	// New model.
	m := new(PSM)
	// Initialize values for the model.
	m.initialize()

	return m
}

func (m *PSM) initialize() {
	// Populate the hierarchy.
	m.code = [TotalTripletsCodes]uint8{}
	// TODO: change for statistical analysis of comments.
	// Just for testing, we initially define the first transform
	// as the expected one.
	for v := range TotalTripletsCodes {
		m.code[v] = uint8(v)
	}
}
