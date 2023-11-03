package tables

import (
	"fmt"
	"testing"

	tbcolor "github.com/guojia99/go-tables/table/color"
)

func Test_table_autoScaling(t *testing.T) {

	has := NewCell(1).SetWordWrap(false)
	not := NewEmptyCell()

	t.Run(
		"ok", func(t *testing.T) {
			t1 := &table{
				body: Cells2D{
					{has, has, has, has, not, not},
					{has, not},
					{has, has, has, not},
					{has, has, has, has, has, not},
					{has},
					{},
					{not, has, not, has},
					{not, not, not, not, not},
					{not, not, not, not, not, not},
				},
			}
			fmt.Println(t1.body)
			fmt.Println("-----------")
			t1.autoScaling()
			fmt.Println(t1.body)
		},
	)

	t.Run(
		"ok2", func(t *testing.T) {
			t2 := &table{}
			fmt.Println(t2.body)
			t2.autoScaling()
			fmt.Println(t2.body)
		},
	)
}

func Test_table_doColWithFn(t1 *testing.T) {
	type fields struct {
		outArea Address
		body    Cells2D
	}
	type args struct {
		fn   func(cell Cell)
		cols []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "all col",
			fields: fields{
				outArea: Address{},
				body: Cells2D{
					{NewCell(1), NewCell(2), NewCell(3)},
					{NewCell(1), NewCell(2), NewCell(3)},
					{NewCell(1), NewCell(2), NewCell(3)},
				},
			},
			args: args{
				fn: func(cell Cell) {
					cell.SetColor(tbcolor.BlueBgRed)
				},
				cols: nil,
			},
			wantErr: false,
		},
		{
			name: "one col",
			fields: fields{
				outArea: Address{},
				body: Cells2D{
					{NewCell(1), NewCell(2), NewCell(3)},
					{NewCell(1), NewCell(2), NewCell(3)},
					{NewCell(1), NewCell(2), NewCell(3)},
				},
			},
			args: args{
				fn: func(cell Cell) {
					cell.SetColor(tbcolor.BlueBgRed)
				},
				cols: []int{1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(
			tt.name, func(t1 *testing.T) {
				t := &table{
					outArea: tt.fields.outArea,
					body:    tt.fields.body,
				}
				if err := t.doColWithFn(tt.args.fn, tt.args.cols); (err != nil) != tt.wantErr {
					t1.Errorf("doColWithFn() error = %v, wantErr %v", err, tt.wantErr)
				}
				fmt.Println(t.body)
			},
		)
	}
}

func Test_table_doRowWithFn(t1 *testing.T) {
	type fields struct {
		outArea Address
		body    Cells2D
	}
	type args struct {
		fn   func(cell Cell)
		rows []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "all row",
			fields: fields{
				outArea: Address{},
				body: Cells2D{
					{NewCell(1), NewCell(2), NewCell(3)},
					{NewCell(1), NewCell(2), NewCell(3)},
					{NewCell(1), NewCell(2), NewCell(3)},
				},
			},
			args: args{
				fn: func(cell Cell) {
					cell.SetColor(tbcolor.BlueBgRed)
				},
				rows: nil,
			},
			wantErr: false,
		},
		{
			name: "once row",
			fields: fields{
				outArea: Address{},
				body: Cells2D{
					{NewCell(1), NewCell(2), NewCell(3)},
					{NewCell(1), NewCell(2), NewCell(3)},
					{NewCell(1), NewCell(2), NewCell(3)},
				},
			},
			args: args{
				fn: func(cell Cell) {
					cell.SetColor(tbcolor.BlueBgRed)
				},
				rows: []int{1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(
			tt.name, func(t1 *testing.T) {
				t := &table{
					outArea: tt.fields.outArea,
					body:    tt.fields.body,
				}
				if err := t.doRowWithFn(tt.args.fn, tt.args.rows); (err != nil) != tt.wantErr {
					t1.Errorf("doRowWithFn() error = %v, wantErr %v", err, tt.wantErr)
				}
				fmt.Println(t.body)
			},
		)
	}
}

func Test_table_doAddressWithFn(t1 *testing.T) {
	type fields struct {
		outArea Address
		body    Cells2D
	}
	type args struct {
		fn      func(cell Cell)
		address Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "once row",
			fields: fields{
				outArea: Address{},
				body: Cells2D{
					{NewCell(1), NewCell(2), NewCell(3)},
					{NewCell(1), NewCell(2), NewCell(3)},
					{NewCell(1), NewCell(2), NewCell(3)},
				},
			},
			args: args{
				fn: func(cell Cell) {
					cell.SetColor(tbcolor.BlueBgRed)
				},
				address: Address{1, 1},
			},
			wantErr: false,
		},
		{
			name: "error number",
			fields: fields{
				outArea: Address{},
				body: Cells2D{
					{NewCell(1), NewCell(2), NewCell(3)},
					{NewCell(1), NewCell(2), NewCell(3)},
					{NewCell(1), NewCell(2), NewCell(3)},
				},
			},
			args: args{
				fn: func(cell Cell) {
					cell.SetColor(tbcolor.BlueBgRed)
				},
				address: Address{10, 1},
			},
			wantErr: true,
		},
		{
			name: "empty table",
			fields: fields{
				outArea: Address{},
				body:    Cells2D{},
			},
			args: args{
				fn: func(cell Cell) {
					cell.SetColor(tbcolor.BlueBgRed)
				},
				address: Address{1, 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(
			tt.name, func(t1 *testing.T) {
				t := &table{
					outArea: tt.fields.outArea,
					body:    tt.fields.body,
				}
				if err := t.doAddressWithFn(tt.args.fn, tt.args.address); (err != nil) != tt.wantErr {
					t1.Errorf("doAddressWithFn() error = %v, wantErr %v", err, tt.wantErr)
				}
				fmt.Println(t.body)
			},
		)
	}
}
