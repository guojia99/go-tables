package table

import (
	"fmt"
	"testing"
)

func TestSimpleTable(t *testing.T) {
	opt := &Option{
		Align:   AlignCenter,
		Contour: DefaultContour,
	}
	tests := []struct {
		name    string
		args    interface{}
		wantErr bool
	}{
		{
			name:    "none",
			args:    "none",
			wantErr: true,
		},
		{
			name: "map table",
			args: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
				"key5": "value5",
				"key6": "value6",
			},
			wantErr: false,
		},
		{
			name: "map table 2",
			args: map[string]interface{}{
				"number":  1,
				"string":  "guojia",
				"float":   2.4,
				"slide":   []int{1, 2, 3, 4},
				"complex": complex(1, -1),
				"key6": struct {
					a string
				}{a: "123"},
			},
			wantErr: false,
		},
		{
			name: "map slice table",
			args: map[string][]string{
				"key1": {"value1", "value11", "value12", "value13"},
				"key2": {"value2", "value2", "value2", "value2"},
				"key3": {"value3", "value3", "value3", "value3"},
				"key4": {"value4", "value4", "value4", "value4"},
				"key5": {"value5", "value4", "value4", "value4"},
				"key6": {"value6", "value4", "value4", "value4"},
			},
			wantErr: false,
		},
		{
			name: "slice 2d table",
			args: [][]string{
				{"DATA1", "DATA2", "DATA3"},
				{"DATA4", "DATA5", "DATA6"},
				{"DATA7", "DATA8", "DATA9"},
				{"DATA10", "DATA11", "DATA12"},
			},
			wantErr: false,
		},
		{
			name: "slice table",
			args: []string{
				"DATA1", "DATA2", "DATA3", "DATA4", "DATA5", "DATA6",
			},
			wantErr: false,
		},
		{
			name: "struct table",
			args: struct {
				Str    string
				Val    string `table:"value"`
				Num    int    `json:"number"`
				NoUse  string `json:"-"`
				NoUse2 string `table:"-"`
			}{
				Str:    "value",
				Val:    "val",
				Num:    111,
				NoUse:  "nouse",
				NoUse2: "nouse",
			},
			wantErr: false,
		},
		{
			name: "struct slice table",
			args: []struct {
				Str    string
				Val    string `table:"value"`
				Num    int    `json:"number"`
				NoUse  string `json:"-"`
				NoUse2 string `table:"-"`
			}{
				{"data1", "val1", 1, "no1", "no2"},
				{"data2", "val2", 2, "no2", "no3"},
				{"data3", "val3", 3, "no3", "no4"},
				{"data4", "val4", 4, "no4", "no5"},
				{"data5", "val5", 5, "no5", "no6"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SimpleTable(tt.args, opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimpleTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
