package gointerview

// Stores input and expected output for one run.
type Case[I any, O any] struct {
	Input    I
	Expected O
}

// Constructor
func NewCase[I any, O any](input I, expected O) Case[I, O] {
	return Case[I, O]{Input: deepCopy(input), Expected: deepCopy(expected)}
}

// Returns expected output as string.
func (c *Case[I, O]) ExpectedToString() string {
	return anyToString(c.Expected)
}

// Returns input as string.
func (c *Case[I, O]) InputToString() string {
	return anyToString(c.Input)
}

// Runs a solution with the input and returns a report of mismatches
// between actual and expected output.
func (c *Case[I, O]) RunSolution(fName string, s func(I) O) Report[I, O] {
	input := deepCopy(c.Input)
	actual := s(input)
	mismatches := findMismatches(actual, c.Expected)
	return Report[I, O]{c, fName, mismatches}
}
