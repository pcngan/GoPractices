package exercise
/*
import (
	"fmt"
)
*/
func Solution07(inner, outter int, pointsX, pointsY []int) int {
	if len(pointsX) != len(pointsY) {
		return 0
	}
	count := 0
	for i := 0; i < len(pointsX); i++ {

		if (pointsX[i]*pointsX[i]+pointsY[i]*pointsY[i]) > inner*inner &&
			(pointsX[i]*pointsX[i]+pointsY[i]*pointsY[i]) < outter*outter {
			count++
		}
	}

	return count
}