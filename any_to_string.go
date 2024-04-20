package gointerview

import (
	"fmt"
	"reflect"
	"strings"
)

func anyToString(pA interface{}) string {
	a := reflect.ValueOf(pA)

	if a.Kind() == reflect.Struct {
		return structToString(a)
	}

	return fmt.Sprintf("%v", a)
}

func structToString(a reflect.Value) string {
	nFields := a.NumField()

	if nFields == 0 {
		return "{}"
	}

	var builder strings.Builder
	lastIndex := nFields - 1
	builder.WriteRune('{')

	for i := 0; i < lastIndex; i++ {
		builder.WriteString(fmt.Sprintf("%v, ", a.Field(i)))
	}

	builder.WriteString(fmt.Sprintf("%v}", a.Field(lastIndex)))
	return builder.String()
}
