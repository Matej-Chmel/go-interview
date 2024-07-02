package gointerview

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	at "github.com/Matej-Chmel/go-any-to-string"
	dc "github.com/Matej-Chmel/go-deep-copy"
	ite "github.com/Matej-Chmel/go-interview/internal"
)

// Class for two input problems.
// It can act as an implementation for Interview class.
type Interview2[I any, I2 any, O any] struct {
	cases         []*ite.TestCase[I, I2, O]
	isSingleInput bool
	options       *at.Options
	solutions1    map[string]func(I) O
	solutions2    map[string]func(I, I2) O
}

// Constructs an Interview2 object
func NewInterview2[I any, I2 any, O any]() Interview2[I, I2, O] {
	return newInterview2Impl[I, I2, O](false)
}

// Internal constructor for Interview2 object.
// Decides which solution map to initialize.
func newInterview2Impl[I any, I2 any, O any](isSingleInput bool) Interview2[I, I2, O] {
	res := Interview2[I, I2, O]{
		cases:         make([]*ite.TestCase[I, I2, O], 0),
		isSingleInput: isSingleInput,
		options:       at.NewOptions(),
		solutions1:    nil,
		solutions2:    nil,
	}

	if isSingleInput {
		res.solutions1 = make(map[string]func(I) O)
	} else {
		res.solutions2 = make(map[string]func(I, I2) O)
	}

	return res
}

// Adds one test case
func (i *Interview2[I, I2, O]) AddCase(input I, input2 I2, expected O) {
	if len(i.cases) == 0 {
		if !dc.IsFullyExported(input) {
			panic("First input struct CANNOT have unexported fields")
		}

		if !dc.IsFullyExported(input2) {
			panic("Second input struct CANNOT have unexported fields")
		}

		if !dc.IsFullyExported(expected) {
			panic("Output struct CANNOT have unexported fields")
		}
	}

	i.cases = append(
		i.cases, ite.NewTestCase(&input, &input2, &expected, i.isSingleInput))
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
func (i *Interview2[I, I2, O]) AddSolution(s func(I, I2) O) {
	i.solutions2[ite.GetFunctionName(s)] = s
}

// Adds multiple solution functions
func (i *Interview2[I, I2, O]) AddSolutions(s ...func(I, I2) O) {
	for _, f := range s {
		i.AddSolution(f)
	}
}

// Runs all solutions against all test cases
// and compiles the output into a single string
func (i *Interview2[I, I2, O]) AllSolutionsToString() string {
	var builder strings.Builder
	i.WriteAllSolutions(&builder)
	return builder.String()
}

// Changes options so that byte, uint8, rune and int32 are all
// printed as characters
func (i *Interview2[I, I2, O]) BytesAsString() {
	i.options.ByteAsString = true
	i.options.RuneAsString = true
}

// Internal constructor for a new ReceiptLine
func (i *Interview2[I, I2, O]) newReceiptLine(
	actual O, c *ite.TestCase[I, I2, O],
) *ite.ReceiptLine {
	var input2 *string = nil

	if !i.isSingleInput {
		val := c.GetInput2String(i.options)
		input2 = &val
	}

	return ite.NewReceiptLineImpl(
		at.AnyToStringCustom(actual, i.options),
		c.GetExpectedString(i.options),
		c.GetInputString(i.options),
		input2,
	)
}

// Runs all solutions against all test cases
// and prints the output to the standard output
func (i *Interview2[I, I2, O]) Print() error {
	return i.WriteAllSolutions(os.Stdout)
}

// Reads one case from relative paths for inputs and output
func (i *Interview2[I, I2, O]) ReadCase(
	input1RelPath, input2RelPath, outRelPath string,
) {
	input1, err := ite.ReadData[I](input1RelPath)

	if err != nil {
		panic(err)
	}

	var input2 I2

	if !i.isSingleInput {
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

	i.AddCase(input1, input2, out)
}

// Reads multiple cases from relative paths for inputs and output
func (i *Interview2[I, I2, O]) ReadCases(
	input1RelPath, input2RelPath, outRelPath string,
) {
	i.ReadCasesSlice(
		input1RelPath, input2RelPath, outRelPath, 0, -1)
}

// Reads multiple cases from relative paths for inputs and output.
// Only the cases in range [begin, end) are added.
func (i *Interview2[I, I2, O]) ReadCasesSlice(
	input1RelPath, input2RelPath, outRelPath string,
	begin, end int,
) {
	input1, err := ite.ReadData[[]I](input1RelPath)

	if err != nil {
		panic(err)
	}

	var input2 []I2

	if !i.isSingleInput {
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

	i.AddCasesSlice(input1, input2, out, begin, end)
}

// Runs a solution for a single input problem
// against all test cases
func (i *Interview2[I, I2, O]) runFunction1(f func(I) O) (r ite.Receipt) {
	for _, c := range i.cases {
		input := dc.DeepCopy(c.Input)
		actual := f(*input)
		r.Lines = append(r.Lines, i.newReceiptLine(actual, c))
	}

	r.Name = ite.GetFunctionName(f)
	return
}

// Runs a solution for a two input problem
// against all test cases
func (i *Interview2[I, I2, O]) runFunction2(f func(I, I2) O) (r ite.Receipt) {
	for _, c := range i.cases {
		input, input2 := dc.DeepCopy(c.Input), dc.DeepCopy(c.Input2)
		actual := f(*input, *input2)
		r.Lines = append(r.Lines, i.newReceiptLine(actual, c))
	}

	r.Name = ite.GetFunctionName(f)
	return
}

// Runs one solution function against all test cases.
// If function cannot be found, an error is returned.
func (i *Interview2[I, I2, O]) RunSolution(name string) (ite.Receipt, error) {
	var exists bool
	var fn1 func(I) O
	var fn2 func(I, I2) O

	if i.isSingleInput {
		fn1, exists = i.solutions1[name]
	} else {
		fn2, exists = i.solutions2[name]
	}

	if !exists {
		res := ite.Receipt{Lines: nil, Name: ""}
		return res, fmt.Errorf("solution %s not found", name)
	}

	if i.isSingleInput {
		return i.runFunction1(fn1), nil
	}

	return i.runFunction2(fn2), nil
}

// Runs all solutions against all test cases
func (i *Interview2[I, I2, O]) RunAllSolutions() (s ite.ReceiptSlice) {
	if i.isSingleInput {
		for _, f := range i.solutions1 {
			receipt := i.runFunction1(f)
			s.Receipts = append(s.Receipts, receipt)
		}
	} else {
		for _, f := range i.solutions2 {
			receipt := i.runFunction2(f)
			s.Receipts = append(s.Receipts, receipt)
		}
	}

	sort.Slice(s.Receipts, func(i, j int) bool {
		return s.Receipts[i].Name < s.Receipts[j].Name
	})
	return
}

// Runs all solutions against all test cases
// and writes the results into a writer w
func (i *Interview2[I, I2, O]) WriteAllSolutions(w io.Writer) error {
	var builder strings.Builder

	s := i.RunAllSolutions()
	s.ContinueBuild(&builder)
	_, err := w.Write([]byte(builder.String()))

	return err
}
