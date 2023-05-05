package p1

import "testing"

func Benchmark_removeDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeDuplicates([]int{1, 1, 1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10})
	}
}

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Example 1",
			args: []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			name: "Example 2",
			args: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "My zero example",
			args: []int{},
			want: []int{},
		},
		{
			name: "My example",
			args: []int{1, 1, 1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			got := removeDuplicates(tt.args)

			lencompare := got <= len(tt.want)
			compare := true
			if lencompare {
				for i, v := range tt.want {
					if tt.args[i] != v {
						compare = false
						break
					}
				}
			}

			if !lencompare || !compare {
				t.Errorf("removeDuplicates() = %v, want %v", tt.args, tt.want)
			}

		})
	}
}
