package dialnumber

import "testing"

func TestSolution(t *testing.T){
	cases := []struct {
		inA, inB []string
		inP, want string
	}{
		{[]string{"pim", "pom"}, []string{"999999999", "777888999"}, "88999", "pom"},
		{[]string{"sander", "amy", "ann", "michael"}, []string{"123456789","234567890","789123456", "123123123"}, "1", "ann"},
		{[]string{"adam", "eva", "leo"},[]string{"121212121", "111111111", "444555666"},"112","NO CONTACT"},
	}
	for _, c := range cases {
		got := Solution(c.inA, c.inB, c.inP)
		if got != c.want {
			t.Errorf("Solution(%v, %v, %v) == %v, want %v", c.inA, c.inB, c.inP, got, c.want)
		}
	}
}