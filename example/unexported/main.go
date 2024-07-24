package main

import goi "github.com/Matej-Chmel/go-interview"

type myStruct struct {
	a int
	b float32
	c complex64
}

func sumOfFields(m myStruct) float32 {
	return float32(m.a) + m.b + real(m.c) + imag(m.c)
}

func main() {
	i := goi.NewInterview[myStruct, float32]()
	i.AddCase(myStruct{1, 2.2, 3.3 + 4.4i}, 10.9)
	i.AddSolution(sumOfFields)
	i.ShowFieldNames()
	i.Print()
}
