/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import (
	"fmt"
	"reflect"
)

func valueInterface(in reflect.Value) interface{} {
	interfaceVal := in.Interface()
	switch data := interfaceVal.(type) {
	case fmt.Stringer:
		return data.String()
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
	if tableTag != "-" || jsonTag != "-" {
		switch {
		case tableTag != "":
			return tableTag, true
		case jsonTag != "":
			return jsonTag, true
		}
	}
	return "", false
}

func valueOf(in interface{}) (out reflect.Value, err error) {
	switch in.(type) {
	case reflect.Value:
		out = in.(reflect.Value)
	default:
		out = reflect.ValueOf(in)
	}
	if !out.IsValid() {
		err = NotValidValue
	}
	return
}

type (
	sortMapKeyValues []sortMapKeyValue
	sortMapKeyValue  struct {
		key   Cell
		value interface{}
	}
)

func (s sortMapKeyValues) Len() int           { return len(s) }
func (s sortMapKeyValues) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortMapKeyValues) Less(i, j int) bool { return s[i].key.String() < s[j].key.String() }
