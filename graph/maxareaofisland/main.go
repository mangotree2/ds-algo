package main

import (
	"fmt"
	"math"
)

//给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
//
//找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)
//
//示例 1:
//
//[[0,0,1,0,0,0,0,1,0,0,0,0,0],
// [0,0,0,0,0,0,0,1,1,1,0,0,0],
// [0,1,1,0,1,0,0,0,0,0,0,0,0],
// [0,1,0,0,1,1,0,0,1,0,1,0,0],
// [0,1,0,0,1,1,0,0,1,1,1,0,0],
// [0,0,0,0,0,0,0,0,0,0,1,0,0],
// [0,0,0,0,0,0,0,1,1,1,0,0,0],
// [0,0,0,0,0,0,0,1,1,0,0,0,0]]
//对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。
//
//示例 2:
//
//[[0,0,0,0,0,0,0,0]]
//对于上面这个给定的矩阵, 返回 0。
//
//注意: 给定的矩阵grid 的长度和宽度都不超过 50。



func maxAreaOfIsland(grid [][]int) int {

	var row ,col int
	row = len(grid)

	if row > 0 {
		col = len(grid[0])
	} else {
		col = 0
	}

	if row == 0 || col == 0 {
		return 0
	}

	max := math.MinInt64

	visit := make([][]bool,row)

	for i:= 0; i < row ; i ++{
		visit[i] = make([]bool,len(grid[i]))
	}


	for i := 0; i < row; i++ {
		for j := 0; j < col;j++ {

			if grid[i][j] == 1 && visit[i][j] == false {
				cnt := 0
				searchByDFS(grid,visit,row,col,i,j,&cnt)

				if max < cnt {
					max = cnt
				}
			}

		}

	}

	if max == math.MinInt64 {
		return 0
	}
	return max

}

func searchByDFS(grid [][]int, visit [][]bool, row, col int, idx, idy int,cnt *int)  {

	if idx < 0 || idx >= row || idy < 0 || idy >= col {
		return
	}

	if grid[idx][idy] == 0 {
		return
	}


	if grid[idx][idy] == 1 &&visit[idx][idy] == false {
		grid[idx][idy] = 0
		visit[idx][idy] = true
		*cnt++

		searchByDFS(grid,visit,row,col,idx-1,idy,cnt)
		searchByDFS(grid,visit,row,col,idx+1,idy,cnt)
		searchByDFS(grid,visit,row,col,idx,idy-1,cnt)
		searchByDFS(grid,visit,row,col,idx,idy+1,cnt)

	}

	return



}




func main() {
	fmt.Println(maxAreaOfIsland([][]int{
		{1,1,0,0,0},{1,1,0,0,0},{0,0,0,1,1},{0,0,0,1,1},
	}))
}























