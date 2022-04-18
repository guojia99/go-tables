package table

import "testing"

func TestTable_String(t1 *testing.T) {
	type fields struct {
		Opt     *Option
		Headers RowCell
		Body    []RowCell
		Footers RowCell
	}
	tests := []struct {
		name    string
		fields  fields
		wantOut string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Table{
				Opt:     tt.fields.Opt,
				Headers: tt.fields.Headers,
				Body:    tt.fields.Body,
				Footers: tt.fields.Footers,
			}
			if gotOut := t.String(); gotOut != tt.wantOut {
				t1.Errorf("String() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
