package main

import goi "github.com/Matej-Chmel/go-interview"

func add(a, b int) int {
	return a + b
}

func main() {
	i := goi.NewInterview2[int, int, int]()
	i.AddCase(1, 1, 2)
	i.AddCase(2, 2, 4)
	i.AddSolution(add)
	i.Print()
}
