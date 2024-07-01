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
	i := goi.NewInterview[int, int]()
	i.AddCase(1, 1)
	i.AddCase(2, 2)
	i.AddCase(3, 6)
	i.AddCase(4, 24)
	i.AddCase(5, 120)
	i.AddCase(6, 720)

	i.AddSolutions(iterativeFactorial, recursiveFactorial)
	i.Print()
}
