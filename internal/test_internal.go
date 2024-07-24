package internal

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

// Wrapper around the test state
type Tester struct {
	failed bool
	*testing.T
}

// Constructs new Tester
func NewTester(t *testing.T) Tester {
	return Tester{failed: false, T: t}
}

// Compares all lines from actual against expected
func (t *Tester) CheckLines(actual, expected []*ReceiptLine) {
	if t.failed {
		return
	}

	if len(actual) != len(expected) {
		t.Throw(2, "Expected %d lines, found %d", len(expected), len(actual))
		return
	}

	for i := 0; i < len(expected); i++ {
		if !actual[i].Equal(expected[i]) {
			a, e := actual[i].String(), expected[i].String()
			t.Throw(2, "Lines at index %d do not equal\n%s\n%s", i, a, e)
			return
		}
	}
}

// Compares actual and expected names
func (t *Tester) CheckName(err error, actual, expected string) {
	if t.failed {
		return
	}

	if err != nil {
		t.Throw(2, "%s not found", expected)
		return
	}

	if actual != expected {
		t.Throw(2, "Expected: %s, found: %s", expected, actual)
		return
	}
}

// Compares names from s against solNames
func (t *Tester) CheckSlice(s *ReceiptSlice, solNames ...string) {
	if t.failed {
		return
	}

	if len(s.Receipts) != len(solNames) {
		t.Throw(2, "Expected %d, found %d solutions", len(s.Receipts), len(solNames))
		return
	}

	for i := 0; i < len(solNames); i++ {
		if a, e := s.Receipts[i].Name, solNames[i]; a != e {
			t.Throw(2, "Name at index %d does not equal\n%s != %s", i, a, e)
			return
		}
	}
}

// Compares actual and expected strings character by character
func (t *Tester) CheckStrings(skip int, actual, expected string) {
	if t.failed {
		return
	}

	i, j, mismatches := 0, 0, make([]mismatch, 0)

	for i < len(actual) && j < len(expected) {
		if actual[i] != expected[j] {
			mismatches = append(mismatches, mismatch{
				actual:   actual[i],
				expected: expected[i],
				index:    i,
			})
		}

		i++
		j++
	}

	if len(mismatches) == 0 {
		return
	}

	var builder strings.Builder

	for _, m := range mismatches[:5] {
		builder.WriteString(strconv.FormatInt(int64(m.index), 10))
		builder.WriteString(fmt.Sprintf(": %q != %q\n", m.actual, m.expected))
	}

	t.Throw(skip+1, "%s\n\n!=\n\n%s\n\nMismatches:\n%s",
		actual, expected, builder.String())
}

// Used for testing.
// An example struct with two exported integer fields.
type Exported struct {
	A int
	B int
}

// Used for testing.
// An example struct with another inner struct.
type ExportedNested struct {
	Exported
	C int
}

// Custom string method
func (e ExportedNested) String() string {
	return fmt.Sprintf("%d/%d/%d", e.A, e.B, e.C)
}

// Used for testing.
// A character mismatch at an index.
type mismatch struct {
	actual   byte
	expected byte
	index    int
}

// Formats and logs an error with a line number of the caller and fails the test.
// Skip indicates the number of stack frames to skip.
func (t *Tester) Throw(skip int, format string, data ...any) {
	if t.failed {
		return
	}

	out := fmt.Sprintf(format, data...)
	_, _, line, ok := runtime.Caller(skip)

	if ok {
		out = fmt.Sprintf("\n\nLine %d\n\n%s", line, out)
	} else {
		out = "\n\n" + out
	}

	t.Error(out)
	t.failed = true
}
