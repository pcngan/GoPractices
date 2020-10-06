package doublemin

import "testing"

func TestSolution(t *testing.T){
	cases := []struct {
		in, want int
	}{
		{14, 19},
		{10, 11},
		{99, 9999},
		{12345, 12999},
		{11111, 11116},
	}
	for _, c := range cases {
		got := Solution(c.in)
		if got != c.want {
			t.Errorf("Solution(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}