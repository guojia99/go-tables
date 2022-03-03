package tables_back

import "reflect"

type kind int

const (
	None        kind = iota // other
	Struct                  // struct{}
	Map                     // map[interface{}]interface{}
	StructSlice             // []struct{}
	Slice                   // []interface{}
	Slice2D                 // [][]interface{}
)

func parsingType(in interface{}) kind {
	v := reflect.ValueOf(in)
	switch v.Kind() {
	case reflect.Struct:
		return Struct
	case reflect.Map:
		return Map
	case reflect.Slice:
		switch v.Type().Elem().Kind() {
		case reflect.Struct:
			return StructSlice
		case reflect.Slice:
			return Slice2D
		default:
			return Slice
		}
	}
	return None
}
func valueInterface(in reflect.Value) interface{} {
	switch in.Type().Kind() {
	case reflect.Interface:
		return in.Interface()
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

func isHeadCapitalLetters(in string) bool {
	if len(in) == 0 {
		return false
	}
	if !('A' <= in[0] && in[0] <= 'Z') {
		return false
	}
	return true
}
