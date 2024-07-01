# Go Interview
Simple generic library for an interview preparation.

## Guide
The library operates as a test driver.

1. User defines 1 or more solutions to an interview problem
2. User defines a set of test cases as pairs of an input and expected values
3. The library then runs each solution with each input
4. The result is a comparison of actual outputs to the expected ones

Here is an example with a factorial problem.

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

## Two input example
Many interview problems require two inputs. This functionality is provided by a different class `Interview2`.

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

## 2D example
The library supports combinations of 1D and 2D number slices as inputs. Methods for reading one or more test cases from a text file are provided by both `Interview` classes.

The supported format is defined as follows:
- Each line represents 1D slice
- Two or more newlines separate 2D slices
- Each number is separated by one or more whitespace characters
- Boolean values are represented as numbers 0 and 1
- Float values can omit leading zeros before decimal point

Here is an example input file:

```none
1 0 -1
2 -2 3
-1 4 0

-5 3 1 2
4 -4 0 0
0 2 -2 1

3 -1 4
2 0 -3
-2 5 -6
```

And here is a user code that reads the file as multiple test cases.

```go
package main

import goi "github.com/Matej-Chmel/go-interview"

func matrixMult(a, b [][]int8) [][]int8 {
	if len(a) != len(b) {
		return nil
	}

	res := make([][]int8, len(a))

	for i := 0; i < len(a); i++ {
		aRow, bRow := a[i], b[i]

		if len(aRow) != len(bRow) {
			return nil
		}

		res[i] = make([]int8, len(aRow))

		for j := 0; j < len(aRow); j++ {
			res[i][j] = aRow[j] * bRow[j]
		}
	}

	return res
}

func main() {
	i := goi.NewInterview2[[][]int8, [][]int8, [][]int8]()
	i.AddSolution(matrixMult)
	i.ReadCases(
		"test_data/matrixMult_in.txt",
		"test_data/matrixMult_in2.txt",
		"test_data/matrixMult_out.txt")
	i.Print()
}
```

### Output

```none
matrixMult
==========
(OK) 1 0 -1  -3 1 4    -3 0 -4
     2 -2 3, 0 -2 2 -> 0 4 6
     -1 4 0  5 -3 1    -5 -12 0

(OK) -5 3 1 2  1 0 -1 3     -5 0 -1 6
     4 -4 0 0, -3 2 4 -2 -> -12 -8 0 0
     0 2 -2 1  2 -2 1 0     0 -4 -2 0

(OK) 3 -1 4   0 -4 5     0 4 20
     2 0 -3 , -2 3 -1 -> -4 0 3
     -2 5 -6  1 2 0      -2 10 0

(OK) 7 0 0 -1   -1 4 -5 0    -7 0 0 0
     -2 8 -3 2, 3 -3 2 1  -> -6 -24 -6 2
     1 1 -4 -5  2 1 -6 -2    2 1 24 10

(OK) 0 -1 2   3 -3 1    0 3 2
     1 1 0  , -2 0 2 -> -2 0 0
     -3 4 -2  4 -5 3    -12 -20 -6
```