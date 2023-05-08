package p1

import (
	"reflect"
	"testing"
)

func Test_getRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				rowIndex: 3,
			},
			want: []int{1, 3, 3, 1},
		},
		{
			name: "Example 2",
			args: args{
				rowIndex: 0,
			},
			want: []int{1},
		},
		{
			name: "Example 3",
			args: args{
				rowIndex: 1,
			},
			want: []int{1, 1},
		},
		{
			name: "My Example 4 rows",
			args: args{
				rowIndex: 4,
			},
			want: []int{1, 4, 6, 4, 1},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := getRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
