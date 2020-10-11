package exercise

import (
	"fmt"
)

func checkOk(A []int, i, left, right int) bool {
	if i == 0 || i == len(A)-1 {
		return true
	}

	if (A[i]-A[left])*(A[right]-A[i]) >= 0 {
		return false
	}

	return true
}

func removeTree(A []int, i int) bool {
	if i == 0 || i == len(A)-1 {
		return true
	}

	return checkOk(A, i-1, i-2, i+1) && checkOk(A, i+1, i-1, i+2)

}

func Solution04(A []int) int {
	N := len(A)

	fails := make([]int, 0, N)

	for i := 1; i < N-1; i++ {
		left := A[i] - A[i-1]
		right := A[i+1] - A[i]
		if left*right >= 0 {
			fails = append(fails, i)
		}
	}
	fmt.Println(fails)
	count := 0
	switch {
	case len(fails) == 0:
		return 0
	case len(fails) > 1:
		count = -1
	default:
		needRemoves := []int{fails[0] - 1, fails[0], fails[0] + 1}
		for _, v := range needRemoves {
			if removeTree(A, v) {
				count++
			}
		}

		if count == 0 {
			count = -1
		}
	}

	return count

}