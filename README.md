# go-interview
Simple generic library for interview preparation.

## Usage
```go
package main

import (
	"fmt"

	goi "github.com/Matej-Chmel/go-interview"
)

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
    it := goi.NewInterview[int, int]()
	it.AddCase(1, 1).AddCase(2, 2).AddCase(3, 6)
	it.AddCase(4, 24).AddCase(5, 120).AddCase(6, 720)

	it.AddSolutions(iterativeFactorial, recursiveFactorial)
    it.Print()
}
```