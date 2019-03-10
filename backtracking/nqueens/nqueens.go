package nqueens

import "fmt"

var result = make([]int,8)

//8皇后问题
func cal8Queens(row int) {
	if row == 8 {
		printQueens(result,8)
		return
	}

	for col :=0; col <8 ;col++{
		if isOK(row,col) {
			result[row] = col
			cal8Queens(row+1)
		}
	}


}

func isOK(row, col int) bool {
	leftUp := col -1
	rightUp := col +1
	for i:= row-1;i>=0;i-- {
		if result[i] == col {
			return false
		}
		if leftUp >0 && result[i] == leftUp {
			return false
		}

		if rightUp < 8 && result[i] == rightUp {
			return false
		}
		leftUp--
		rightUp++
	}

	return true


}

func printQueens(result []int,N int) {
	for i := 0; i < N ; i ++ {
		for j :=0 ; j <N; j++ {
			if result[i] == j {
				fmt.Printf("Q ")
			} else {
				fmt.Printf("*  ")
			}
		}
	}

}
