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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumSubarraySum(tt.nums, tt.k)

			if got != tt.want {
				t.Errorf("maxinumSubarrySum(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
