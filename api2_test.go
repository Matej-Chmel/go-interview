package gointerview_test

import (
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
