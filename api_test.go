package gointerview_test

import (
	"sort"
	"testing"

	goi "github.com/Matej-Chmel/go-interview"
	ite "github.com/Matej-Chmel/go-interview/internal"
)

func badFactorial(n int) (r int) {
	for i := 1; i <= n; i++ {
		r += i
	}

	return
}

func loopFactorial(n int) (r int) {
	r = 1

	for i := 1; i <= n; i++ {
		r *= i
	}

	return
}

func recursiveFactorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * recursiveFactorial(n-1)
}

func TestFactorial(t *testing.T) {
	i := goi.NewInterview[int, int]()

	i.AddCase(0, 1)
	i.AddCase(1, 1)
	i.AddCase(2, 2)
	i.AddCase(3, 6)
	i.AddCase(4, 24)
	i.AddCase(5, 120)

	i.AddSolutions(badFactorial, loopFactorial, recursiveFactorial)

	bad, err := i.RunSolution("badFactorial")
	success := ite.CheckName(err, bad.Name, "badFactorial", t)

	if !success {
		return
	}

	success = ite.CheckLines(bad.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine("0", "0", "1"),
		ite.NewReceiptLine("1", "1", "1"),
		ite.NewReceiptLine("2", "3", "2"),
		ite.NewReceiptLine("3", "6", "6"),
		ite.NewReceiptLine("4", "10", "24"),
		ite.NewReceiptLine("5", "15", "120"),
	}, t)

	if !success {
		return
	}

	loop, err := i.RunSolution("loopFactorial")
	success = ite.CheckName(err, loop.Name, "loopFactorial", t)

	if !success {
		return
	}

	success = ite.CheckLines(loop.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine("0", "1", "1"),
		ite.NewReceiptLine("1", "1", "1"),
		ite.NewReceiptLine("2", "2", "2"),
		ite.NewReceiptLine("3", "6", "6"),
		ite.NewReceiptLine("4", "24", "24"),
		ite.NewReceiptLine("5", "120", "120"),
	}, t)

	if !success {
		return
	}

	rec, err := i.RunSolution("recursiveFactorial")
	success = ite.CheckName(err, rec.Name, "recursiveFactorial", t)

	if !success {
		return
	}

	ite.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine("0", "1", "1"),
		ite.NewReceiptLine("1", "1", "1"),
		ite.NewReceiptLine("2", "2", "2"),
		ite.NewReceiptLine("3", "6", "6"),
		ite.NewReceiptLine("4", "24", "24"),
		ite.NewReceiptLine("5", "120", "120"),
	}, t)
}

func bytes(s []byte) []byte {
	s[0] = 65
	return s
}

func runes(s []rune) []rune {
	s[0] = 'A'
	return s
}

func TestBytes(t *testing.T) {
	i := goi.NewInterview[[]byte, []byte]()
	i.BytesAsString()
	input := []byte("hello")

	i.AddCase(input, []byte("Aello"))
	i.AddSolution(bytes)

	s := i.RunAllSolutions()
	data := string(input)

	if data != "hello" {
		ite.Throw(t, 1, "Input changed from hello to %s", data)
		return
	}

	success := ite.CheckSlice(&s, t, "bytes")

	if !success {
		return
	}

	if d := len(s.Receipts[0].Lines); d != 1 {
		ite.Throw(t, 1, "Number of lines: %d", d)
		return
	}

	if a := s.Receipts[0].Lines[0].Actual; a != "Aello" {
		ite.Throw(t, 1, "%s != Aello", a)
	}
}

func TestRunes(t *testing.T) {
	i := goi.NewInterview[[]rune, []rune]()
	i.BytesAsString()
	input := []rune("hello")

	i.AddCase(input, []rune("Aello"))
	i.AddSolution(runes)

	s := i.RunAllSolutions()
	data := string(input)

	if data != "hello" {
		ite.Throw(t, 1, "Input changed from hello to %s", data)
		return
	}

	success := ite.CheckSlice(&s, t, "runes")

	if !success {
		return
	}

	if d := len(s.Receipts[0].Lines); d != 1 {
		ite.Throw(t, 1, "Number of lines: %d", d)
		return
	}

	if a := s.Receipts[0].Lines[0].Actual; a != "Aello" {
		ite.Throw(t, 1, "%s != Aello", a)
	}
}

