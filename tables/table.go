package tables

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

type TableOption struct {
	HeaderColors ColorStyles
	DataColors   ColorStyles
	Align        align
	Contour      *Contour
	// The time of the output table is the serialized time
	useTimeEngine  bool
	timeEngineFunc func(time time.Time) string
}

func NewTableOption() *TableOption {
	return &TableOption{
		Align:          AlignLeft,
		HeaderColors:   NewDefaultColorStyles(),
		DataColors:     NewDefaultColorStyles(),
		Contour:        DefaultContour,
		useTimeEngine:  true,
		timeEngineFunc: DefaultSerializationTime,
	}
}

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
	String() string
	Parse(interface{}) error
	//Header() []string
	//Get(int, int) Cell
}

func NewTable(data interface{}, opt *TableOption) (t Tables, err error) {
	kind := parsingType(data)

	switch kind {
	case StructSlice:
		t = &StructSliceTable{opt: opt}
	case None:
	default:
		return nil, errors.New("kind 类型错误")
	}

	if t == nil {
		return nil, errors.New("kind 类型错误")
	}
	if err = t.Parse(data); err != nil {
		return nil, err
	}
	return
}

// basic table
type Table struct {
	Opt     *TableOption
	Headers Cells
	Body1D  Cells
	Body    []Cells
	Footers Cells
}

func (t *Table) String() (out string) {
	out = t.Headers.Parse(t.Opt, nil)
	if len(t.Body) >= 1 {
		mws := t.Body[0].MaxWidths()
		for _, cells := range t.Body {
			for idx, c := range cells.MaxWidths() {
				if mws[idx] < c {
					mws[idx] = c
				}
			}
		}
		for _, c := range t.Body {
			out += c.Parse(t.Opt, mws)
		}
	}
	return
}
func (t *Table) Parse(interface{}) error { return nil }

func EmptyTable() *Table {
	return &Table{
		Headers: Cells{},
		Body1D:  Cells{},
		Body:    []Cells{},
		Footers: Cells{},
	}
}

type StructSliceTable struct {
	opt  *TableOption
	data *Table
}

func (t *StructSliceTable) String() string { return t.data.String() }
func (t *StructSliceTable) Parse(in interface{}) error {
	inValue := reflect.ValueOf(in)
	structs := make([]interface{}, inValue.Len())
	for i := 0; i < inValue.Len(); i++ {
		structs[i] = inValue.Index(i).Interface()
	}

	tableData := EmptyTable()
	tableData.Opt = t.opt
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
			//if tp.Field(n).PkgPath != "" {
			//	continue
			//}
			baseName := tp.Field(n).Name
			if idx == 0 {
				tableData.Headers = append(tableData.Headers, NewCell(t.opt.Align, baseName))
			}
			colValue := fmt.Sprintf("%v", valueInterface(val.FieldByName(baseName)))
			cells = append(cells, NewCell(t.opt.Align, colValue))
		}
		tableData.Body = append(tableData.Body, cells)
	}
	t.data = tableData
	return nil
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
		return in.Slice(0, in.Len()-1)
	case reflect.Invalid, reflect.Chan, reflect.Func, reflect.Struct, reflect.UnsafePointer, reflect.Ptr:
	default:
	}
	return ""
}
