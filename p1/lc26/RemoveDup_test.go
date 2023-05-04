package lc26

import "testing"

func Benchmark_removeDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeDuplicates([]int{1, 1, 1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{[]int{1, 1, 2}},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want: 5,
		},
		{
			name: "My zero example",
			args: args{[]int{}},
			want: 0,
		},
		{
			name: "My example",
			args: args{[]int{1, 1, 1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10}},
			want: 10,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
