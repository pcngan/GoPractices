package exercise
/*
import (
	"fmt"
	"math"
)
*/
func Solution02(N int, A, B []int) bool{
	// init a matrix [N+1][N+1] to store edges
	edges := make([][]bool, N+1)
	for i:=0; i<=N; i++ {
		edges[i] = make([]bool, N+1)
	}
	
	M := len(A) //number of edges
	for i:=0; i<M; i++ {
		edges[A[i]][B[i]] = true
		edges[B[i]][A[i]] = true
	}
	
	//check path 0 ->N
	for i:=1; i< N; i++ {
		if !edges[i][i+1] {
			return false
		}
	}
	return true
}