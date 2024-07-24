package main

import goi "github.com/Matej-Chmel/go-interview"

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

func main() {
	iv := goi.NewInterview[int, int]()
	iv.AddCase(1, 1)
	iv.AddCase(2, 2)
	iv.AddCase(3, 6)
	iv.AddCase(4, 24)
	iv.AddCase(5, 120)
	iv.AddCase(6, 720)

	iv.AddSolutions(iterativeFactorial, recursiveFactorial)
	iv.Print()
}
