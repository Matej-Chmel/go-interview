package gointerview_test

import (
	"strings"
	"testing"

	goi "github.com/Matej-Chmel/go-interview"
	ite "github.com/Matej-Chmel/go-interview/internal"
)

func checkLines2(actual, expected []ite.ReceiptLine2, t *testing.T) {
	if len(actual) != len(expected) {
		ite.Throw(t, 2, "Expected %d lines, found %d", len(expected), len(actual))
	}

	for i := 0; i < len(expected); i++ {
		if !actual[i].Equal(&expected[i]) {
			a, e := actual[i].String(), expected[i].String()
			ite.Throw(t, 2, "Lines at index %d do not equal\n%s\n%s", i, a, e)
		}
	}
}

type result struct {
	A, B string
}

func badSwap(a, b string) result {
	return result{A: a, B: b}
}

func goodSwap(a, b string) result {
	return result{A: b, B: a}
}

func TestSwap(t *testing.T) {
	i := goi.NewInterview2[string, string, result]()

	i.AddCase("hello", "world", result{
		A: "world",
		B: "hello",
	})
	i.AddCase("123", ".", result{
		A: ".",
		B: "123",
	})

	i.AddSolutions(badSwap, goodSwap)

	bad, err := i.RunSolution("badSwap")
	ite.CheckName(err, bad.Name, "badSwap", t)
	checkLines2(bad.Lines, []ite.ReceiptLine2{
		ite.NewReceiptLine2("{hello world}", "{world hello}", "hello", "world"),
		ite.NewReceiptLine2("{123 .}", "{. 123}", "123", "."),
	}, t)

	good, err := i.RunSolution("goodSwap")
	ite.CheckName(err, good.Name, "goodSwap", t)
	checkLines2(good.Lines, []ite.ReceiptLine2{
		ite.NewReceiptLine2("{world hello}", "{world hello}", "hello", "world"),
		ite.NewReceiptLine2("{. 123}", "{. 123}", "123", "."),
	}, t)
}

func addOne(i, j float64) float64 {
	return i + j
}

func addTwo(i, j float64) float64 {
	return i * j
}

func TestWrite2(t *testing.T) {
	i := goi.NewInterview2[float64, float64, float64]()

	i.AddCase(1.1, 2.2, 1.1+2.2)
	i.AddCase(5.24, 0.0, 5.24)

	i.AddSolutions(addOne, addTwo)

	var builder strings.Builder
	i.WriteAllSolutions(&builder)

	actual := builder.String()
	expected := ite.ReadFile("test_data/test_write2.txt", t)

	if actual != expected {
		ite.Throw(t, 1, "\nActual:\n%s\n\nExpected:\n%s", actual, expected)
		ite.Throw(t, 1, "Actual: %d, Expected: %d", len(actual), len(expected))
	}
}

func TestExported2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			ite.Throw(t, 1, "TestExported2 did panic but should have NOT")
		}
	}()

	i := goi.NewInterview2[ite.Exported, ite.Exported, ite.Exported]()
	data := ite.Exported{A: 1, B: 2}
	i.AddCase(data, data, data)
}

func TestExportedNested2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			ite.Throw(t, 1, "TestExportedNested2 did panic but should have NOT")
		}
	}()

	i := goi.NewInterview2[ite.ExportedNested, ite.ExportedNested, ite.ExportedNested]()
	data := ite.ExportedNested{Exported: ite.Exported{A: 1, B: 2}, C: 3}
	i.AddCase(data, data, data)
}

type Unexported2 struct {
	a int
	B int
}

func TestUnexported2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			ite.Throw(t, 1, "TestUnexported2 did NOT panic")
		}
	}()

	i := goi.NewInterview2[Unexported2, Unexported2, Unexported2]()
	data := Unexported2{a: 1, B: 2}
	i.AddCase(data, data, data)
}

type UnexportedNested2 struct {
	U Unexported2
	C int
}

func TestUnexportedNested2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			ite.Throw(t, 1, "TestUnexportedNested2 did NOT panic")
		}
	}()

	i := goi.NewInterview2[UnexportedNested2, UnexportedNested2, UnexportedNested2]()
	data := UnexportedNested2{U: Unexported2{a: 1, B: 2}, C: 3}
	i.AddCase(data, data, data)
}
