package exercise
/*
import(
	"fmt"
)
*/
func checkMagic(A [][]int, size,rowIndex,ColIndex int) bool {
	//fmt.Println(size, rowIndex, ColIndex)
	Sum, temp := 0, 0

	for i := 0; i < size; i++ {
		Sum += A[rowIndex+i][ColIndex+i]
	}
	//fmt.Println("Sum: ", Sum)
	for i := 0; i < size; i++ {
		temp += A[rowIndex+i][ColIndex+size-i-1]
	}
	//fmt.Println("temp: ",temp)
	if temp != Sum {
		return false
	}

	for i := rowIndex; i < rowIndex+size; i++ {
		temp = 0
		for j := ColIndex; j < ColIndex+size; j++ {
			temp += A[i][j]
		}
		//fmt.Println("rowIndex ",i," temp ", temp)
		if temp != Sum {
			return false
		}
	}

	for i := ColIndex; i < ColIndex+size; i++ {
		temp = 0
		for j := rowIndex; j < rowIndex+size; j++ {
			temp += A[j][i]
		}
		//fmt.Println("ColIndex ",i," temp ", temp)
		if temp != Sum {
			return false
		}
	}
	return true
}

func Solution06(A [][]int) int {
	//N, M numbers of rows, columns
	N, M := len(A), len(A[0])
	Size := N
	if M < Size {
		Size = M
	}
	//fmt.Println(N,M,Size)
	i:=Size
	
	for ; i > 1; i-- {
		//fmt.Println(Size, i)
		for j := 0; j <= N-i; j++ {
			for k := 0; k <= M-i; k++ {
				if checkMagic(A, i, j, k) {
					return i
				}
			}
		}
	}

	return 1
}