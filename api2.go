package gointerview

import (
	"fmt"
	"io"
	"os"
	"sort"

	at "github.com/Matej-Chmel/go-any-to-string"
	dc "github.com/Matej-Chmel/go-deep-copy"
	ite "github.com/Matej-Chmel/go-interview/internal"
)

type interview2[I any, I2 any, O any] struct {
	cases     []*ite.TestCase2[I, I2, O]
	solutions map[string]func(I, I2) O
}

func NewInterview2[I any, I2 any, O any]() interview2[I, I2, O] {
	return interview2[I, I2, O]{
		cases:     make([]*ite.TestCase2[I, I2, O], 0),
		solutions: make(map[string]func(I, I2) O),
	}
}

func (i *interview2[I, I2, O]) AddCase(input I, input2 I2, expected O) {
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

	i.cases = append(i.cases, &ite.TestCase2[I, I2, O]{
		Expected: dc.DeepCopy(&expected),
		Input:    dc.DeepCopy(&input),
		Input2:   dc.DeepCopy(&input2),
	})
}

func (i *interview2[I, I2, O]) AddSolution(s func(I, I2) O) {
	i.solutions[ite.GetFunctionName(s)] = s
}

func (i *interview2[I, I2, O]) AddSolutions(s ...func(I, I2) O) {
	for _, f := range s {
		i.AddSolution(f)
	}
}

func (i *interview2[I, I2, O]) getReceipt(f func(I, I2) O, name string) (r ite.Receipt2, e error) {
	r.Lines, e = i.runFunction(f)
	r.Name = name
	return
}

func (i *interview2[I, I2, O]) runFunction(f func(I, I2) O) ([]ite.ReceiptLine2, error) {
	results := make([]ite.ReceiptLine2, 0)

	o := at.NewOptions()
	o.ByteAsString = true
	o.RuneAsString = true

	for _, c := range i.cases {
		i, i2 := dc.DeepCopy(c.Input), dc.DeepCopy(c.Input2)
		actual := f(*i, *i2)
		res := ite.ReceiptLine2{
			Actual:   at.AnyToStringCustom(actual, o),
			Expected: at.AnyToStringCustom(*c.Expected, o),
			Input:    at.AnyToStringCustom(*i, o),
			Input2:   at.AnyToStringCustom(*i2, o),
		}
		results = append(results, res)
	}

	return results, nil
}

func (i *interview2[I, I2, O]) RunSolution(name string) (r ite.Receipt2, e error) {
	s, exists := i.solutions[name]

	if !exists {
		e = fmt.Errorf("solution %s not found", name)
		return
	}

	return i.getReceipt(s, name)
}

func (i *interview2[I, I2, O]) RunAllSolutions() (s ite.ReceiptSlice2) {
	for name, f := range i.solutions {
		receipt, _ := i.getReceipt(f, name)
		s.Receipts = append(s.Receipts, receipt)
	}

	sort.Slice(s.Receipts, func(i, j int) bool {
		return s.Receipts[i].Name < s.Receipts[j].Name
	})
	return
}

func (i *interview2[I, I2, O]) Print() error {
	return i.WriteAllSolutions(os.Stdout)
}

func (i *interview2[I, I2, O]) WriteAllSolutions(w io.Writer) error {
	s := i.RunAllSolutions()
	return s.Write(w)
}
