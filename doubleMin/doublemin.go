package doublemin

import (
	"fmt"
	"math"
	)

func Solution(N int) int{
	if N<=0{
		panic("the input is not a positive number")
	}
	s := fmt.Sprintf("%v", N)
	ls := len(s)
	//letters := strings.Split(s, "")
	a := make([]int, ls)
	sum := 0
	for i, v := range s {
		j := int(v)- '0'
		if j<0 || j>9 {
			panic("value is not  digits")
		}
		a[i] = j
		sum += j
	}
	sum2 := 2 * sum

	var b []int
	for i := ls - 1; i >= 0; i-- {
		dif := sum2 + a[i] - sum
		switch {
		case dif > 9:
			b = append(b, 9)
			sum2 -= 9
		case dif > 0:
			b = append(b, dif)
			sum2 -= dif
		default:
			b = append(b, a[i])
			sum2 -= a[i]
		}
		sum -= a[i]
	}
	for sum2 > 0 {
		if sum2 > 9 {
			b = append(b, 9)
			sum2 -= 9
		} else {
			b = append(b, sum2)
			break
		}
	}
	value := 0
	for i, v := range b {	
		value += v * int(math.Pow10(i))
	}
	return value
}

