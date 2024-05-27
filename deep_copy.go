package gointerview

import "reflect"

func deepCopy[T any](v T) T {
	return deepCopyImpl(reflect.ValueOf(v)).Interface().(T)
}

func deepCopyImpl(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		copyPtr := reflect.New(v.Elem().Type())
		copyPtr.Elem().Set(deepCopyImpl(v.Elem()))
		return copyPtr

	case reflect.Interface:
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		copyInterface := deepCopyImpl(v.Elem())
		return copyInterface

	case reflect.Struct:
		copyStruct := reflect.New(v.Type()).Elem()
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).CanSet() {
				copyStruct.Field(i).Set(deepCopyImpl(v.Field(i)))
			}
		}
		return copyStruct

	case reflect.Slice:
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		copySlice := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
		for i := 0; i < v.Len(); i++ {
			copySlice.Index(i).Set(deepCopyImpl(v.Index(i)))
		}
		return copySlice

	case reflect.Map:
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		copyMap := reflect.MakeMapWithSize(v.Type(), v.Len())
		for _, key := range v.MapKeys() {
			copyMap.SetMapIndex(deepCopyImpl(key), deepCopyImpl(v.MapIndex(key)))
		}
		return copyMap

	case reflect.Array:
		copyArray := reflect.New(v.Type()).Elem()
		for i := 0; i < v.Len(); i++ {
			copyArray.Index(i).Set(deepCopyImpl(v.Index(i)))
		}
		return copyArray

	default:
		return v
	}
}
