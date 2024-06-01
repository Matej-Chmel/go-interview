package gointerview_test

import (
	"strings"
	"testing"

	goi "github.com/Matej-Chmel/go-interview"
)

type ExportedPair struct {
	A string
	B string
}

func exportedSolution(p ExportedPair) bool {
	return true
}

func TestStruct(t *testing.T) {
	eIV := goi.NewInterview[ExportedPair, bool]()
	eIV.AddCase(ExportedPair{"abc", "def"}, true)
	eIV.AddSolution(exportedSolution)
	eRes := eIV.RunAllSolutions()
	eABC := strings.Contains(eRes[0].String(), "abc")
	eDEF := strings.Contains(eRes[0].String(), "def")

	if !eABC && !eDEF {
		t.Errorf("Exported members cannot be seen\n%s", eRes[0].String())
	}
}
