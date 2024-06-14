# go-interview
Simple library for interview preparation.

## How to use

1. Define one or more functions that attempt to solve an interview question.
2. Create new interview object
3. Add test cases to the object
4. Add solution functions to the object
5. Print results

```go
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
```

### Output
```none
iterativeFactorial
==================
(OK) 1 -> 1
(OK) 2 -> 2
(OK) 3 -> 6
(OK) 4 -> 24
(OK) 5 -> 120
(OK) 6 -> 720

recursiveFactorial
==================
(OK) 1 -> 1
(OK) 2 -> 2
(OK) 3 -> 6
(OK) 4 -> 24
(OK) 5 -> 120
(OK) 6 -> 720
```

## Alternative for two inputs

Some interview questions provide two variables as inputs. The code can then be easily changed to use `goi.NewInterview2` that accepts 3 generic parameters - 2 for inputs and 1 for output.

```go
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
```

### Output

```none
add
===
(OK) 1, 1 -> 2
(OK) 2, 2 -> 4
```
