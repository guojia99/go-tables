package table

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_structTable(t *testing.T) {
	s := map[string]string{
		"aaa":     "addsa",
		"dasdas":  "dasdas",
		"dasdsas": "dasda",
	}
	inValue := reflect.ValueOf(s)
	inType := reflect.TypeOf(s)
	if inValue.Kind() == reflect.Ptr {
		inValue = inValue.Elem()
		inType = inType.Elem()
	}
	for _, val := range inValue.MapKeys() {
		fmt.Println(inValue.MapIndex(val))
	}
}
