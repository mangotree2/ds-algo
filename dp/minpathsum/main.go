package main

import "math"

//给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//说明：每次只能向下或者向右移动一步。
//
//示例:
//
//输入:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//输出: 7
//解释: 因为路径 1→3→1→1→1 的总和最小。

func minPathSum(grid [][]int) int {

	status := make([][]int,len(grid))
	for i,v := range grid {
		status[i] = make([]int,len(v))
		for j :=0;j < len(v);j++ {
			status[i][j]= math.MaxInt64
		}
	}

	status[0][0] = grid[0][0]
	for i := 1; i < len(status[0]);i++ {
		status[0][i] = status[0][i-1]+grid[0][i]
	}

	for i := 1; i < len(status);i++ {
		status[i][0] = status[i-1][0]+grid[i][0]
	}

	for i:=1;i<len(status);i++{
		for j := 1;j<len(status[i]);j ++ {
			status[i][j] = min(status[i][j-1],status[i-1][j])+grid[i][j]

		}
	}


	return status[len(grid)-1][len(grid[0])-1]

}

func min(s1, s2 int) int {
	if s1 < s2 {
		return s1
	}else {
		return s2
	}
}

func main() {
	
}
