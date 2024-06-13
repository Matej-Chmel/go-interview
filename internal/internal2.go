package internal

import (
	"fmt"
	"io"
	"strings"
)

type ReceiptLine2 struct {
	Actual   string
	Expected string
	Input    string
	Input2   string
}

func NewReceiptLine2(a, e, i, i2 string) ReceiptLine2 {
	return ReceiptLine2{Actual: a, Expected: e, Input: i, Input2: i2}
}

func (r *ReceiptLine2) Equal(o *ReceiptLine2) bool {
	return r.Actual == o.Actual && r.Expected == o.Expected &&
		r.Input == o.Input && r.Input2 == o.Input2
}

func (r ReceiptLine2) String() string {
	var end, status string

	if r.Actual == r.Expected {
		status = "(OK)"
	} else {
		end = fmt.Sprintf(" != %s", r.Expected)
		status = "(  )"
	}

	return fmt.Sprintf("%s %s, %s -> %s%s",
		status, r.Input, r.Input2, r.Actual, end)
}

type Receipt2 struct {
	Lines []ReceiptLine2
	Name  string
}

func (s Receipt2) String() string {
	var builder strings.Builder

	builder.WriteString(s.Name)
	builder.WriteRune('\n')
	builder.WriteString(strings.Repeat("=", len(s.Name)))

	for _, l := range s.Lines {
		builder.WriteRune('\n')
		builder.WriteString(l.String())
	}

	return builder.String()
}

func (s *ReceiptLine2) Write(w io.Writer) error {
	_, err := w.Write([]byte(s.String()))
	return err
}

type ReceiptSlice2 struct {
	Receipts []Receipt2
}

func (s ReceiptSlice2) String() string {
	var builder strings.Builder
	last := len(s.Receipts) - 1

	for i := 0; i < last; i++ {
		builder.WriteString(s.Receipts[i].String())
		builder.WriteString("\n\n")
	}

	builder.WriteString(s.Receipts[last].String())
	return builder.String()
}

func (s *ReceiptSlice2) Write(w io.Writer) error {
	_, err := w.Write([]byte(s.String()))
	return err
}

type TestCase2[I any, I2 any, O any] struct {
	Expected *O
	Input    *I
	Input2   *I2
}
