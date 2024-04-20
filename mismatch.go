package gointerview

import (
	"fmt"
	"reflect"
)

// Stores values of one mismatched field of actual and expected input.
type Mismatch struct {
	Field string
	Left  string
	Right string
}

// Compares two instances of struct Mismatch.
func (m *Mismatch) Equals(o Mismatch) bool {
	return (m.Field == o.Field &&
		m.Left == o.Left &&
		m.Right == o.Right)
}

// Converts mismatch to a string.
func (m Mismatch) String() string {
	return fmt.Sprintf("%s: \"%s\" != \"%s\"", m.Field, m.Left, m.Right)
}

// Checks for mismatch between primitive types.
func checkForMismatch(a, b reflect.Value, fType string) *Mismatch {
	if reflect.DeepEqual(a.Interface(), b.Interface()) {
		return nil
	}

	return &Mismatch{
		Field: fType,
		Left:  fmt.Sprintf("%v", a.Interface()),
		Right: fmt.Sprintf("%v", b.Interface()),
	}
}

// Checks for multiple mismatches between two interfaces.
// Interfaces can be primitive types or structs.
func findMismatches(pA, pB interface{}) []Mismatch {
	a := reflect.ValueOf(pA)
	b := reflect.ValueOf(pB)
	res := []Mismatch{}

	if a.Type() != b.Type() {
		res = append(res, Mismatch{"Type", a.Type().Name(), b.Type().Name()})
	} else if a.Kind() == reflect.Struct {
		nFields := a.NumField()

		for i := 0; i < nFields; i++ {
			res = processComparison(
				a.Field(i), b.Field(i), a.Type().Field(i).Name, res)
		}
	} else {
		res = processComparison(a, b, "Value", res)
	}

	return res
}

// Attempts to find and add a mismatch to the result slice.
func processComparison(
	a, b reflect.Value, fType string, res []Mismatch,
) []Mismatch {
	mismatch := checkForMismatch(a, b, fType)

	if mismatch != nil {
		res = append(res, *mismatch)
	}

	return res
}
