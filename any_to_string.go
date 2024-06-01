package gointerview

import (
	"reflect"

	ats "github.com/Matej-Chmel/go-any-to-string"
)

func anyToString(a any) string {
	return ats.AnyToStringCustom(a, options())
}

func options() ats.Options {
	o := ats.NewOptions()
	o.ByteAsString = true
	o.RuneAsString = true
	return o
}

func valueToString(v *reflect.Value) string {
	return ats.ValueToStringCustom(v, options())
}
