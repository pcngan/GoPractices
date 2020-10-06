package cyclegraph

import "testing"

func TestSolution(t *testing.T){
	cases := []struct {
		inA, inB []int
		want bool
	}{
		{[]int{1,3,2,4}, []int{4,1,3,2}, true},
		{[]int{1,2,3,4}, []int{2,1,4,3}, false},
		{[]int{3,1,2},[]int{2,3,1},true},
		{[]int{1,2,1},[]int{2,3,3},false},
		{[]int{1,2,3,4},[]int{2,1,4,4},false},
	}
	for _, c := range cases {
		got := Solution(c.inA, c.inB)
		if got != c.want {
			t.Errorf("Solution(%v, %v) == %v, want %v", c.inA, c.inB, got, c.want)
		}
	}
}