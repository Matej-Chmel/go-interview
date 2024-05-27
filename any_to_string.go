package gointerview

import (
	"fmt"
	"reflect"
	"strings"
)

func anyToString(pA interface{}) string {
	a := reflect.ValueOf(pA)

	if a.Kind() == reflect.Array || a.Kind() == reflect.Slice {
		kind := a.Type().Elem().Kind()

		if kind == reflect.Uint8 {
			return byteArrayAsString(a)
		}

		if kind == reflect.Int32 {
			return runeArrayAsString(a)
		}
	} else if a.Kind() == reflect.Struct {
		return structToString(a)
	}

	return fmt.Sprintf("%v", a)
}

func byteArrayAsString(a reflect.Value) string {
	var builder strings.Builder
	last := a.Len() - 1

	for i := 0; i < last; i++ {
		s := string(byte(a.Index(i).Uint()))
		builder.WriteString(s)
	}

	s := string(byte(a.Index(last).Uint()))
	builder.WriteString(s)
	return builder.String()
}

func runeArrayAsString(a reflect.Value) string {
	var builder strings.Builder
	last := a.Len() - 1

	for i := 0; i < last; i++ {
		s := string(rune(a.Index(i).Int()))
		builder.WriteString(s)
	}

	s := string(rune(a.Index(last).Int()))
	builder.WriteString(s)
	return builder.String()
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
