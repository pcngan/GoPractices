package exercise

import (
	//"fmt"
	"math"
)

func Solution01(A []int) int{
	N:= len(A)
	var steps float64 = 0.0
	for i:=1; i<5; i++ {
		var temp float64 = 0.0
		for j:=0; j<N; j++ {
			temp += math.Abs(float64(A[j]-i))
		}
		if i==1 || temp < steps {
			steps = temp
		}
	}
	
	return int(steps)
}