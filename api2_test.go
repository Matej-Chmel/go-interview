package gointerview_test

import (
	"strings"
	"testing"

	goi "github.com/Matej-Chmel/go-interview"
	ite "github.com/Matej-Chmel/go-interview/internal"
)

type swapResult struct {
	A, B string
}

type unexported2 struct {
	a float32
	B float32
}

type unexportedNested2 struct {
	unexported2
	C bool
}

func addOne(i, j float64) float64 {
	return i + j
}

func addTwo(i, j float64) float64 {
	return i * j
}

func badSwap(a, b string) swapResult {
	return swapResult{A: a, B: b}
}

func exportedSum(a, b ite.Exported) ite.Exported {
	return ite.Exported{
		A: a.A + b.A,
		B: a.B + b.B,
	}
}

func exportedNestedSum(a, b ite.ExportedNested) ite.ExportedNested {
	return ite.ExportedNested{
		Exported: exportedSum(a.Exported, b.Exported),
		C:        a.C + b.C,
	}
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

func goodSwap(a, b string) swapResult {
	return swapResult{A: b, B: a}
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

func unexportedNestedProduct(a, b unexportedNested2) unexportedNested2 {
	return unexportedNested2{
		unexported2: unexported2{a: a.a * b.a, B: a.B * b.B},
		C:           a.C && b.C,
	}
}

func unexportedProduct(a, b unexported2) unexported2 {
	return unexported2{
		a: a.a * b.a,
		B: a.B * b.B,
	}
}

func Test2Exported(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview2[ite.Exported, ite.Exported, ite.Exported]()
	iv.AddCase(
		ite.Exported{A: 1, B: 2},
		ite.Exported{A: 10, B: 20},
		ite.Exported{A: 11, B: 22})
	iv.AddSolution(exportedSum)
	iv.ShowFieldNames()

	rec, err := iv.RunSolution("exportedSum")
	t.CheckName(err, rec.Name, "exportedSum")
	t.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine2(
			"{A:1 B:2}", "{A:10 B:20}", "{A:11 B:22}", "{A:11 B:22}"),
	})
}

func Test2ExportedNested(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview2[ite.ExportedNested, ite.ExportedNested, ite.ExportedNested]()
	iv.AddCase(
		ite.ExportedNested{
			Exported: ite.Exported{A: 1, B: 2},
			C:        -100,
		},
		ite.ExportedNested{
			Exported: ite.Exported{A: -1, B: -2},
			C:        100,
		},
		ite.ExportedNested{
			Exported: ite.Exported{A: 0, B: 0},
			C:        0,
		},
	)
	iv.AddSolution(exportedNestedSum)
	iv.GetOptions().IgnoreCustomMethod = true
	iv.ShowFieldNames()

	rec, err := iv.RunSolution("exportedNestedSum")
	t.CheckName(err, rec.Name, "exportedNestedSum")
	t.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine2(
			"{Exported:{A:1 B:2} C:-100}",
			"{Exported:{A:-1 B:-2} C:100}",
			"{Exported:{A:0 B:0} C:0}",
			"{Exported:{A:0 B:0} C:0}"),
	})
}

func Test2Filter(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview2[[]float32, []bool, []float32]()
	iv.AddSolution(floatFilter)
	iv.ReadCases(
		"test_data/filter_in.txt",
		"test_data/filter_in2.txt",
		"test_data/filter_out.txt")

	rec, err := iv.RunSolution("floatFilter")
	t.CheckName(err, rec.Name, "floatFilter")
	t.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine2("[1.1 2.22 3.333]", "[true false true]", "[1.1 3.333]", "[1.1 3.333]"),
		ite.NewReceiptLine2("[4.44 5.555 6.667 7.778]", "[false true]", "[5.555]", "[5.555]"),
		ite.NewReceiptLine2("[8.888 9.999 10.101 11.111]", "[true false false true]", "[8.888 11.111 8.888]", "[8.888 11.111]"),
		ite.NewReceiptLine2("[12.121 13.131 14.141 15.152 16.162]", "[false true true false]", "[13.131 14.141 12.121]", "[13.131 14.141]"),
		ite.NewReceiptLine2("[17.172 18.182 19.192 20.202]", "[true false]", "[17.172 17.172]", "[17.172]"),
	})

	if expected, err := ite.ReadAllText("test_data/filter_stdout.txt"); err != nil {
		t.Throw(1, err.Error())
	} else {
		t.CheckStrings(1, iv.AllSolutionsToString(), expected)
	}
}

