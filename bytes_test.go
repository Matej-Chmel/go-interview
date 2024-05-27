package gointerview_test

import (
	"strings"
	"testing"

	goi "github.com/Matej-Chmel/go-interview"
)

func badSolution(slice []byte) []byte {
	slice[0] = byte('a')
	return slice
}

func byteSolution(slice []byte) int {
	res := 0

	for _, v := range slice {
		if int(v) < 90 {
			res++
		}
	}

	return res
}

func runeSolution(slice []rune) int {
	res := 0

	for _, v := range slice {
		if int(v) < 90 {
			res++
		}
	}

	return res
}

func TestBadBytes(t *testing.T) {
	it := goi.NewInterview[[]byte, []byte]()

	it.AddCase([]byte("helloWORLD"), []byte("helloWORLD"))
	it.AddSolution(badSolution)

	c := NewChecker[int, int]("badSolution", t)
	res, err := it.RunSolution("badSolution")

	if err != nil {
		c.TestError("Function not found", 1)
	}

	r0 := res[0].String()

	if !strings.Contains(r0, "\"aelloWORLD\" != \"helloWORLD\"") {
		c.TestErrorf("Found: %s", r0)
	}
}

func TestBytes(t *testing.T) {
	it := goi.NewInterview[[]byte, int]()

	it.AddCase([]byte("helloWORLD"), 5)
	it.AddCase([]byte("abc"), 0)
	it.AddCase([]byte("  A"), 1)

	it.AddSolution(byteSolution)

	c := NewChecker[int, int]("byteSolution", t)
	res, err := it.RunSolution("byteSolution")

	if err != nil {
		c.TestError("Function not found", 1)
	}

	r0 := res[0].String()
	r1 := res[1].String()
	r2 := res[2].String()

	if !strings.HasPrefix(r0, "byteSolution(helloWORLD) -> 5") {
		c.TestErrorf("Found: %s", r0)
	}

	if !strings.HasPrefix(r1, "byteSolution(abc) -> 0") {
		c.TestErrorf("Found: %s", r0)
	}

	if !strings.HasPrefix(r2, "byteSolution(  A) -> 1") {
		c.TestErrorf("Found: %s", r0)
	}
}

func TestRunes(t *testing.T) {
	it := goi.NewInterview[[]rune, int]()

	it.AddCase([]rune("helloWORLD"), 5)
	it.AddCase([]rune("abc"), 0)
	it.AddCase([]rune("  A"), 1)

	it.AddSolution(runeSolution)

	c := NewChecker[int, int]("runeSolution", t)
	res, err := it.RunSolution("runeSolution")

	if err != nil {
		c.TestError("Function not found", 1)
	}

	r0 := res[0].String()
	r1 := res[1].String()
	r2 := res[2].String()

	if !strings.HasPrefix(r0, "runeSolution(helloWORLD) -> 5") {
		c.TestErrorf("Found: %s", r0)
	}

	if !strings.HasPrefix(r1, "runeSolution(abc) -> 0") {
		c.TestErrorf("Found: %s", r1)
	}

	if !strings.HasPrefix(r2, "runeSolution(  A) -> 1") {
		c.TestErrorf("Found: %s", r2)
	}
}
