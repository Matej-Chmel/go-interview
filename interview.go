package gointerview

import (
	"errors"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
)

// Stores cases and map of solutions.
type Interview[I any, O any] struct {
	Cases     []Case[I, O]
	Solutions map[string]func(I) O
}

// Interview constructor.
func NewInterview[I any, O any]() *Interview[I, O] {
	return &Interview[I, O]{Solutions: make(map[string]func(I) O)}
}

// Adds case from input and expected output.
func (i *Interview[I, O]) AddCase(input I, output O) *Interview[I, O] {
	i.Cases = append(i.Cases, NewCase(input, output))
	return i
}

// Adds a solution function.
func (i *Interview[I, O]) AddSolution(s func(I) O) *Interview[I, O] {
	i.Solutions[getFunctionName(s)] = s
	return i
}

// Adds multiple solution functions.
func (i *Interview[I, O]) AddSolutions(solutions ...func(I) O) *Interview[I, O] {
	for _, s := range solutions {
		i.Solutions[getFunctionName(s)] = s
	}
	return i
}

// Runs all cases with all solutions and prints the reports.
func (i *Interview[I, O]) Print() {
	i.Write(os.Stdout)
}

// Runs all cases with all solutions and returns the reports.
func (i *Interview[I, O]) RunAllSolutions() []Report[I, O] {
	slice := []Report[I, O]{}

	for name, f := range i.Solutions {
		slice = append(slice, i.runSolutionImpl(name, f)...)
	}

	return slice
}

// Runs all cases with one solution and returns the reports.
func (i *Interview[I, O]) RunSolution(name string) ([]Report[I, O], error) {
	f, status := i.Solutions[name]

	if status {
		return i.runSolutionImpl(name, f), nil
	}

	return make([]Report[I, O], 0), errors.New("unknown solution")
}

// Private function for running a solution.
func (i *Interview[I, O]) runSolutionImpl(name string, f func(I) O) []Report[I, O] {
	slice := []Report[I, O]{}

	for _, aCase := range i.Cases {
		res := aCase.RunSolution(name, f)
		slice = append(slice, res)
	}

	return slice
}

// Runs all cases with all solutions and returns the reports as string.
func (i Interview[I, O]) String() string {
	return sliceToString(i.RunAllSolutions(), "\n\n")
}

// Runs all cases with all solutions and writes the reports to writer w.
func (i *Interview[I, O]) Write(w io.Writer) (int, error) {
	return io.WriteString(w, i.String())
}

// Returns name of function f.
func getFunctionName(f interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	tokens := strings.Split(name, ".")
	return tokens[len(tokens)-1]
}
