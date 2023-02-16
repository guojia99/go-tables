package table

import (
	`fmt`
	`reflect`
)

func valueInterface(in reflect.Value) interface{} {
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
	case reflect.Func:
		return in.Type().Name()
	case reflect.Struct:
		return fmt.Sprintf("%+v", interfaceVal)
	case reflect.Invalid, reflect.Chan, reflect.UnsafePointer, reflect.Ptr:
	default:
	}
	return ""
}

func structTagName(tag reflect.StructTag) (string, bool) {
	tableTag := tag.Get("table")
	jsonTag := tag.Get("json")
	yamlTag := tag.Get("yaml")
	if jsonTag == "-" || tableTag == "-" {
		return "", false
	}
	switch {
	case tableTag != "":
		return tableTag, true
	case jsonTag != "":
		return jsonTag, true
	case yamlTag != "":
		return yamlTag, true
	}
	return "", false
}
