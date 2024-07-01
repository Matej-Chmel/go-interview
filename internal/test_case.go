package internal

import (
	at "github.com/Matej-Chmel/go-any-to-string"
	dc "github.com/Matej-Chmel/go-deep-copy"
)

// Test case with one or two inputs and an output
type TestCase[I any, I2 any, O any] struct {
	expectedString string
	inputString    string
	input2String   string
	Expected       *O
	Input          *I
	Input2         *I2
}

// Constructs a test case
func NewTestCase[I any, I2 any, O any](
	i1 *I, i2 *I2, o *O, isSingleInput bool) *TestCase[I, I2, O] {

	res := &TestCase[I, I2, O]{
		Expected: dc.DeepCopy(o),
		Input:    dc.DeepCopy(i1),
		Input2:   nil,
	}

	if !isSingleInput {
		res.Input2 = dc.DeepCopy(i2)
	}

	return res
}

// Lazy loads and returns string representing expected result
func (c *TestCase[I, I2, O]) GetExpectedString(o *at.Options) string {
	if c.expectedString == "" {
		c.expectedString = at.AnyToStringCustom(*c.Expected, o)
	}

	return c.expectedString
}

// Lazy loads and returns string representing first input
func (c *TestCase[I, I2, O]) GetInputString(o *at.Options) string {
	if c.inputString == "" {
		c.inputString = at.AnyToStringCustom(*c.Input, o)
	}

	return c.inputString
}

// Lazy loads and returns string representing second input
func (c *TestCase[I, I2, O]) GetInput2String(o *at.Options) string {
	if c.input2String == "" && c.Input2 != nil {
		c.input2String = at.AnyToStringCustom(*c.Input2, o)
	}

	return c.input2String
}
