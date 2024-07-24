package internal

import (
	"fmt"
	r "reflect"
)

func ConvertToSlice[T any](data string, flags uint, mask uint) T {
	if (flags & mask) == mask {
		var a any = []byte(data)
		converted, ok := a.(T)
		panicIfFailed[T](ok, "byte")
		return converted
	}

	var a any = []rune(data)
	converted, ok := a.(T)
	panicIfFailed[T](ok, "rune")
	return converted
}

func IsByte[T any]() bool {
	var data T
	val := r.ValueOf(data)
	return val.Kind() == r.Slice && val.Type().Elem().Kind() == r.Uint8
}

func panicIfFailed[T any](ok bool, targetType string) {
	if ok {
		return
	}

	var data T
	panic(fmt.Sprintf("Cannot convert type %T to []%s", data, targetType))
}