func Test2MatrixMult(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview2[[][]int8, [][]int8, [][]int8]()
	iv.AddSolution(matrixMult)
	iv.ReadCases(
		"test_data/matrixMult_in.txt",
		"test_data/matrixMult_in2.txt",
		"test_data/matrixMult_out.txt")

	if expected, err := ite.ReadAllText("test_data/matrixMult_stdout.txt"); err != nil {
		t.Throw(1, err.Error())
	} else {
		t.CheckStrings(1, iv.AllSolutionsToString(), expected)
	}
}

func Test2Stdout(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview2[float64, float64, float64]()
	iv.AddCase(1.1, 2.2, 1.1+2.2)
	iv.AddCase(5.24, 0.0, 5.24)
	iv.AddSolutions(addOne, addTwo)

	var builder strings.Builder
	iv.WriteAllSolutions(&builder)

	if expected, err := ite.ReadAllText("test_data/add_stdout.txt"); err != nil {
		t.Throw(1, err.Error())
	} else {
		t.CheckStrings(1, iv.AllSolutionsToString(), expected)
	}
}

func Test2Swap(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview2[string, string, swapResult]()

	iv.AddCase("hello", "world", swapResult{
		A: "world",
		B: "hello",
	})
	iv.AddCase("123", ".", swapResult{
		A: ".",
		B: "123",
	})

	iv.AddSolutions(badSwap, goodSwap)

	bad, err := iv.RunSolution("badSwap")
	t.CheckName(err, bad.Name, "badSwap")

	t.CheckLines(bad.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine2("hello", "world", "{hello world}", "{world hello}"),
		ite.NewReceiptLine2("123", ".", "{123 .}", "{. 123}"),
	})

	good, err := iv.RunSolution("goodSwap")
	t.CheckName(err, good.Name, "goodSwap")

	t.CheckLines(good.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine2("hello", "world", "{world hello}", "{world hello}"),
		ite.NewReceiptLine2("123", ".", "{. 123}", "{. 123}"),
	})
}

func Test2Unexported(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview2[unexported2, unexported2, unexported2]()
	iv.AddCase(
		unexported2{a: 100, B: 2},
		unexported2{a: -71, B: 24},
		unexported2{a: -7100, B: 48})
	iv.AddSolution(unexportedProduct)
	iv.ShowFieldNames()

	rec, err := iv.RunSolution("unexportedProduct")
	t.CheckName(err, rec.Name, "unexportedProduct")
	t.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine2(
			"{a:100.0 B:2.0}", "{a:-71.0 B:24.0}",
			"{a:-7100.0 B:48.0}", "{a:-7100.0 B:48.0}"),
	})
}

func Test2UnexportedNested(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview2[unexportedNested2, unexportedNested2, unexportedNested2]()
	iv.AddCase(
		unexportedNested2{unexported2: unexported2{a: 2, B: 4}, C: true},
		unexportedNested2{unexported2: unexported2{a: 0, B: 3}, C: false},
		unexportedNested2{unexported2: unexported2{a: 0, B: 12}, C: false},
	)
	iv.AddSolution(unexportedNestedProduct)
	iv.ShowFieldNames()

	rec, err := iv.RunSolution("unexportedNestedProduct")
	t.CheckName(err, rec.Name, "unexportedNestedProduct")
	t.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine2(
			"{unexported2:{a:2.0 B:4.0} C:true}",
			"{unexported2:{a:0.0 B:3.0} C:false}",
			"{unexported2:{a:0.0 B:12.0} C:false}",
			"{unexported2:{a:0.0 B:12.0} C:false}"),
	})
}
