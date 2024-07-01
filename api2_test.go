package gointerview_test

import (
	"strings"
	"testing"

	goi "github.com/Matej-Chmel/go-interview"
	ite "github.com/Matej-Chmel/go-interview/internal"
)

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
	success := ite.CheckName(err, bad.Name, "badSwap", t)

	if !success {
		return
	}

	success = ite.CheckLines(bad.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine2("hello", "world", "{hello world}", "{world hello}"),
		ite.NewReceiptLine2("123", ".", "{123 .}", "{. 123}"),
	}, t)

	if !success {
		return
	}

	good, err := i.RunSolution("goodSwap")
	success = ite.CheckName(err, good.Name, "goodSwap", t)

	if !success {
		return
	}

	ite.CheckLines(good.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine2("hello", "world", "{world hello}", "{world hello}"),
		ite.NewReceiptLine2("123", ".", "{. 123}", "{. 123}"),
	}, t)
}

func addOne(i, j float64) float64 {
	return i + j
}

func addTwo(i, j float64) float64 {
	return i * j
}

func TestStdout2(t *testing.T) {
	i := goi.NewInterview2[float64, float64, float64]()
	i.AddCase(1.1, 2.2, 1.1+2.2)
	i.AddCase(5.24, 0.0, 5.24)
	i.AddSolutions(addOne, addTwo)

	var builder strings.Builder
	i.WriteAllSolutions(&builder)

	if expected, err := ite.ReadAllText("test_data/add_stdout.txt"); err != nil {
		ite.Throw(t, 1, err.Error())
	} else {
		ite.CheckStrings(t, 1, i.AllSolutionsToString(), expected)
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

func floatFilter(data []float32, filter []bool) (res []float32) {
	for i, v := range filter {
		if v && i < len(data) {
			res = append(res, data[i])
		}
	}

	if len(data) > 0 && data[0] > 8. {
		res = append(res, data[0])
	}

	return
}

func TestFilter(t *testing.T) {
	i := goi.NewInterview2[[]float32, []bool, []float32]()
	i.AddSolution(floatFilter)
	i.ReadCases(
		"test_data/filter_in.txt",
		"test_data/filter_in2.txt",
		"test_data/filter_out.txt")

	rec, err := i.RunSolution("floatFilter")
	success := ite.CheckName(err, rec.Name, "floatFilter", t)

	if !success {
		return
	}

	success = ite.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine2("[1.1 2.22 3.333]", "[true false true]", "[1.1 3.333]", "[1.1 3.333]"),
		ite.NewReceiptLine2("[4.44 5.555 6.667 7.778]", "[false true]", "[5.555]", "[5.555]"),
		ite.NewReceiptLine2("[8.888 9.999 10.101 11.111]", "[true false false true]", "[8.888 11.111 8.888]", "[8.888 11.111]"),
		ite.NewReceiptLine2("[12.121 13.131 14.141 15.152 16.162]", "[false true true false]", "[13.131 14.141 12.121]", "[13.131 14.141]"),
		ite.NewReceiptLine2("[17.172 18.182 19.192 20.202]", "[true false]", "[17.172 17.172]", "[17.172]"),
	}, t)

	if !success {
		return
	}

	if expected, err := ite.ReadAllText("test_data/filter_stdout.txt"); err != nil {
		ite.Throw(t, 1, err.Error())
	} else {
		ite.CheckStrings(t, 1, i.AllSolutionsToString(), expected)
	}
}

func matrixMult(a, b [][]int8) [][]int8 {
	if len(a) != len(b) {
		return nil
	}

	res := make([][]int8, len(a))

	for i := 0; i < len(a); i++ {
		aRow, bRow := a[i], b[i]

		if len(aRow) != len(bRow) {
			return nil
		}

		res[i] = make([]int8, len(aRow))

		for j := 0; j < len(aRow); j++ {
			res[i][j] = aRow[j] * bRow[j]
		}
	}

	return res
}

func TestMatrixMult(t *testing.T) {
	i := goi.NewInterview2[[][]int8, [][]int8, [][]int8]()
	i.AddSolution(matrixMult)
	i.ReadCases(
		"test_data/matrixMult_in.txt",
		"test_data/matrixMult_in2.txt",
		"test_data/matrixMult_out.txt")

	if expected, err := ite.ReadAllText("test_data/matrixMult_stdout.txt"); err != nil {
		ite.Throw(t, 1, err.Error())
	} else {
		ite.CheckStrings(t, 1, i.AllSolutionsToString(), expected)
	}
}
