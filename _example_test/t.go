package main

import (
	"fmt"
	"reflect"
)

type Kind int

var kindName = []string{
	"None",
	"Struct",
	"Map",
	"StructSlice",
	"Slice",
	"Slice2D",
}

func (k Kind) String() string {
	return kindName[k]
}

const (
	None        Kind = iota // other
	Struct                  // struct{}
	Map                     // map[interface{}]interface{}
	StructSlice             // []struct{}
	Slice                   // []interface{}
	Slice2D                 // [][]interface{}
)

func parsingType(in interface{}) Kind {
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

func main() {
	type s struct {
		k string
	}

	fmt.Println(parsingType(map[string]string{"aa": ""}))
	fmt.Println(parsingType(s{k: ""}))
	fmt.Println(parsingType([]s{{k: ""}, {k: "12"}}))
	fmt.Println(parsingType([]string{"1", "2"}))
	fmt.Println(parsingType([][]string{{"dasd"}, {"das", "das"}}))
}
