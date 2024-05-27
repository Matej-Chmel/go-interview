package gointerview

import (
	"fmt"
	"reflect"
	"strings"
)

func anyToString(pA interface{}) string {
	return valueToString(reflect.ValueOf(pA))
}

func byteArrayAsString(val reflect.Value) string {
	var builder strings.Builder
	last := val.Len() - 1

	for i := 0; i < last; i++ {
		s := string(byte(val.Index(i).Uint()))
		builder.WriteString(s)
	}

	s := string(byte(val.Index(last).Uint()))
	builder.WriteString(s)
	return builder.String()
}

func runeArrayAsString(val reflect.Value) string {
	var builder strings.Builder
	last := val.Len() - 1

	for i := 0; i < last; i++ {
		s := string(rune(val.Index(i).Int()))
		builder.WriteString(s)
	}

	s := string(rune(val.Index(last).Int()))
	builder.WriteString(s)
	return builder.String()
}

func structToString(val reflect.Value) string {
	nFields := val.NumField()

	if nFields == 0 {
		return "{}"
	}

	var builder strings.Builder
	lastIndex := nFields - 1
	builder.WriteRune('{')

	for i := 0; i < lastIndex; i++ {
		builder.WriteString(fmt.Sprintf("%v, ", val.Field(i)))
	}

	builder.WriteString(fmt.Sprintf("%v}", val.Field(lastIndex)))
	return builder.String()
}

func valueToString(val reflect.Value) string {
	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		kind := val.Type().Elem().Kind()

		if kind == reflect.Uint8 {
			return byteArrayAsString(val)
		}

		if kind == reflect.Int32 {
			return runeArrayAsString(val)
		}
	} else if val.Kind() == reflect.Struct {
		return structToString(val)
	}

	return fmt.Sprintf("%v", val)
}
