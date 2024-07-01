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
