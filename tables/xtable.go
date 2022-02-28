package tables

import (
	"errors"
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

type Tables interface {
	Type() Kind
	String() string
	Parse(interface{}) error
}

func NewXTable(data interface{}, opt *Option) (t Tables, err error) {
	kind := parsingType(data)

	switch kind {
	case StructSlice:
		t = &StructSliceTable{Opt: opt}
	case Slice:
		t = &SliceTable{Opt: opt}
	case None:
		return nil, errors.New("the data body required to create a new table frame does not support this type")
	}
	if t == nil {
		return nil, errors.New("unable to create table")
	}
	if err = t.Parse(data); err != nil {
		return nil, err
	}
	return
}

type StructTable struct{}

func (t *StructTable) Type() Kind              { return Struct }
func (t *StructTable) String() string          { return "" }
func (t *StructTable) Parse(interface{}) error { return nil }

type MapTable struct{}

func (t *MapTable) Type() Kind              { return Map }
func (t *MapTable) String() string          { return "" }
func (t *MapTable) Parse(interface{}) error { return nil }

// StructSliceTable create a table from a list of structs
type StructSliceTable struct {
	Opt *Option
	tb  *Table
}

func (t *StructSliceTable) Type() Kind { return StructSlice }
func (t *StructSliceTable) String() (out string) {
	if len(t.tb.Title) == 0 {
		return "error table"
	}
	mws := MaxCellsWidths(t.tb.Title)
	for _, bodyCells := range t.tb.Body {
		for idx, c := range MaxCellsWidths(bodyCells) {
			if mws[idx] < c {
				mws[idx] = c
			}
		}
	}

	header := t.tb.Title.Parse(mws, t.Opt)
	out += t.Opt.Contour.SlideHeader(mws...) + header + t.Opt.Contour.SlideCenter(mws...)
	for _, c := range t.tb.Body {
		out += c.Parse(mws, t.Opt)
	}
	out += t.Opt.Contour.SlideFooter(mws...)
	return
}
func (t *StructSliceTable) Parse(in interface{}) error {
	inValue := reflect.ValueOf(in)
	structs := make([]interface{}, inValue.Len())
	for i := 0; i < inValue.Len(); i++ {
		structs[i] = inValue.Index(i).Interface()
	}

	t.tb = emptyTable(t.Opt)
	for idx, s := range structs {
		val := reflect.ValueOf(s)
		tp := reflect.TypeOf(s)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
			tp = tp.Elem()
		}
		if val.Kind() != reflect.Struct {
			return errors.New("the content of the struct list is not a struct")
		}

		var cells []*Cell
		for n := 0; n < val.NumField(); n++ {
			field := tp.Field(n)
			baseName := field.Name
			if !isHeadCapitalLetters(baseName) {
				continue
			}
			colValue := fmt.Sprintf("%v", valueInterface(val.FieldByName(baseName)))
			tableTag := field.Tag.Get("table")
			jsonTag := field.Tag.Get("json")
			if jsonTag == "-" || tableTag == "-" {
				continue
			}
			if idx == 0 {
				if tableTag != "" {
					baseName = tableTag
				} else if jsonTag != "" {
					baseName = jsonTag
				}
				t.tb.Title = append(t.tb.Title, NewCell(t.Opt.Align, baseName))
			}
			cells = append(cells, NewCell(t.Opt.Align, colValue))
		}
		t.tb.Body = append(t.tb.Body, cells)
	}
	return nil
}

type SliceTable struct {
	Opt *Option
	tb  *Table
}

func (t *SliceTable) Type() Kind     { return Slice }
func (t *SliceTable) String() string { return t.tb.String() }
func (t *SliceTable) Parse(in interface{}) error {
	inValue := reflect.ValueOf(in)
	slice := make([]interface{}, inValue.Len())
	for i := 0; i < inValue.Len(); i++ {
		slice[i] = inValue.Index(i).Interface()
	}

	t.tb = emptyTable(t.Opt)
	t.tb.Title = append(t.tb.Title, NewCell(t.Opt.Align, "No"), NewCell(t.Opt.Align, reflect.TypeOf(inValue).Name()))

	for idx, s := range slice {
		val := reflect.ValueOf(s)
		tp := reflect.TypeOf(s)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
			tp = tp.Elem()
		}
		cells := Cells{
			NewCell(t.Opt.Align, fmt.Sprintf("%d", idx)),
			NewCell(t.Opt.Align, fmt.Sprintf("%v", valueInterface(val))),
		}
		t.tb.Body = append(t.tb.Body, cells)
	}
	return nil
}

type Slice2DTable struct{}

func (t *Slice2DTable) Type() Kind              { return Slice2D }
func (t *Slice2DTable) String() string          { return "" }
func (t *Slice2DTable) Parse(interface{}) error { return nil }

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