func noInc(i int) int {
	return i
}

func inc(i int) int {
	return i + 1
}

func TestStdout(t *testing.T) {
	i := goi.NewInterview[int, int]()
	i.AddCase(1, 2)
	i.AddCase(65, 66)
	i.AddSolutions(noInc, inc)

	if expected, err := ite.ReadAllText("test_data/inc_stdout.txt"); err != nil {
		ite.Throw(t, 1, err.Error())
	} else {
		ite.CheckStrings(t, 1, i.AllSolutionsToString(), expected)
	}
}

func TestExported(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			ite.Throw(t, 1, "TestExported did panic but should have NOT")
		}
	}()

	i := goi.NewInterview[ite.Exported, ite.Exported]()
	data := ite.Exported{A: 1, B: 2}
	i.AddCase(data, data)
}

func TestExportedNested(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			ite.Throw(t, 1, "TestExportedNested did panic but should have NOT")
		}
	}()

	i := goi.NewInterview[ite.ExportedNested, ite.ExportedNested]()
	data := ite.ExportedNested{Exported: ite.Exported{A: 1, B: 2}, C: 3}
	i.AddCase(data, data)
}

type Unexported struct {
	a int
	B int
}

func TestUnexported(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			ite.Throw(t, 1, "TestUnexported did NOT panic")
		}
	}()

	i := goi.NewInterview[Unexported, Unexported]()
	data := Unexported{a: 1, B: 2}
	i.AddCase(data, data)
}

type UnexportedNested struct {
	U Unexported
	C int
}

func TestUnexportedNested(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			ite.Throw(t, 1, "TestUnexportedNested did NOT panic")
		}
	}()

	i := goi.NewInterview[UnexportedNested, UnexportedNested]()
	data := UnexportedNested{U: Unexported{a: 1, B: 2}, C: 3}
	i.AddCase(data, data)
}

func badSort(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	nums[0] = 0
	return nums
}

func TestSort1D(t *testing.T) {
	i := goi.NewInterview[[]int, []int]()
	i.AddSolution(badSort)
	i.ReadCases("test_data/sort_in.txt", "test_data/sort_out.txt")

	rec, err := i.RunSolution("badSort")
	success := ite.CheckName(err, rec.Name, "badSort", t)

	if !success {
		return
	}

	success = ite.CheckLines(rec.Lines, []*ite.ReceiptLine{
		ite.NewReceiptLine("[1 3 5 7 9]", "[0 3 5 7 9]", "[1 3 5 7 9]"),
		ite.NewReceiptLine("[9 0 7 8 9]", "[0 7 8 9 9]", "[0 7 8 9 9]"),
		ite.NewReceiptLine("[3 3 3 2 2]", "[0 2 3 3 3]", "[2 2 3 3 3]"),
		ite.NewReceiptLine("[0 0 2 -1 -3 -2]", "[0 -2 -1 0 0 2]", "[-3 -2 -1 0 0 2]"),
		ite.NewReceiptLine("[51236 3237 908 -90000 90100]", "[0 908 3237 51236 90100]", "[-90000 908 3237 51236 90100]"),
	}, t)

	if !success {
		return
	}
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

func TestIncMatrix(t *testing.T) {
	i := goi.NewInterview[[][]int32, [][]int32]()
	i.AddSolution(incMatrix)
	i.ReadCases("test_data/incMatrix_in.txt", "test_data/incMatrix_out.txt")

	if expected, err := ite.ReadAllText("test_data/incMatrix_stdout.txt"); err != nil {
		ite.Throw(t, 1, err.Error())
	} else {
		ite.CheckStrings(t, 1, i.AllSolutionsToString(), expected)
	}
}
