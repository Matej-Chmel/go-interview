package gointerview_test

import (
	"sort"
	"testing"

	goi "github.com/Matej-Chmel/go-interview"
	ite "github.com/Matej-Chmel/go-interview/internal"
)

type unexported struct {
	a int
	B int
}

type unexportedNested struct {
	unexported
	C int
}

func badFactorial(n int) (r int) {
	for i := 1; i <= n; i++ {
		r += i
	}

	return
}

func badSort(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	nums[0] = 0
	return nums
}

func bytes(s []byte) []byte {
	s[0] = 65
	return s
}

func exportedNestedSolution(e ite.ExportedNested) ite.ExportedNested {
	return ite.ExportedNested{
		Exported: ite.Exported{A: e.A + 1, B: e.B + 2},
		C:        100,
	}
}

func exportedSolution(e ite.Exported) ite.Exported {
	return ite.Exported{A: 1, B: 2}
}

func inc(i int) int {
	return i + 1
}

func incMatrix(a [][]int32) [][]int32 {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j]++

			if j >= 3 {
				a[i][j]++
			}
		}
	}
	return a
}

func loopFactorial(n int) (r int) {
	r = 1

	for i := 1; i <= n; i++ {
		r *= i
	}

	return
}

func noInc(i int) int {
	return i
}

func recursiveFactorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * recursiveFactorial(n-1)
}

func runes(s []rune) []rune {
	s[0] = 'A'
	return s
}

func unexportedDouble(e unexported) unexported {
	return unexported{a: e.a * 2, B: e.B * 2}
}

func unexportedNestedInc(e unexportedNested) unexportedNested {
	return unexportedNested{
		unexported: unexported{a: e.a + 1, B: e.B + 2},
		C:          100,
	}
}

func TestBytes(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview[[]byte, []byte]()
	iv.ShowBytesAsString()
	input := []byte("hello")

	// Test that input is deep copied
	iv.AddCase(input, []byte("Aello"))

	// Test that conversion from string works
	iv.AddCaseString("abcd", "Abcd")
	iv.AddSolution(bytes)

	rec := iv.RunAllSolutions()
	data := string(input)

	if data != "hello" {
		t.Throw(1, "Input changed from hello to %s", data)
		return
	}

	if len(rec.Receipts) != 1 {
		t.Throw(1, "More receipts (%d)", len(rec.Receipts))
		return
	}

	t.CheckSlice(&rec, "bytes")
	t.CheckLines(rec.Receipts[0].Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine("hello", "Aello", "Aello"),
		ite.NewReceiptLine("abcd", "Abcd", "Abcd"),
	})
}

func TestExported(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview[ite.Exported, ite.Exported]()
	data := ite.Exported{A: 1, B: 2}
	iv.AddCase(data, ite.Exported{A: 3, B: 4})
	iv.AddSolution(exportedSolution)

	rec, err := iv.RunSolution("exportedSolution")
	t.CheckName(err, rec.Name, "exportedSolution")
	t.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine("{1 2}", "{1 2}", "{3 4}"),
	})
}

func TestExportedNested(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview[ite.ExportedNested, ite.ExportedNested]()
	iv.AddCase(
		ite.ExportedNested{Exported: ite.Exported{A: 50, B: 100}, C: 3},
		ite.ExportedNested{Exported: ite.Exported{A: 51, B: 102}, C: 100})
	iv.AddSolution(exportedNestedSolution)

	rec, err := iv.RunSolution("exportedNestedSolution")
	t.CheckName(err, rec.Name, "exportedNestedSolution")
	t.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine("50/100/3", "51/102/100", "51/102/100"),
	})
}

func TestFactorial(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview[int, int]()

	iv.AddCase(0, 1)
	iv.AddCase(1, 1)
	iv.AddCase(2, 2)
	iv.AddCase(3, 6)
	iv.AddCase(4, 24)
	iv.AddCase(5, 120)

	iv.AddSolutions(badFactorial, loopFactorial, recursiveFactorial)

	bad, err := iv.RunSolution("badFactorial")
	t.CheckName(err, bad.Name, "badFactorial")
	t.CheckLines(bad.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine("0", "0", "1"),
		ite.NewReceiptLine("1", "1", "1"),
		ite.NewReceiptLine("2", "3", "2"),
		ite.NewReceiptLine("3", "6", "6"),
		ite.NewReceiptLine("4", "10", "24"),
		ite.NewReceiptLine("5", "15", "120"),
	})

	loop, err := iv.RunSolution("loopFactorial")
	t.CheckName(err, loop.Name, "loopFactorial")
	t.CheckLines(loop.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine("0", "1", "1"),
		ite.NewReceiptLine("1", "1", "1"),
		ite.NewReceiptLine("2", "2", "2"),
		ite.NewReceiptLine("3", "6", "6"),
		ite.NewReceiptLine("4", "24", "24"),
		ite.NewReceiptLine("5", "120", "120"),
	})

	rec, err := iv.RunSolution("recursiveFactorial")
	t.CheckName(err, rec.Name, "recursiveFactorial")
	t.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine("0", "1", "1"),
		ite.NewReceiptLine("1", "1", "1"),
		ite.NewReceiptLine("2", "2", "2"),
		ite.NewReceiptLine("3", "6", "6"),
		ite.NewReceiptLine("4", "24", "24"),
		ite.NewReceiptLine("5", "120", "120"),
	})
}

