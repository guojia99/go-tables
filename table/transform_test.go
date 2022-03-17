package table

import (
	"reflect"
	"testing"
	"time"
)

func TestDefaultTransformContentByTime(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "second ago",
			args: args{
				in: time.Now().Add(-time.Second),
			},
			want: "1 second ago",
		},
		{
			name: "second after",
			args: args{
				in: time.Now().Add(2 * time.Second),
			},
			want: "1 second after",
		},

		{
			name: "minute ago",
			args: args{
				in: time.Now().Add(-time.Minute - time.Second),
			},
			want: "1 minute ago",
		},
		{
			name: "minute after",
			args: args{
				in: time.Now().Add(2 * time.Minute),
			},
			want: "1 minute after",
		},

		{
			name: "hour ago",
			args: args{
				in: time.Now().Add(-time.Hour - time.Minute),
			},
			want: "1 hour ago",
		},
		{
			name: "hour after",
			args: args{
				in: time.Now().Add(time.Hour + time.Minute),
			},
			want: "1 hour after",
		},

		{
			name: "days ago",
			args: args{
				in: time.Now().AddDate(0, 0, -3),
			},
			want: "3 days ago",
		},
		{
			name: "days after",
			args: args{
				in: time.Now().AddDate(0, 0, 3),
			},
			want: "2 days after",
		},

		{
			name: "month ago",
			args: args{
				in: time.Now().AddDate(0, -2, 0),
			},
			want: "1 month ago",
		},
		{
			name: "month after",
			args: args{
				in: time.Now().AddDate(0, 1, 0),
			},
			want: "1 month after",
		},

		{
			name: "year ago",
			args: args{
				in: time.Now().AddDate(-1, 0, 0),
			},
			want: "1 year ago",
		},
		{
			name: "year after",
			args: args{
				in: time.Now().AddDate(1, 0, 0),
			},
			want: "1 year after",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultTransformContentByTime(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultTransformContentByTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
