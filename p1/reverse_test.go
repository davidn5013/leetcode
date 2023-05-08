package p1

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		/*
			Example 1:

			Input: x = 123
			Output: 321

			Example 2:

			Input: x = -123
			Output: -321

			Example 3:

			Input: x = 120
			Output: 21
		*/
		{
			name: "Example 1",
			args: args{x: 123},
			want: 321,
		},
		{
			name: "Example 2",
			args: args{x: -123},
			want: -321,
		},
		{
			name: "Example 3",
			args: args{x: 120},
			want: 21,
		},
		{
			name: "Example extra",
			args: args{x: 1534236469},
			want: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
