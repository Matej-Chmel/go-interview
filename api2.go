package gointerview

import (
	"fmt"
	"io"
	"os"
	"strings"

	at "github.com/Matej-Chmel/go-any-to-string"
	dc "github.com/Matej-Chmel/go-deep-copy"
	ite "github.com/Matej-Chmel/go-interview/internal"
)

// Class for two input problems.
// It can act as an implementation for Interview class.
type Interview2[I any, I2 any, O any] struct {
	*ite.EmbeddedOptions
	byteFlags     uint
	cases         []*ite.TestCase[I, I2, O]
	isSingleInput bool
	solutions1    map[string]func(I) O
	solutions2    map[string]func(I, I2) O
}

// Constructs an Interview2 object
func NewInterview2[I any, I2 any, O any]() Interview2[I, I2, O] {
	options := ite.NewEmbeddedOptions()
	return newInterview2Impl[I, I2, O](false, &options)
}

// Internal constructor for Interview2 object.
// Decides which solution map to initialize.
func newInterview2Impl[I any, I2 any, O any](
	isSingleInput bool,
	options *ite.EmbeddedOptions,
) Interview2[I, I2, O] {
	res := Interview2[I, I2, O]{
		byteFlags:       0,
		cases:           make([]*ite.TestCase[I, I2, O], 0),
		EmbeddedOptions: options,
		isSingleInput:   isSingleInput,
		solutions1:      nil,
		solutions2:      nil,
	}

	if ite.IsByte[I]() {
		res.byteFlags |= ite.Input1Byte
	}

	if ite.IsByte[I2]() {
		res.byteFlags |= ite.Input2Byte
	}

	if ite.IsByte[O]() {
		res.byteFlags |= ite.OutputByte
	}

	if isSingleInput {
		res.solutions1 = make(map[string]func(I) O)
	} else {
		res.solutions2 = make(map[string]func(I, I2) O)
	}

	return res
}

// Adds one test case
func (iv *Interview2[I, I2, O]) AddCase(input I, input2 I2, expected O) {
	testCase := ite.NewTestCase(&input, &input2, &expected, iv.isSingleInput)
	iv.cases = append(iv.cases, testCase)
}

// Converts strings to a byte or rune slices and attempts
// to add those slices as a new test case.
// Panics if any string cannot be converted to its target type.
func (iv *Interview2[I, I2, O]) AddCaseString(
	s1 string, s2 string, exp string,
) {
	input := ite.ConvertToSlice[I](s1, iv.byteFlags, ite.Input1Byte)
	expected := ite.ConvertToSlice[O](exp, iv.byteFlags, ite.OutputByte)

	var input2 I2

	if !iv.isSingleInput {
		input2 = ite.ConvertToSlice[I2](s2, iv.byteFlags, ite.Input2Byte)
	}

	iv.AddCase(input, input2, expected)
}

// Adds multiple test cases
func (iv *Interview2[I, I2, O]) AddCases(input1 []I, input2 []I2, expected []O) {
	iv.AddCasesSlice(input1, input2, expected, 0, -1)
}

// Adds multiple test cases in range [begin, end)
func (iv *Interview2[I, I2, O]) AddCasesSlice(
	input1 []I, input2 []I2, expected []O,
	begin, end int,
) {
	len1, lenO := len(input1), len(expected)
	var len2 int

	if iv.isSingleInput {
		len2 = lenO
	} else {
		len2 = len(input2)
	}

	if len1 != len2 || len1 != lenO || len2 != lenO {
		panic(fmt.Errorf(
			"Length of inputs and outputs don't match %d:%d:%d",
			len1, len2, lenO))
	}

	if end < 0 {
		end = lenO
	} else {
		end = min(lenO, end)
	}

	if iv.isSingleInput {
		var i2 I2

		for i := begin; i < end; i++ {
			iv.AddCase(input1[i], i2, expected[i])
		}
	} else {
		for i := begin; i < end; i++ {
			iv.AddCase(input1[i], input2[i], expected[i])
		}
	}
}

// Adds one solution function
func (iv *Interview2[I, I2, O]) AddSolution(s func(I, I2) O) {
	iv.solutions2[ite.GetFunctionName(s)] = s
}

// Adds multiple solution functions
func (iv *Interview2[I, I2, O]) AddSolutions(s ...func(I, I2) O) {
	for _, f := range s {
		iv.AddSolution(f)
	}
}

// Runs all solutions against all test cases
// and compiles the output into a single string
func (iv *Interview2[I, I2, O]) AllSolutionsToString() string {
	var builder strings.Builder
	iv.WriteAllSolutions(&builder)
	return builder.String()
}

// Returns true if no test cases are available
func (iv *Interview2[I, I2, O]) noCases() bool {
	return len(iv.cases) == 0
}

// Returns true if no solutions are available
func (iv *Interview2[I, I2, O]) noSolutions() bool {
	return (iv.isSingleInput && len(iv.solutions1) == 0) ||
		(!iv.isSingleInput && len(iv.solutions2) == 0)
}

