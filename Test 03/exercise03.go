package tomato
/*
import (
)
*/

// Exercise #01
func Solution01(N int, K int) int {
	if K > (N * (N + 1) / 2) {
		return -1
	}
	count := 0
	for K > 0 {
		if K > N {
			K -= N
			N--
		} else {
			K = 0
		}
		count++
	}

	return count
}

//Exercise #02
func Solution02(S string) int {
	N := len(S)

	mDups := make(map[string]bool)
	mUniques := make(map[string]bool)

	for i := 1; i < N; i++ {
		for j := 0; j <= N-i; j++ {
			sub := S[j : j+i]
			if _, ok := mDups[sub]; ok {
				continue
			}

			if _, ok := mUniques[sub]; ok {
				mDups[sub] = true
				delete(mUniques, sub)
			} else {
				mUniques[sub] = true
			}
		}

		if len(mUniques) > 0 {
			return i
		}
	}

	return N
}