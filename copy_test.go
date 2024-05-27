package gointerview_test

import (
	"testing"

	goi "github.com/Matej-Chmel/go-interview"
)

func correctSolution(slice []byte) []byte {
	return slice
}

func wrongSolution(slice []byte) []byte {
	slice[0] = byte('a')
	return slice
}

func TestCopy(t *testing.T) {
	i := goi.NewInterview[[]byte, []byte]()

	i.AddCase([]byte("hello"), []byte("hello"))
	i.AddSolutions(wrongSolution, correctSolution)

	wChecker := NewChecker[[]byte, []byte]("wrongSolution", t)
	wChecker.AddMismatch("Value", "aello", "hello")

	cChecker := NewChecker[[]byte, []byte]("correctSolution", t)

	wRes, _ := i.RunSolution("wrongSolution")
	wChecker.CheckSlice(wRes)

	cRes, _ := i.RunSolution("correctSolution")
	cChecker.CheckSlice(cRes)
}
