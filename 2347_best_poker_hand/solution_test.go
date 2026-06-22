package bestpokerhand

import "testing"

func TestBestHand(t *testing.T) {
	tests := []struct {
		name  string
		ranks []int
		suits []byte
		want  string
	}{
		{
			name:  "example 1",
			ranks: []int{13, 2, 3, 1, 9},
			suits: []byte{'a', 'a', 'a', 'a', 'a'},
			want:  "Flush",
		},
		{
			name:  "example 2",
			ranks: []int{4, 4, 2, 4, 4},
			suits: []byte{'d', 'a', 'a', 'b', 'c'},
			want:  "Three of a Kind",
		},
		{
			name:  "example 3",
			ranks: []int{10, 10, 2, 12, 9},
			suits: []byte{'a', 'b', 'c', 'a', 'd'},
			want:  "Pair",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bestHand(tt.ranks, tt.suits)

			if got != tt.want {
				t.Errorf("bestHand(ranks=%v, suits=%q) = %s, want %s", tt.ranks, tt.suits, got, tt.want)
			}
		})
	}
}
