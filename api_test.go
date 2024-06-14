package gointerview_test

import (
	"strings"
	"testing"

	goi "github.com/Matej-Chmel/go-interview"
	ite "github.com/Matej-Chmel/go-interview/internal"
)

func checkLines(actual, expected []ite.ReceiptLine, t *testing.T) {
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

func checkSlice(s *ite.ReceiptSlice, t *testing.T, solNames ...string) {
	if len(s.Receipts) != len(solNames) {
		ite.Throw(t, 2, "Expected %d, found %d solutions", len(s.Receipts), len(solNames))
	}

	for i := 0; i < len(solNames); i++ {
		if a, e := s.Receipts[i].Name, solNames[i]; a != e {
			ite.Throw(t, 2, "Name at index %d does not equal\n%s != %s", i, a, e)
		}
	}
}

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
	ite.CheckName(err, bad.Name, "badFactorial", t)
	checkLines(bad.Lines, []ite.ReceiptLine{
		ite.NewReceiptLine("0", "1", "0"),
		ite.NewReceiptLine("1", "1", "1"),
		ite.NewReceiptLine("3", "2", "2"),
		ite.NewReceiptLine("6", "6", "3"),
		ite.NewReceiptLine("10", "24", "4"),
		ite.NewReceiptLine("15", "120", "5"),
	}, t)

	loop, err := i.RunSolution("loopFactorial")
	ite.CheckName(err, loop.Name, "loopFactorial", t)
	checkLines(loop.Lines, []ite.ReceiptLine{
		ite.NewReceiptLine("1", "1", "0"),
		ite.NewReceiptLine("1", "1", "1"),
		ite.NewReceiptLine("2", "2", "2"),
		ite.NewReceiptLine("6", "6", "3"),
		ite.NewReceiptLine("24", "24", "4"),
		ite.NewReceiptLine("120", "120", "5"),
	}, t)

	rec, err := i.RunSolution("recursiveFactorial")
	ite.CheckName(err, rec.Name, "recursiveFactorial", t)
	checkLines(rec.Lines, []ite.ReceiptLine{
		ite.NewReceiptLine("1", "1", "0"),
		ite.NewReceiptLine("1", "1", "1"),
		ite.NewReceiptLine("2", "2", "2"),
		ite.NewReceiptLine("6", "6", "3"),
		ite.NewReceiptLine("24", "24", "4"),
		ite.NewReceiptLine("120", "120", "5"),
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
	input := []byte("hello")

	i.AddCase(input, []byte("Aello"))
	i.AddSolution(bytes)

	s := i.RunAllSolutions()
	data := string(input)

	if data != "hello" {
		ite.Throw(t, 1, "Input changed from hello to %s", data)
	}

	checkSlice(&s, t, "bytes")

	if d := len(s.Receipts[0].Lines); d != 1 {
		ite.Throw(t, 1, "Number of lines: %d", d)
	}

	if a := s.Receipts[0].Lines[0].Actual; a != "Aello" {
		ite.Throw(t, 1, "%s != Aello", a)
	}
}

func TestRunes(t *testing.T) {
	i := goi.NewInterview[[]rune, []rune]()
	input := []rune("hello")

	i.AddCase(input, []rune("Aello"))
	i.AddSolution(runes)

	s := i.RunAllSolutions()
	data := string(input)

	if data != "hello" {
		ite.Throw(t, 1, "Input changed from hello to %s", data)
	}

	checkSlice(&s, t, "runes")

	if d := len(s.Receipts[0].Lines); d != 1 {
		ite.Throw(t, 1, "Number of lines: %d", d)
	}

	if a := s.Receipts[0].Lines[0].Actual; a != "Aello" {
		ite.Throw(t, 1, "%s != Aello", a)
	}
}

func solutionOne(i int) int {
	return i
}

func solutionTwo(i int) int {
	return i + 1
}

func TestWrite(t *testing.T) {
	i := goi.NewInterview[int, int]()

	i.AddCase(1, 2)
	i.AddCase(65, 66)

	i.AddSolutions(solutionOne, solutionTwo)

	var builder strings.Builder
	i.WriteAllSolutions(&builder)

	actual := builder.String()
	expected := ite.ReadFile("test_data/test_write.txt", t)

	if actual != expected {
		ite.Throw(t, 1, "\nActual:\n%s\n\nExpected:\n%s", actual, expected)
		ite.Throw(t, 1, "Actual: %d, Expected: %d", len(actual), len(expected))
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
