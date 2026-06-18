package maxgoodsubarraysum

import "testing"

func TestMaximumSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{
			name: "example 1",
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    1,
			want: 11,
		},
		{
			name: "example 2",
			nums: []int{-1, 3, 2, 4, 5},
			k:    3,
			want: 11,
		},
		{
			name: "example 3",
			nums: []int{-1, -2, -3, -4},
			k:    2,
			want: -6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumSubarraySum(tt.nums, tt.k)

			if got != tt.want {
				t.Errorf("maximumSubarrySum(nums=%v, k=%d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
