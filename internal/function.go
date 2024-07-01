package internal

import (
	"reflect"
	"runtime"
	"strings"
)

// Returns name of f without package information
func GetFunctionName(f interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	tokens := strings.Split(name, ".")
	return tokens[len(tokens)-1]
}
