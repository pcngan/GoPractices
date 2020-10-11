package tomato

import (
	"testing"
)

func TestSolution01(t *testing.T){
	cases := []struct{
		inA int
		inB int
		want int
	}{
		{5,8,2},
		{4,10,4},
		{1,2,-1},
		{10,5,1},
	}
	
	for _,c := range cases {
		got := Solution01(c.inA,c.inB)
		if got!= c.want {
			t.Errorf("Solution01(%v, %v) = %v, want = %v", c.inA, c.inB, got, c.want)
		}
	}
}

func TestSolution02(t *testing.T){
	cases := []struct{
		inA string
		want int
	}{
		{"abaaba", 2},
		{"zyzyzyz", 5},
		{"aabbbabaaa",3},
	}
	
	for _,c := range cases {
		got := Solution02(c.inA)
		if got!= c.want {
			t.Errorf("Solution01(%v) = %v, want = %v", c.inA, got, c.want)
		}
	}
}