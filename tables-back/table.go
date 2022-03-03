package tables_back

import (
	"errors"
	"fmt"
	"reflect"
)

type Table struct {
	Opt           *Option
	Title, Footer Cells
	Body          []Cells
}

func (t *Table) String() (out string) {
	if len(t.Title) == 0 {
		return "error table"
	}
	mws := MaxCellsWidths(t.Title)
	for _, bodyCells := range t.Body {
		for idx, c := range MaxCellsWidths(bodyCells) {
			if mws[idx] < c {
				mws[idx] = c
			}
		}
	}

	header := t.Title.Parse(mws, t.Opt)
	out += t.Opt.Contour.SlideHeader(mws...) + header + t.Opt.Contour.SlideCenter(mws...)
	for _, c := range t.Body {
		out += c.Parse(mws, t.Opt)
	}
	out += t.Opt.Contour.SlideFooter(mws...)
	return
}

func NewTable(data interface{}, opt *Option) (t *Table, err error) {
	switch parsingType(data) {
	case StructSlice:
		t, err = structSlice2Table(data, opt)
	case Slice:
		t, err = slice2Table(data, opt)
	case Slice2D:
		t, err = slice2D2Table(data, opt)
	case Map:
		t, err = Map2Table(data, opt)
	case Struct:
		t, err = Struct2Table(data, opt)
	case None:
		return nil, errors.New("the data body required to create a new table frame does not support this type")
	}
	return
}

func emptyTable(opt *Option) *Table {
	return &Table{
		Opt:   opt,
		Title: Cells{},
		Body:  []Cells{},
	}
}

func structSlice2Table(in interface{}, opt *Option) (*Table, error) {
	tb := emptyTable(opt)

	inValue := reflect.ValueOf(in)
	structs := make([]interface{}, inValue.Len())
	for i := 0; i < inValue.Len(); i++ {
		structs[i] = inValue.Index(i).Interface()
	}
	for idx, s := range structs {
		val := reflect.ValueOf(s)
		tp := reflect.TypeOf(s)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
			tp = tp.Elem()
		}
		if val.Kind() != reflect.Struct {
			return nil, errors.New("the content of the struct list is not a struct")
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
				tb.Title = append(tb.Title, NewCell(opt.Align, baseName))
			}
			cells = append(cells, NewCell(opt.Align, colValue))
		}
		tb.Body = append(tb.Body, cells)
	}
	return tb, nil
}
func slice2Table(in interface{}, opt *Option) (*Table, error) {
	tb := emptyTable(opt)

	inValue := reflect.ValueOf(in)
	slice := make([]interface{}, inValue.Len())
	for i := 0; i < inValue.Len(); i++ {
		slice[i] = inValue.Index(i).Interface()
	}

	tb.Title = append(tb.Title, NewCell(opt.Align, "No"), NewCell(opt.Align, reflect.TypeOf(inValue).Name()))

	for idx, s := range slice {
		val := reflect.ValueOf(s)
		tp := reflect.TypeOf(s)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
			tp = tp.Elem()
		}
		cells := Cells{
			NewCell(opt.Align, fmt.Sprintf("%d", idx)),
			NewCell(opt.Align, fmt.Sprintf("%v", valueInterface(val))),
		}
		tb.Body = append(tb.Body, cells)
	}
	return tb, nil
}
func slice2D2Table(in interface{}, opt *Option) (*Table, error) { return emptyTable(opt), nil }
func Map2Table(in interface{}, opt *Option) (*Table, error)     { return emptyTable(opt), nil }
func Struct2Table(in interface{}, opt *Option) (*Table, error)  { return emptyTable(opt), nil }
