package internal

import (
	"strings"
)

// Output information for all test cases under one solution name
type Receipt struct {
	Lines []*ReceiptLine
	Name  string
}

// Writes itself to builder
// Each multi-line test case is separated by double newline
func (s *Receipt) ContinueBuild(builder *strings.Builder) {
	builder.WriteString(s.Name)
	builder.WriteRune('\n')
	builder.WriteString(strings.Repeat("=", len(s.Name)))
	isMultiLine := false

	for _, l := range s.Lines {
		builder.WriteRune('\n')

		if isMultiLine {
			builder.WriteRune('\n')
		}

		isMultiLine = l.ContinueBuild(builder)
	}
}

// Output information about a single test case
type ReceiptLine struct {
	Actual   string
	Expected string
	Input    string
	Input2   *string
}

// Constructs ReceiptLine for a single input problem
func NewReceiptLine(input, actual, expected string) *ReceiptLine {
	return NewReceiptLineImpl(actual, expected, input, nil)
}

// Constructs ReceiptLine for a two input problem
func NewReceiptLine2(input1, input2, actual, expected string) *ReceiptLine {
	return NewReceiptLineImpl(actual, expected, input1, &input2)
}

// Internal constructor for ReceiptLine
func NewReceiptLineImpl(actual, expected, input1 string, input2 *string) *ReceiptLine {
	return &ReceiptLine{Actual: actual, Expected: expected, Input: input1, Input2: input2}
}

// Writes itself to builder using a new IteratorCollection
// Returns a flag indicating whether a newline character was written
func (r *ReceiptLine) ContinueBuild(builder *strings.Builder) bool {
	col := NewIteratorCollection(r.Actual, r.Expected, r.Input, r.Input2)
	return WriteCollection(builder, col)
}

// Returns a flag indicating whether two ReceiptLines match
func (r *ReceiptLine) Equal(o *ReceiptLine) bool {
	var i2 bool

	if r.Input2 == o.Input2 {
		i2 = true
	} else if r.Input2 == nil || o.Input2 == nil {
		return false
	} else {
		i2 = *r.Input2 == *o.Input2
	}

	return r.Actual == o.Actual && r.Expected == o.Expected &&
		r.Input == o.Input && i2
}

// Returns a string representation of the line
func (r ReceiptLine) String() string {
	var builder strings.Builder
	r.ContinueBuild(&builder)
	return builder.String()
}

// Slice of Receipts
type ReceiptSlice struct {
	Receipts []Receipt
}

// Writes itself to builder
func (s *ReceiptSlice) ContinueBuild(builder *strings.Builder) {
	last := len(s.Receipts) - 1

	for i := 0; i < last; i++ {
		s.Receipts[i].ContinueBuild(builder)
		builder.WriteString("\n\n")
	}

	s.Receipts[last].ContinueBuild(builder)
}

func (r *ReceiptSlice) Len() int {
	return len(r.Receipts)
}

func (r *ReceiptSlice) Less(i, j int) bool {
	return r.Receipts[i].Name <= r.Receipts[j].Name
}

func (r *ReceiptSlice) Swap(i, j int) {
	r.Receipts[i], r.Receipts[j] = r.Receipts[j], r.Receipts[i]
}
