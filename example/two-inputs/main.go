package main

import goi "github.com/Matej-Chmel/go-interview"

func add(a, b int) int {
	return a + b
}

func main() {
	iv := goi.NewInterview2[int, int, int]()
	iv.AddCase(1, 1, 2)
	iv.AddCase(2, 2, 4)
	iv.AddSolution(add)
	iv.Print()
}
