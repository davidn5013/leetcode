package p1

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			/*
				   Input: nums1 = [1,3], nums2 = [2]
				   Output: 2.00000
				   Explanation: merged array = [1,2,3] and median is 2.
					Example 2:

					Input: nums1 = [1,2], nums2 = [3,4]
					Output: 2.50000
					Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

			*/
			name: "Example 1",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{2},
			},
			want: 2.0,
		},
		{
			name: "Example 2",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
