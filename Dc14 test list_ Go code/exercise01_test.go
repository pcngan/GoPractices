package exercise

import (
	"testing"
	"fmt"
)

func TestSolution01(t *testing.T){
	cases := []struct{
		inA []int
		want int
	}{
		{[]int{3,2,1,1,2,3,1}, 5},
		{[]int{4,1,4,1},6},
		{[]int{3,3,3},0},
	}
	
	for _,c := range cases {
		got := Solution01(c.inA)
		if got!= c.want {
			t.Errorf("Solution01(%v) = %v, want = %v", c.inA, got, c.want)
		}
	}
}

func TestSolution02(t *testing.T){
	cases := []struct{
		inN int
		inA, inB []int
		want bool
	}{
		{4, []int{1,2,4,4,3}, []int{2,3,1,3,1}, true},
		{4, []int{1,2,1,3}, []int{2,4,3,4}, false},
		{6, []int{2,4,5,3}, []int{3,5,6,4}, false},
		{3, []int{1,3}, []int{2,2}, true},
	}
	
	for _,c := range cases {
		got := Solution02(c.inN, c.inA, c.inB)
		if got!= c.want {
			t.Errorf("Solution02(%v, %v, %v) = %v, want = %v", c.inN, c.inA, c.inB, got, c.want)
		}
	}
}

func TestSolution03(t *testing.T){
	cases := []struct{
		inA string
		want int
	}{
		{"23",7},
		{"0081", 11},
		{"022", 9},
	}
	
	for _,c := range cases {
		got := Solution03(c.inA)
		if got!= c.want {
			t.Errorf("Solution03(%v) = %v, want = %v",c.inA, got, c.want)
		}
	}
}

func TestSolution04(t *testing.T){
	cases := []struct{
		inA []int
		want int
	}{
		{[]int{3,4,5,3,7},3},
		{[]int{1,2,3,4}, -1},
		{[]int{1,3,1,2}, 0},
	}
	
	for _,c := range cases {
		testName := fmt.Sprintf("Solution04(%v)", c.inA)
		t.Run(testName, func (t *testing.T){
			got := Solution04(c.inA)
			if got != c.want {
				t.Errorf("return %v, want = %v", got, c.want)
			}
		})
	}
}

func TestSolution06(t *testing.T){
	cases := []struct{
		inA [][]int
		want int
	}{
		{[][]int{
			[]int{4,3,4,5,3},
			[]int{2,7,3,8,4},
			[]int{1,7,6,5,2},
			[]int{8,4,9,5,5},
			},3},
		{[][]int{
			[]int{2,2,1,1},
			[]int{2,2,2,2},
			[]int{1,2,2,2},
			},2},
		{[][]int{
			[]int{7,2,4},
			[]int{2,7,6},
			[]int{9,5,1},
			[]int{4,3,8},
			[]int{3,5,4},
			},3},
		{[][]int{
			[]int{2,7,6},
			[]int{9,5,1},
			[]int{4,3,8},
			},3},
	}
	
	for _,c := range cases {
		testName := fmt.Sprintf("Solution06(%v)", c.inA)
		t.Run(testName, func (t *testing.T){
			got := Solution06(c.inA)
			if got != c.want {
				t.Errorf("return %v, want = %v", got, c.want)
			}
		})
	}
}

func TestSolution07(t *testing.T){
	cases := []struct{
		inA int
		inB int
		inX []int
		inY []int
		want int
	}{
		{2,4,
		[]int{4,0,1,-2},
		[]int{-4,4,3,0},
		1},
		{1,3,
		[]int{0,1,2,-2,3},
		[]int{0,1,4,1,0},
		2},
	}
	
	for _,c := range cases {
		testName := fmt.Sprintf("Solution07(%v, %v, %v, %v)", c.inA, c.inB, c.inX, c.inY)
		t.Run(testName, func (t *testing.T){
			got := Solution07(c.inA, c.inB, c.inX, c.inY)
			if got != c.want {
				t.Errorf("return %v, want = %v", got, c.want)
			}
		})
	}
}

func TestSolution08(t *testing.T){
	cases := []struct{
		inA int
		inB int
		inC int
		inD int
		want int
	}{
		{1,8,3,2,6},
		{2,3,3,2,3},
		{6,2,4,7,0},
	}
	
	for _,c := range cases {
		testName := fmt.Sprintf("Solution08(%v, %v, %v, %v)", c.inA, c.inB, c.inC, c.inD)
		t.Run(testName, func (t *testing.T){
			got := Solution08(c.inA, c.inB, c.inC, c.inD)
			if got != c.want {
				t.Errorf("return %v, want = %v", got, c.want)
			}
		})
	}
}