package exercise

import (
	"fmt"
)

func getDigits(s string) []int {
	digits := make([]int, len(s))
	for i, v := range s {
		digits[i] = int(v - '0')
	}
	return digits
}

func Solution03(s string) int {

	DigitsMod3 := [][]int{
		{0, 3, 6, 9}, //%3=0
		{1, 4, 7},    //%3=1
		{2, 5, 8},    //%3=2
	}
	//
	digits := getDigits(s)
	fmt.Println(digits)
	// sum of all digits
	sum := 0
	for _, v := range digits {
		sum += v
	}
	count := 0
	if (sum%3 == 0){
		count = 1
	}
	for i := 0; i < len(digits); i++ {
		mod := (sum - digits[i]) % 3
		if (sum%3 == 0){
			count += len(DigitsMod3[mod]) -1
		} else {
			count += len(DigitsMod3[mod])
		}
	}
	return count
}