package internal

import (
	"reflect"
	"runtime"
	"sort"
	"strings"
)

// Executes all solution functions from a map
func ExecuteSolutions[T any](
	solutions map[string]T, target func(T) Receipt,
) (res ReceiptSlice) {
	i := 0
	res.Receipts = make([]Receipt, len(solutions))

	for _, sol := range solutions {
		res.Receipts[i] = target(sol)
		i++
	}

	sort.Sort(&res)
	return
}

// Returns name of f without package information
func GetFunctionName(f interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	tokens := strings.Split(name, ".")
	return tokens[len(tokens)-1]
}
