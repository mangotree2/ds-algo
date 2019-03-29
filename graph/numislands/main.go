package main

import (
	"github.com/mangotree2/ds-algo/queue/circulardqueue"
)

//给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
//
//示例 1:
//
//输入:
//11110
//11010
//11000
//00000
//
//输出: 1
//示例 2:
//
//输入:
//11000
//11000
//00100
//00011
//
//输出: 3


//核心就是把1->0
var m = [][]byte{
	{'1', '1', '0', '1', '0'},
	{'1', '1', '0', '1', '0'},
	{'1', '1', '0', '1', '1'},
	{'0', '0', '1', '1', '0'},
}

func numIslandsByDFS(grid [][]byte) int {

	height:=len(grid)
	if height==0{
		return 0
	}
	width:=len(grid[0])
	visit:=make([][]bool,height)
	for i:=0;i<height;i++{
		visit[i]=make([]bool,width)
	}

	num:=0


	for i:=0;i<height;i++{
		for j:=0;j<width;j++{

			if grid[i][j]=='1'&&visit[i][j]==false{
				search(i,j,visit,grid,height,width)

				num+=1

			}
		}

	}


	return num

}
func search(idxi,idxj int,visit[][]bool,grid [][]byte,height,width int){
	if idxi<0||idxj<0||idxi==height||idxj==width{
		return
	}
	if grid[idxi][idxj]=='0'{
		return
	}
	if visit[idxi][idxj]==false{
		visit[idxi][idxj]=true
		search(idxi+1,idxj,visit,grid,height,width)

		search(idxi-1,idxj,visit,grid,height,width)

		search(idxi,idxj+1,visit,grid,height,width)

		search(idxi,idxj-1,visit,grid,height,width)


	}


}

/////////////////////////////////////////////


func numIslandsByBFS(grid [][]byte) int {
	var x, y, xx, yy, count, rows, cols int
	rows = len(grid)
	if rows > 0 {
		cols = len(grid[0])
	} else {
		cols = 0
	}

	if rows == 0 || cols == 0 {
		return 0
	}

	var dx, dy = []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}

	que := circulardqueue.Constructor(2*(rows+cols))

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				que.InsertLast(i)
				que.InsertLast(j)
				grid[i][j] = '0'
				for !que.IsEmpty() {
					x = que.GetFront()//因为储存的是坐标，所以是int，这里要强制转化，因为que.Front()返回的是interface{}类型
					que.DeleteFront()
					y = que.GetFront()
					que.DeleteFront()
					for k := 0; k < 4; k++ {
						xx = x + dx[k]
						yy = y + dy[k]
						if xx < 0 || xx >= rows || yy < 0 || yy >= cols {
							continue
						}
						if grid[xx][yy] == '1' {
							grid[xx][yy] = '0'
							que.InsertLast(xx)
							que.InsertLast(yy)
						}
					}
				}
				count++
			}
		}
	}
	return count
}