package main

import "fmt"

//班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。
//
//给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。
//
//示例 1:
//
//输入:
//[[1,1,0],
// [1,1,0],
// [0,0,1]]
//输出: 2
//说明：已知学生0和学生1互为朋友，他们在一个朋友圈。
//第2个学生自己在一个朋友圈。所以返回2。
//示例 2:
//
//输入:
//[[1,1,0],
// [1,1,1],
// [0,1,1]]
//输出: 1
//说明：已知学生0和学生1互为朋友，学生1和学生2互为朋友，所以学生0和学生2也是朋友，所以他们三个在一个朋友圈，返回1。
//注意：
//
//N 在[1,200]的范围内。
//对于所有学生，有M[i][i] = 1。
//如果有M[i][j] = 1，则有M[j][i] = 1。

//A->B
//B->C


func findCircleNum(M [][]int) int {

	row := len(M)
	col := 0
	if row > 0 {
		col = len(M[0])
	} else  {
		col = 0
	}

	if row == 0 || col == 0 {
		return 0
	}

	visit := make([][]bool,row)

	for i := range visit{
		visit[i] = make([]bool,col)
		//visit[i][i] = true
	}


	cnt := 0
	for i := 0;i<row;i++{
		for j := 0; j < col;j++ {
			if visit[i][j] == false && M[i][j] == 1 {
				findByDFS(M,row,col,i,j,visit)

				cnt++
			}
		}

	}

	return cnt

}

func findByDFS(M [][]int, row, col, idx, idy int,visit [][]bool) {
	if idx < 0 || idx > row || idy < 0 || idy > col {

		return
	}


	if M[idx][idy] == 1 {

		visit[idx][idy] = true
		if idx == idy {
			for i :=idy+1; i < len(M[idx]);i++{
				if visit[idx][i] == false && M[idx][i] == 1{
					findByDFS(M,row,col,idx,i,visit)
				}
			}

			return
		}

		if idy <= row-1 {
			for i :=0; i < len(M[idy]);i++{
				if visit[idy][i] == false && M[idy][i] == 1{
					findByDFS(M,row,col,idy,i,visit)
				}
			}
		}


	}


}

func main() {
	//[[1,0,0],
	// [0,1,0],
	// [0,0,1]]
	fmt.Println(findCircleNum([][]int{{1,0,0},{0,1,0},{0,0,1}}))
	//[[1,1,0],
	// [1,1,0],
	// [0,0,1]]
	fmt.Println(findCircleNum([][]int{{1,1,0},{1,1,0},{0,0,1}}))
	//[[1,0,0,1],
	// [0,1,1,0],
	// [0,1,1,1],
	// [1,0,1,1]]
	fmt.Println(findCircleNum([][]int{{1,0,0,1},{0,1,1,0},{0,1,1,1},{1,0,1,1}}))



}
