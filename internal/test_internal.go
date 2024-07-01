package internal

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

// Used for testing.
// Compares all lines from actual against expected.
// Returns flag indicating whether all lines match.
func CheckLines(actual, expected []*ReceiptLine, t *testing.T) bool {
	if len(actual) != len(expected) {
		Throw(t, 2, "Expected %d lines, found %d", len(expected), len(actual))
		return false
	}

	for i := 0; i < len(expected); i++ {
		if !actual[i].Equal(expected[i]) {
			a, e := actual[i].String(), expected[i].String()
			Throw(t, 2, "Lines at index %d do not equal\n%s\n%s", i, a, e)
			return false
		}
	}

	return true
}

// Used for testing.
// Compares actual and expected names.
// Returns flag indicating whether the names match.
func CheckName(err error, actual, expected string, t *testing.T) bool {
	if err != nil {
		Throw(t, 2, "%s not found", expected)
		return false
	}

	if actual != expected {
		Throw(t, 2, "Expected: %s, found: %s", expected, actual)
		return false
	}

	return true
}

// Used for testing.
// Compares names from s against solNames.
// Returns flag indicating whether all names match.
func CheckSlice(s *ReceiptSlice, t *testing.T, solNames ...string) bool {
	if len(s.Receipts) != len(solNames) {
		Throw(t, 2, "Expected %d, found %d solutions", len(s.Receipts), len(solNames))
		return false
	}

	for i := 0; i < len(solNames); i++ {
		if a, e := s.Receipts[i].Name, solNames[i]; a != e {
			Throw(t, 2, "Name at index %d does not equal\n%s != %s", i, a, e)
			return false
		}
	}

	return true
}

// Used for testing.
// Compares actual and expected strings character by character.
// Returns flag indicating whether the strings match.
func CheckStrings(t *testing.T, skip int, actual, expected string) bool {
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
		return true
	}

	var builder strings.Builder

	for _, m := range mismatches[:5] {
		builder.WriteString(strconv.FormatInt(int64(m.index), 10))
		builder.WriteString(fmt.Sprintf(": %q != %q\n", m.actual, m.expected))
	}

	Throw(t, skip+1, "%s\n\n!=\n\n%s\n\nMismatches:\n%s",
		actual, expected, builder.String())
	return false
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

// Used for testing.
// A character mismatch at an index.
type mismatch struct {
	actual   byte
	expected byte
	index    int
}

// Used for testing.
// Formats and logs an error with a line number of the caller and fails the test.
// Skip indicates the number of stack frames to skip.
func Throw(t *testing.T, skip int, format string, data ...any) {
	out := fmt.Sprintf(format, data...)
	_, _, line, ok := runtime.Caller(skip)

	if ok {
		out = fmt.Sprintf("\n\nLine %d\n\n%s", line, out)
	} else {
		out = "\n\n" + out
	}

	t.Error(out)
}