func TestIncMatrix(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview[[][]int32, [][]int32]()
	iv.AddSolution(incMatrix)
	iv.ReadCases("test_data/incMatrix_in.txt", "test_data/incMatrix_out.txt")

	if expected, err := ite.ReadAllText("test_data/incMatrix_stdout.txt"); err != nil {
		t.Throw(1, err.Error())
	} else {
		t.CheckStrings(1, iv.AllSolutionsToString(), expected)
	}
}

func TestNil(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview[*ite.ExportedNested, *ite.ExportedNested]()
	iv.AddCase(nil, nil)
	iv.AddSolution(func(e *ite.ExportedNested) *ite.ExportedNested {
		return nil
	})

	rec, err := iv.RunSolution("func1")
	t.CheckName(err, rec.Name, "func1")
	t.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine("nil", "nil", "nil"),
	})
}

func TestNoData(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview[bool, int]()

	t.CheckStrings(1, iv.AllSolutionsToString(),
		"No test cases provided by the user!")

	iv.AddCase(false, 0)
	t.CheckStrings(1, iv.AllSolutionsToString(),
		"No solution functions provided by the user!")
}

func TestRunes(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview[[]rune, []rune]()
	iv.ShowBytesAsString()
	input := []rune("hello")

	iv.AddCase(input, []rune("Aello"))
	iv.AddSolution(runes)

	s := iv.RunAllSolutions()
	data := string(input)

	if data != "hello" {
		t.Throw(1, "Input changed from hello to %s", data)
		return
	}

	t.CheckSlice(&s, "runes")

	if d := len(s.Receipts[0].Lines); d != 1 {
		t.Throw(1, "Number of lines: %d", d)
		return
	}

	if a := s.Receipts[0].Lines[0].Actual; a != "Aello" {
		t.Throw(1, "%s != Aello", a)
	}
}

func TestSort1D(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview[[]int, []int]()
	iv.AddSolution(badSort)
	iv.ReadCases("test_data/sort_in.txt", "test_data/sort_out.txt")

	rec, err := iv.RunSolution("badSort")
	t.CheckName(err, rec.Name, "badSort")
	t.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine("[1 3 5 7 9]", "[0 3 5 7 9]", "[1 3 5 7 9]"),
		ite.NewReceiptLine("[9 0 7 8 9]", "[0 7 8 9 9]", "[0 7 8 9 9]"),
		ite.NewReceiptLine("[3 3 3 2 2]", "[0 2 3 3 3]", "[2 2 3 3 3]"),
		ite.NewReceiptLine("[0 0 2 -1 -3 -2]", "[0 -2 -1 0 0 2]", "[-3 -2 -1 0 0 2]"),
		ite.NewReceiptLine("[51236 3237 908 -90000 90100]", "[0 908 3237 51236 90100]", "[-90000 908 3237 51236 90100]"),
	})
}

func TestStdout(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview[int, int]()
	iv.AddCase(1, 2)
	iv.AddCase(65, 66)
	iv.AddSolutions(noInc, inc)

	if expected, err := ite.ReadAllText("test_data/inc_stdout.txt"); err != nil {
		t.Throw(1, err.Error())
	} else {
		t.CheckStrings(1, iv.AllSolutionsToString(), expected)
	}
}

func TestUnexported(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview[unexported, unexported]()
	iv.AddCase(
		unexported{a: 3, B: 5},
		unexported{a: 6, B: 10})
	iv.AddSolution(unexportedDouble)
	iv.ShowFieldNames()

	rec, err := iv.RunSolution("unexportedDouble")
	t.CheckName(err, rec.Name, "unexportedDouble")
	t.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine("{a:3 B:5}", "{a:6 B:10}", "{a:6 B:10}"),
	})
}

func TestUnexportedNested(ot *testing.T) {
	t := ite.NewTester(ot)
	iv := goi.NewInterview[unexportedNested, unexportedNested]()
	iv.AddCase(
		unexportedNested{unexported: unexported{a: 1, B: 2}, C: 3},
		unexportedNested{unexported: unexported{a: 2, B: 4}, C: 100})
	iv.AddSolution(unexportedNestedInc)
	iv.ShowFieldNames()

	rec, err := iv.RunSolution("unexportedNestedInc")
	t.CheckName(err, rec.Name, "unexportedNestedInc")
	t.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine(
			"{unexported:{a:1 B:2} C:3}",
			"{unexported:{a:2 B:4} C:100}",
			"{unexported:{a:2 B:4} C:100}"),
	})
}
