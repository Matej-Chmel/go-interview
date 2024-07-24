package main

import goi "github.com/Matej-Chmel/go-interview"

func flipCase(data []byte) []byte {
	for i, v := range data {
		if v >= 'A' && v <= 'Z' {
			data[i] = v + 32
		} else if v >= 'a' && v <= 'z' {
			data[i] = v - 32
		}
	}

	return data
}

func main() {
	i := goi.NewInterview[[]byte, []byte]()
	i.AddCaseString("Hello world", "hELLO WORLD")
	i.AddSolution(flipCase)
	i.ShowBytesAsString()
	i.Print()
}
