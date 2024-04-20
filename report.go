package gointerview

import (
	"fmt"
	"io"
	"strings"
)

// Stores the found mismatches after calling a solution function
// with the input from the specified Case.
type Report[I any, O any] struct {
	*Case[I, O]
	FunctionName string
	Mismatches   []Mismatch
}

// Adds a new mismatch to the report.
func (c *Report[I, O]) AddMismatch(field, left, right string) {
	c.Mismatches = append(c.Mismatches, Mismatch{field, left, right})
}

// Returns a mismatch at index i.
func (c *Report[I, O]) Get(i int) Mismatch {
	return c.Mismatches[i]
}

// Returns number of stored mismatches.
func (c *Report[I, O]) Len() int {
	return len(c.Mismatches)
}

// Converts the report to a string.
func (c Report[I, O]) String() string {
	var builder strings.Builder
	builder.WriteString(
		fmt.Sprintf(
			"%s(%s) -> %s\n",
			c.FunctionName, c.InputToString(), c.ExpectedToString()))

	if len(c.Mismatches) == 0 {
		builder.WriteString("OK!")
	} else {
		sliceToBuilder(&builder, c.Mismatches, "\n")
	}

	return builder.String()
}

// Writes the report to a writer.
func (c *Report[I, O]) Write(w io.Writer) (int, error) {
	return io.WriteString(w, c.String())
}
