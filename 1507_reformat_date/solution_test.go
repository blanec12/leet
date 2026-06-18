package reformatedate

import "testing"

func TestReformatDate(t *testing.T) {
	tests := []struct {
		name string
		date string
		want string
	}{
		{
			name: "example 1",
			date: "20th Oct 2052",
			want: "2052-10-20",
		},
		{
			name: "example 2",
			date: "6th Jun 1933",
			want: "1933-06-06",
		},
		{
			name: "example 3",
			date: "26th May 1960",
			want: "1960-05-26",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reformatDate(tt.date)

			if got != tt.want {
				t.Errorf("reformatDate(%s) = %s, want %s", tt.date, got, tt.want)
			}
		})
	}
}
