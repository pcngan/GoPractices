package exercise

import (
	"fmt"
)

func timeValid(A []int) bool {

	if len(A) < 4 {
		return false
	}

	if (A[0]*10+A[1]) > 23 || (A[2]*10+A[3]) > 59 {
		return false
	}

	return true
}

func Solution08(A, B, C, D int) int {
	Digits := []int{A, B, C, D}
	for _,v:= range Digits{
		if v<0 || v>9 {
			return 0
		}
	}
	
	tempValue := make([]int, 4)
	timeMap := make(map[string]int)
	for i, v := range Digits {
		tempValue[0] = v
		for j, vv := range Digits {
			if j == i {
				continue
			}
			tempValue[1] = vv
			for k, vvv := range Digits {
				if k == i || k == j {
					continue
				}
				tempValue[2] = vvv
				for m, vvvv := range Digits {
					if m == i || m == j || m == k {
						continue
					}
					tempValue[3] = vvvv
					if timeValid(tempValue) {
						key := fmt.Sprintf("%v%v%v%v", tempValue[0], tempValue[1], tempValue[2], tempValue[3])
						timeMap[key] = 1
					}
				}
			}
		}
	}
	return len(timeMap)
}