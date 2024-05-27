package gointerview_test

import (
	"testing"

	goi "github.com/Matej-Chmel/go-interview"
)

func badFactorial(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 6
	case 4:
		return 24
	case 5:
		return 120
	default:
		return n * n
	}
}

func iterativeFactorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func recursiveFactorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * recursiveFactorial(n-1)
}

func TestBad(t *testing.T) {
	it := goi.NewInterview[int, int]()
	it.AddCase(1, 1).AddCase(2, 2).AddCase(3, 6)
	it.AddCase(4, 24).AddCase(5, 120).AddCase(6, 720)

	it.AddSolution(badFactorial)

	c := NewChecker[int, int]("badFactorial", t)
	res, err := it.RunSolution("badFactorial")

	if err != nil {
		c.TestError("Function not found", 1)
	}

	c.AddMismatch("Value", "36", "720")
	c.Check(res[5])
	c.Clear()

	c.CheckSlice(res[0:5])
}

func TestMultiple(t *testing.T) {
	it := goi.NewInterview[int, int]()
	it.AddCase(1, 1).AddCase(2, 2).AddCase(3, 6)
	it.AddCase(4, 24).AddCase(5, 120).AddCase(6, 720)

	it.AddSolutions(iterativeFactorial, recursiveFactorial)

	res := it.RunAllSolutions()

	c := NewChecker[int, int]("iterativeFactorial", t)
	c.CheckSlice(res[0:6])

	c = NewChecker[int, int]("recursiveFactorial", t)
	c.CheckSlice(res[7:])
}
