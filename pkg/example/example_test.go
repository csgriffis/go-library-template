package example

import "testing"

func TestAdd(t *testing.T) {
	var tt = []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{3, 4, 7},
	}

	for _, tc := range tt {
		got := Add(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Add(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}
