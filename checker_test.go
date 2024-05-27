package gointerview_test

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	goi "github.com/Matej-Chmel/go-interview"
)

type Checker[I any, O any] struct {
	functionName string
	t            *testing.T
	mismatches   []goi.Mismatch
}

func NewChecker[I any, O any](fname string, t *testing.T) Checker[I, O] {
	return Checker[I, O]{fname, t, make([]goi.Mismatch, 0)}
}

func (c *Checker[I, O]) AddMismatch(field string, left string, right string) {
	c.mismatches = append(c.mismatches, goi.Mismatch{
		Field: field, Left: left, Right: right})
}

func (c *Checker[I, O]) Clear() {
	c.mismatches = make([]goi.Mismatch, 0)
}

func (c *Checker[I, O]) Check(report goi.Report[I, O]) {
	if c.functionName != report.FunctionName {
		c.Throw(report)
		return
	}

	if len(c.mismatches) != report.Len() {
		c.Throw(report)
		return
	}

	limit := report.Len()

	for i := 0; i < limit; i++ {
		if !c.mismatches[i].Equals(report.Get(i)) {
			c.Throw(report)
			return
		}
	}
}

func (c *Checker[I, O]) CheckSlice(res []goi.Report[I, O]) {
	for _, report := range res {
		c.Check(report)
	}
}

func (c *Checker[I, O]) TestError(reason string, skip int) {
	_, _, line, _ := runtime.Caller(skip)
	c.t.Errorf("Error at line %d\n\n%s\n", line, reason)
}

func (c *Checker[I, O]) TestErrorf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	c.TestError(s, 2)
}

func (c *Checker[I, O]) Throw(report goi.Report[I, O]) {
	c.TestError(fmt.Sprintf("%s\n\n%s", report.String(), c.string()), 1)
}

func (c *Checker[I, O]) string() string {
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
