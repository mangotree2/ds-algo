package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	row := [9][10]bool{}
	col := [9][10]bool{}
	box := [9][10]bool{}

	for i :=0; i < len(board);i++ {
		for j:=0;j<len(board[i]); j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				//fmt.Println(num)
				//fmt.Println(row[i][num])
				//fmt.Println(col[j][num])
				//fmt.Println(box[int((i/3)/3 + j/3)][num])


				if row[i][num] || col[j][num] || box[int((i/3)*3 + j/3)][num] {
					return false
				} else {
					row[i][num] = true
					col[j][num] = true
					box[int((i/3)*3 + j/3)][num] = true
				}
			}
		}
	}

	return true



}
func main() {

	board :=[][]byte{
		{'.','.','.','.','.','.','.','.','2'},
		{'.','.','.','.','.','.','6','.','.'},
		{'.','.','1','4','.','.','8','.','.'},
		{'.','.','.','.','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','.','.'},
		{'.','.','.','.','3','.','.','.','.'},
		{'5','.','8','6','.','.','.','.','.'},
		{'.','9','.','.','.','.','4','.','.'},
		{'.','.','.','.','5','.','.','.','.'},
	}
	
	fmt.Println(isValidSudoku(board))
}


