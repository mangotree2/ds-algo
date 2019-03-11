package _0_1bag

//01 背包最大可承重下装可以装最多的价值

//weight 每个物品的重量
//value 每个物品的价值
//n多少个物品，w 背包可承重
func Knapsack(weight []int, value []int, n int, w int) int {
	status := make([][]int,n)
	for i :=range status {
		status[i] = make([]int,w+1)
	}
	for i := 0;i < n ;i++{
		for j:=0 ; j <= w ; j++ {
			status[i][j] = -1
		}
	}


	status[0][0] = 0
	status[0][weight[0]] = value[0]

	for i := 1;i < n; i++ {
		for j := 0 ; j <= w - weight[i];j ++ {
			if status[i-1][j] >= 0 {

				if status[i][j+weight[i]] <  status[i-1][j]+value[i]  {
					status[i][j+weight[i]] = status[i-1][j]+value[i]
				}
			}
		}

	}

	max := -1

	for j:=0 ;j <= w ;j ++ {
		if status[n-1][j] > max {max = status[n-1][j]}
	}

	return max

}