// Internal constructor for a new ReceiptLine
func (iv *Interview2[I, I2, O]) newReceiptLine(
	actual O, c *ite.TestCase[I, I2, O],
) *ite.ReceiptLine {
	var input2 *string = nil
	options := iv.GetOptions()

	if !iv.isSingleInput {
		val := c.GetInput2String(options)
		input2 = &val
	}

	return ite.NewReceiptLineImpl(
		at.AnyToStringCustom(actual, options),
		c.GetExpectedString(options),
		c.GetInputString(options),
		input2,
	)
}

// Runs all solutions against all test cases
// and prints the output to the standard output
func (iv *Interview2[I, I2, O]) Print() error {
	return iv.WriteAllSolutions(os.Stdout)
}

// Reads one case from relative paths for inputs and output
func (iv *Interview2[I, I2, O]) ReadCase(
	input1RelPath, input2RelPath, outRelPath string,
) {
	input1, err := ite.ReadData[I](input1RelPath)

	if err != nil {
		panic(err)
	}

	var input2 I2

	if !iv.isSingleInput {
		val, err := ite.ReadData[I2](input2RelPath)

		if err != nil {
			panic(err)
		}

		input2 = val
	}

	out, err := ite.ReadData[O](outRelPath)

	if err != nil {
		panic(err)
	}

	iv.AddCase(input1, input2, out)
}

// Reads multiple cases from relative paths for inputs and output
func (iv *Interview2[I, I2, O]) ReadCases(
	input1RelPath, input2RelPath, outRelPath string,
) {
	iv.ReadCasesSlice(
		input1RelPath, input2RelPath, outRelPath, 0, -1)
}

// Reads multiple cases from relative paths for inputs and output.
// Only the cases in range [begin, end) are added.
func (iv *Interview2[I, I2, O]) ReadCasesSlice(
	input1RelPath, input2RelPath, outRelPath string,
	begin, end int,
) {
	input1, err := ite.ReadData[[]I](input1RelPath)

	if err != nil {
		panic(err)
	}

	var input2 []I2

	if !iv.isSingleInput {
		val, err := ite.ReadData[[]I2](input2RelPath)

		if err != nil {
			panic(err)
		}

		input2 = val
	}

	out, err := ite.ReadData[[]O](outRelPath)

	if err != nil {
		panic(err)
	}

	iv.AddCasesSlice(input1, input2, out, begin, end)
}

// Runs a solution for a single input problem
// against all test cases
func (iv *Interview2[I, I2, O]) runFunction1(f func(I) O) (r ite.Receipt) {
	r.Lines = make([]*ite.ReceiptLine, len(iv.cases))

	for i, c := range iv.cases {
		input := dc.DeepCopy(c.Input)
		actual := f(*input)
		r.Lines[i] = iv.newReceiptLine(actual, c)
	}

	r.Name = ite.GetFunctionName(f)
	return
}

// Runs a solution for a two input problem
// against all test cases
func (iv *Interview2[I, I2, O]) runFunction2(f func(I, I2) O) (r ite.Receipt) {
	r.Lines = make([]*ite.ReceiptLine, len(iv.cases))

	for i, c := range iv.cases {
		input, input2 := dc.DeepCopy(c.Input), dc.DeepCopy(c.Input2)
		actual := f(*input, *input2)
		r.Lines[i] = iv.newReceiptLine(actual, c)
	}

	r.Name = ite.GetFunctionName(f)
	return
}

// Runs one solution function against all test cases.
// If function cannot be found, an error is returned.
func (iv *Interview2[I, I2, O]) RunSolution(name string) (ite.Receipt, error) {
	var exists bool
	var fn1 func(I) O
	var fn2 func(I, I2) O

	if iv.isSingleInput {
		fn1, exists = iv.solutions1[name]
	} else {
		fn2, exists = iv.solutions2[name]
	}

	if !exists {
		res := ite.Receipt{Lines: nil, Name: ""}
		return res, fmt.Errorf("solution %s not found", name)
	}

	if iv.isSingleInput {
		return iv.runFunction1(fn1), nil
	}

	return iv.runFunction2(fn2), nil
}

// Runs all solutions against all test cases
func (iv *Interview2[I, I2, O]) RunAllSolutions() ite.ReceiptSlice {
	if iv.isSingleInput {
		return ite.ExecuteSolutions(iv.solutions1, iv.runFunction1)
	}

	return ite.ExecuteSolutions(iv.solutions2, iv.runFunction2)
}

// Runs all solutions against all test cases
// and writes the results into a writer w
func (iv *Interview2[I, I2, O]) WriteAllSolutions(w io.Writer) error {
	var err error

	if iv.noCases() {
		_, err = w.Write([]byte("No test cases provided by the user!"))
	} else if iv.noSolutions() {
		_, err = w.Write([]byte("No solution functions provided by the user!"))
	} else {
		var builder strings.Builder
		slice := iv.RunAllSolutions()
		slice.ContinueBuild(&builder)
		_, err = w.Write([]byte(builder.String()))
	}

	return err
}
