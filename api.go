package gointerview

import (
	"io"

	ite "github.com/Matej-Chmel/go-interview/internal"
)

// Class for single input problems.
// Delegates all implementation to Interview2
// with a dummy type for the second input.
type Interview[I any, O any] struct {
	ite.EmbeddedOptions
	iv Interview2[I, int, O]
}

// Constructs an Interview object
func NewInterview[I any, O any]() Interview[I, O] {
	options := ite.NewEmbeddedOptions()

	return Interview[I, O]{
		EmbeddedOptions: options,
		iv:              newInterview2Impl[I, int, O](true, &options),
	}
}

// Adds one test case
func (iv *Interview[I, O]) AddCase(input I, expected O) {
	iv.iv.AddCase(input, 0, expected)
}

// Converts strings to a byte or rune slices and attempts
// to add those slices as a new test case.
// Panics if any string cannot be converted to its target type.
func (iv *Interview[I, O]) AddCaseString(input string, expected string) {
	iv.iv.AddCaseString(input, "", expected)
}

// Adds multiple test cases
func (iv *Interview[I, O]) AddCases(input []I, expected []O) {
	iv.iv.AddCasesSlice(input, []int{}, expected, 0, -1)
}

// Adds multiple test cases in range [begin, end)
func (iv *Interview[I, O]) AddCasesSlice(input []I, expected []O, begin, end int) {
	iv.iv.AddCasesSlice(input, []int{}, expected, begin, end)
}

// Adds one solution function
func (iv *Interview[I, O]) AddSolution(s func(I) O) {
	iv.iv.solutions1[ite.GetFunctionName(s)] = s
}

// Adds multiple solution functions
func (iv *Interview[I, O]) AddSolutions(s ...func(I) O) {
	for _, f := range s {
		iv.AddSolution(f)
	}
}

// Runs all solutions against all test cases
// and compiles the output into a single string
func (iv *Interview[I, O]) AllSolutionsToString() string {
	return iv.iv.AllSolutionsToString()
}

// Runs all solutions against all test cases
// and prints the output to the standard output
func (iv *Interview[I, O]) Print() error {
	return iv.iv.Print()
}

// Reads one case from relative paths for input and output
func (iv *Interview[I, O]) ReadCase(inputRelPath, expectedRelPath string) {
	iv.iv.ReadCase(inputRelPath, "", expectedRelPath)
}

// Reads multiple cases from relative paths for input and output
func (iv *Interview[I, O]) ReadCases(inputRelPath, expectedRelPath string) {
	iv.iv.ReadCases(inputRelPath, "", expectedRelPath)
}

// Reads multiple cases from relative paths for input and output.
// Only the cases in range [begin, end) are added.
func (iv *Interview[I, O]) ReadCasesSlice(
	inputRelPath, expectedRelPath string, begin, end int,
) {
	iv.iv.ReadCasesSlice(inputRelPath, "", expectedRelPath, begin, end)
}

// Runs one solution function against all test cases.
// If function cannot be found, an error is returned.
func (iv *Interview[I, O]) RunSolution(name string) (ite.Receipt, error) {
	return iv.iv.RunSolution(name)
}

// Runs all solutions against all test cases
func (iv *Interview[I, O]) RunAllSolutions() ite.ReceiptSlice {
	return iv.iv.RunAllSolutions()
}

// Runs all solutions against all test cases
// and writes the results into a writer w
func (iv *Interview[I, O]) WriteAllSolutions(w io.Writer) error {
	return iv.iv.WriteAllSolutions(w)
}
