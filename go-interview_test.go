package gointerview_test

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	goi "github.com/Matej-Chmel/go-interview"
)

func badFactorial(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 6
	case 4:
		return 24
	case 5:
		return 120
	default:
		return n * n
	}
}

func iterativeFactorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func recursiveFactorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * recursiveFactorial(n-1)
}

type checker[I any, O any] struct {
	functionName string
	t            *testing.T
	mismatches   []goi.Mismatch
}

func newChecker[I any, O any](fname string, t *testing.T) checker[I, O] {
	return checker[I, O]{fname, t, make([]goi.Mismatch, 0)}
}

func (c *checker[I, O]) addMismatch(field string, left string, right string) {
	c.mismatches = append(c.mismatches, goi.Mismatch{field, left, right})
}

func (c *checker[I, O]) clear() {
	c.mismatches = make([]goi.Mismatch, 0)
}

func (c *checker[I, O]) check(report goi.Report[I, O]) {
	if c.functionName != report.FunctionName {
		c.throw(report)
		return
	}

	if len(c.mismatches) != report.Len() {
		c.throw(report)
		return
	}

	limit := report.Len()

	for i := 0; i < limit; i++ {
		if !c.mismatches[i].Equals(report.Get(i)) {
			c.throw(report)
			return
		}
	}
}

func (c *checker[I, O]) checkSlice(res []goi.Report[I, O]) {
	for _, report := range res {
		c.check(report)
	}
}

func (c *checker[I, O]) testError(reason string, skip int) {
	_, _, line, _ := runtime.Caller(skip)
	c.t.Errorf("Error at line %d\n\n%s\n", line, reason)
}

func (c *checker[I, O]) throw(report goi.Report[I, O]) {
	c.testError(fmt.Sprintf("%s\n\n%s", report.String(), c.string()), 1)
}

func (c *checker[I, O]) string() string {
	return sliceToStringCustom(c.mismatches, "\n\n", func(m goi.Mismatch) string {
		return fmt.Sprintf("%s: \"%s\" <--> \"%s\"", m.Field, m.Left, m.Right)
	})
}

func sliceToBuilderCustom[T any](
	builder *strings.Builder, slice []T,
	sep string, conv func(T) string,
) {
	lastIndex := len(slice) - 1

	for i := 0; i < lastIndex; i++ {
		builder.WriteString(conv(slice[i]))
		builder.WriteString(sep)
	}

	builder.WriteString(conv(slice[lastIndex]))
}

func sliceToStringCustom[T any](
	slice []T, sep string, conv func(T) string,
) string {
	var builder strings.Builder
	sliceToBuilderCustom(&builder, slice, sep, conv)
	return builder.String()
}

func TestBad(t *testing.T) {
	it := goi.NewInterview[int, int]()
	it.AddCase(1, 1).AddCase(2, 2).AddCase(3, 6)
	it.AddCase(4, 24).AddCase(5, 120).AddCase(6, 720)

	it.AddSolution(badFactorial)

	c := newChecker[int, int]("badFactorial", t)
	res, err := it.RunSolution("badFactorial")

	if err != nil {
		c.testError("Function not found", 1)
	}

	c.addMismatch("Value", "36", "720")
	c.check(res[5])
	c.clear()

	c.checkSlice(res[0:5])
}

func TestMultiple(t *testing.T) {
	it := goi.NewInterview[int, int]()
	it.AddCase(1, 1).AddCase(2, 2).AddCase(3, 6)
	it.AddCase(4, 24).AddCase(5, 120).AddCase(6, 720)

	it.AddSolutions(iterativeFactorial, recursiveFactorial)

	res := it.RunAllSolutions()

	c := newChecker[int, int]("iterativeFactorial", t)
	c.checkSlice(res[0:6])

	c = newChecker[int, int]("recursiveFactorial", t)
	c.checkSlice(res[7:])
}
