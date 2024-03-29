package utils

import (
	"fmt"
	"reflect"
)

type Kind int

const (
	None        Kind = iota // other
	Struct                  // struct{}
	Map                     // map[interface{}]interface{}
	MapSlice                // map[interface{}][]interface{}
	StructSlice             // []struct{}
	Slice                   // []interface{}
	Slice2D                 // [][]interface{}
)

func ParsingType(in interface{}) Kind {
	v := reflect.ValueOf(in)
	switch v.Kind() {
	case reflect.Struct:
		return Struct
	case reflect.Map:
		switch v.Type().Elem().Kind() {
		case reflect.Slice:
			return MapSlice
		}
		return Map
	case reflect.Slice:
		p := v.Type().Elem()
		switch p.Kind() {
		case reflect.Struct:
			return StructSlice
		case reflect.Ptr:
			if p.Elem().Kind() == reflect.Struct {
				return StructSlice
			}
			return Slice
		case reflect.Slice:
			return Slice2D
		}
		return Slice
	}
	return None
}

func ValueInterface(in reflect.Value) interface{} {
	interfaceVal := in.Interface()
	if stringer, ok := interfaceVal.(fmt.Stringer); ok {
		return stringer.String()
	}

	switch in.Type().Kind() {
	case reflect.Interface:
		return interfaceVal
	case reflect.Bool:
		return in.Bool()
	case reflect.String:
		return in.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return in.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return in.Uint()
	case reflect.Float32, reflect.Float64:
		return in.Float()
	case reflect.Complex64, reflect.Complex128:
		return in.Complex()
	case reflect.Slice, reflect.Array:
		return in.Slice(0, in.Len())
	case reflect.Invalid, reflect.Chan, reflect.Func, reflect.Struct, reflect.UnsafePointer, reflect.Ptr:
	default:
	}
	return ""
}

func IsHeadCapitalLetters(in string) bool {
	if len(in) == 0 {
		return false
	}
	if !('A' <= in[0] && in[0] <= 'Z') {
		return false
	}
	return true
}
