package p1

import "testing"

func Test_zeroFilledSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 3, 0, 0, 2, 0, 0, 4},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := zeroFilledSubarray(tt.args.nums); got != tt.want {
				t.Errorf("zeroFilledSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
