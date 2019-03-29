package main

import "fmt"

//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
//例如，给定三角形：
//
//[
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
//]
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//
//说明：
//
//如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。



///**
//     * 解法1 二维数组进行求解
//     * @param triangle
//     * @return
//     */
//    public int minimumTotal(List<List<Integer>> triangle) {
//        if (triangle == null || triangle.size() == 0){
//            return 0;
//        }
//        // 加1可以不用初始化最后一层
//        int[][] dp = new int[triangle.size()+1][triangle.size()+1];
//
//        for (int i = triangle.size()-1; i>=0; i--){
//            List<Integer> curTr = triangle.get(i);
//            for(int j = 0 ; j< curTr.size(); j++){
//                dp[i][j] = Math.min(dp[i+1][j], dp[i+1][j+1]) + curTr.get(j);
//            }
//        }
//        return dp[0][0];
//    }
//    /**
//     * 解法2 一维数组进行求解
//     * @param triangle
//     * @return
//     */
//    public int minimumTotal(List<List<Integer>> triangle) {
//        if (triangle == null || triangle.size() == 0){
//            return 0;
//        }
//        // 只需要记录每一层的最小值即可
//        int[] dp = new int[triangle.size()+1];
//
//        for (int i = triangle.size() - 1; i >= 0; i--) {
//            List<Integer> curTr = triangle.get(i);
//            for (int j = 0; j < curTr.size(); j++) {
//                //这里的dp[j] 使用的时候默认是上一层的，赋值之后变成当前层
//                dp[j] = Math.min(dp[j],dp[j+1]) + curTr.get(j);
//            }
//        }
//        return dp[0];
//    }

//一维解决
func minimumTotal(triangle [][]int) int {

	if len(triangle) == 0 {
		return 0
	}


	status := make([]int,len(triangle)+1)

	for i := len(triangle)-1; i >= 0 ; i --{
		for j := 0;j < len(triangle[i]); j++ {
			//这里的status[j] 使用的时候默认是上一层的，赋值之后变成当前层
			status[j] = min(status[j],status[j+1]) + triangle[i][j]
		}

	}

	return status[0]

}

//二维数组
func minimumTotal2(triangle [][]int) int {

	if len(triangle) == 0 {
		return 0
	}

	status := make([][]int,len(triangle)+1)
	for i,v := range triangle {
		status[i] = make([]int,len(v))
	}
	status[len(triangle)] = make([]int,len(triangle[len(triangle)-1])+1)

	for i := len(triangle)-1; i >= 0 ; i --{
		for j := 0;j < len(triangle[i]); j++ {
			status[i][j] = min(status[i+1][j],status[i+1][j+1]) + triangle[i][j]
		}

	}

	return status[0][0]

}

func min(a,b int) int {
	if a <b {
		return a
	}

	return b
}


func main() {

	fmt.Println(minimumTotal([][]int{
		{2},{3,4},{6,5,7},{4,1,8,3},
	}))
}
