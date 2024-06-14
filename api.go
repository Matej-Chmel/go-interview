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

type interview[I any, O any] struct {
	cases     []*ite.TestCase[I, O]
	solutions map[string]func(I) O
}

func NewInterview[I any, O any]() interview[I, O] {
	return interview[I, O]{
		cases:     make([]*ite.TestCase[I, O], 0),
		solutions: make(map[string]func(I) O),
	}
}

func (i *interview[I, O]) AddCase(input I, expected O) {
	if len(i.cases) == 0 {
		if !dc.IsFullyExported(input) {
			panic("Input struct CANNOT have unexported fields")
		}

		if !dc.IsFullyExported(expected) {
			panic("Output struct CANNOT have unexported fields")
		}
	}

	i.cases = append(i.cases, &ite.TestCase[I, O]{
		Expected: dc.DeepCopy(&expected),
		Input:    dc.DeepCopy(&input),
	})
}

func (i *interview[I, O]) AddSolution(s func(I) O) {
	i.solutions[ite.GetFunctionName(s)] = s
}

func (i *interview[I, O]) AddSolutions(s ...func(I) O) {
	for _, f := range s {
		i.AddSolution(f)
	}
}

func (i *interview[I, O]) getReceipt(f func(I) O, name string) (r ite.Receipt, e error) {
	r.Lines, e = i.runFunction(f)
	r.Name = name
	return
}

func (i *interview[I, O]) runFunction(f func(I) O) ([]ite.ReceiptLine, error) {
	results := make([]ite.ReceiptLine, 0)

	o := at.NewOptions()
	o.ByteAsString = true
	o.RuneAsString = true

	for _, c := range i.cases {
		input := dc.DeepCopy(c.Input)
		actual := f(*input)
		res := ite.ReceiptLine{
			Actual:   at.AnyToStringCustom(actual, o),
			Expected: at.AnyToStringCustom(*c.Expected, o),
			Input:    at.AnyToStringCustom(*input, o),
		}
		results = append(results, res)
	}

	return results, nil
}

func (i *interview[I, O]) RunSolution(name string) (r ite.Receipt, e error) {
	s, exists := i.solutions[name]

	if !exists {
		e = fmt.Errorf("solution %s not found", name)
		return
	}

	return i.getReceipt(s, name)
}

func (i *interview[I, O]) RunAllSolutions() (s ite.ReceiptSlice) {
	for name, f := range i.solutions {
		receipt, _ := i.getReceipt(f, name)
		s.Receipts = append(s.Receipts, receipt)
	}

	sort.Slice(s.Receipts, func(i, j int) bool {
		return s.Receipts[i].Name < s.Receipts[j].Name
	})
	return
}

func (i *interview[I, O]) Print() error {
	return i.WriteAllSolutions(os.Stdout)
}

func (i *interview[I, O]) WriteAllSolutions(w io.Writer) error {
	s := i.RunAllSolutions()
	return s.Write(w)
}
