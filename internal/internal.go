package internal

import (
	"fmt"
	"io"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type ReceiptLine struct {
	Actual   string
	Expected string
	Input    string
}

func NewReceiptLine(a, e, i string) ReceiptLine {
	return ReceiptLine{Actual: a, Expected: e, Input: i}
}

func (r *ReceiptLine) Equal(o *ReceiptLine) bool {
	return r.Actual == o.Actual && r.Expected == o.Expected && r.Input == o.Input
}

func (r ReceiptLine) String() string {
	var end, status string

	if r.Actual == r.Expected {
		status = "(OK)"
	} else {
		end = fmt.Sprintf(" != %s", r.Expected)
		status = "(  )"
	}

	return fmt.Sprintf("%s %s -> %s%s", status, r.Input, r.Actual, end)
}

type Receipt struct {
	Lines []ReceiptLine
	Name  string
}

func (s Receipt) String() string {
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

func (s *ReceiptLine) Write(w io.Writer) error {
	_, err := w.Write([]byte(s.String()))
	return err
}

type ReceiptSlice struct {
	Receipts []Receipt
}

func (s ReceiptSlice) String() string {
	var builder strings.Builder
	last := len(s.Receipts) - 1

	for i := 0; i < last; i++ {
		builder.WriteString(s.Receipts[i].String())
		builder.WriteRune('\n')
	}

	builder.WriteString(s.Receipts[last].String())
	return builder.String()
}

func (s *ReceiptSlice) Write(w io.Writer) error {
	_, err := w.Write([]byte(s.String()))
	return err
}

type TestCase[I any, O any] struct {
	Expected *O
	Input    *I
}

func CheckName(err error, actual, expected string, t *testing.T) {
	if err != nil {
		Throw(t, 2, "%s not found", expected)
	}

	if actual != expected {
		Throw(t, 2, "Expected: %s, found: %s", expected, actual)
	}
}

func GetFunctionName(f interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	tokens := strings.Split(name, ".")
	return tokens[len(tokens)-1]
}

func Throw(t *testing.T, skip int, format string, data ...any) {
	_, _, line, _ := runtime.Caller(skip)
	format = fmt.Sprintf("Line %d: %s", line, format)
	t.Errorf(format, data...)
}
