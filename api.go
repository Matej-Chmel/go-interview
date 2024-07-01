package gointerview

import (
	"io"

	ite "github.com/Matej-Chmel/go-interview/internal"
)

// Class for single input problems.
// Delegates all implementation to Interview2
// with a dummy type for the second input.
type Interview[I any, O any] struct {
	iv Interview2[I, int, O]
}

// Constructs an Interview object
func NewInterview[I any, O any]() Interview[I, O] {
	return Interview[I, O]{
		iv: newInterview2Impl[I, int, O](true),
	}
}

// Adds one test case
func (i *Interview[I, O]) AddCase(input I, expected O) {
	i.iv.AddCase(input, 0, expected)
}

// Adds multiple test cases
func (i *Interview[I, O]) AddCases(input []I, expected []O) {
	i.iv.AddCasesSlice(input, []int{}, expected, 0, -1)
}

// Adds multiple test cases in range [begin, end)
func (i *Interview[I, O]) AddCasesSlice(input []I, expected []O, begin, end int) {
	i.iv.AddCasesSlice(input, []int{}, expected, begin, end)
}

// Adds one solution function
func (i *Interview[I, O]) AddSolution(s func(I) O) {
	i.iv.solutions1[ite.GetFunctionName(s)] = s
}

// Adds multiple solution functions
func (i *Interview[I, O]) AddSolutions(s ...func(I) O) {
	for _, f := range s {
		i.AddSolution(f)
	}
}

// Runs all solutions against all test cases
// and compiles the output into a single string
func (i *Interview[I, O]) AllSolutionsToString() string {
	return i.iv.AllSolutionsToString()
}

// Runs all solutions against all test cases
// and prints the output to the standard output
func (i *Interview[I, O]) Print() error {
	return i.iv.Print()
}

// Reads one case from relative paths for input and output
func (i *Interview[I, O]) ReadCase(inputRelPath, expectedRelPath string) {
	i.iv.ReadCase(inputRelPath, "", expectedRelPath)
}

// Reads multiple cases from relative paths for input and output
func (i *Interview[I, O]) ReadCases(inputRelPath, expectedRelPath string) {
	i.iv.ReadCases(inputRelPath, "", expectedRelPath)
}

// Reads multiple cases from relative paths for input and output.
// Only the cases in range [begin, end) are added.
func (i *Interview[I, O]) ReadCasesSlice(
	inputRelPath, expectedRelPath string, begin, end int,
) {
	i.iv.ReadCasesSlice(inputRelPath, "", expectedRelPath, begin, end)
}

// Runs one solution function against all test cases.
// If function cannot be found, an error is returned.
func (i *Interview[I, O]) RunSolution(name string) (ite.Receipt, error) {
	return i.iv.RunSolution(name)
}

// Runs all solutions against all test cases
func (i *Interview[I, O]) RunAllSolutions() ite.ReceiptSlice {
	return i.iv.RunAllSolutions()
}

// Runs all solutions against all test cases
// and writes the results into a writer w
func (i *Interview[I, O]) WriteAllSolutions(w io.Writer) error {
	return i.iv.WriteAllSolutions(w)
}
