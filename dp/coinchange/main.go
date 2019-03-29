package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
//示例 1:
//
//输入: coins = [1, 2, 5], amount = 11
//输出: 3
//解释: 11 = 5 + 5 + 1
//示例 2:
//
//输入: coins = [2], amount = 3
//输出: -1



func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}

	min := math.MaxInt64
	record := make(map[int]int)
	sort.IntsAreSorted(coins)
	backtracking(coins,amount,0,0,&min,record)
	if min == math.MaxInt64{
		return -1
	}

	return min
}

func backtracking(coins []int, amount, cnt,ca int,min *int,r map[int]int)  {


	if ca == amount {
		if cnt < *min && cnt != 0 {
			*min = cnt
		}

		return
	}

	if v,ok := r[ca];ok && cnt > v {
		//fmt.Println(ca, " ",v)
		return
	}
	r[ca] = cnt

	for _,v := range coins {
		if ca +v  <= amount {
			backtracking(coins,amount,cnt +1,ca +v,min,r)
		}

	}


}



func coinChangeDP(coins []int, amount int) int {
	dp := make([]int, amount+1)
	sort.Ints(coins)
	m := make(map[int]int)
	//dp存得对应金额最少coins数
	for i := 1; i < len(dp); i++ {
		dp[i] = 2 << 31
		for j := 0; j < len(coins); j++ {
			if i - coins[j] < 0 {
				break
			}

			if dp[i] < dp[i - coins[j]] + 1 {
				dp[i] = dp[i]
				//m[i] = coins[j]

			} else {
				dp[i] = dp[i - coins[j]] + 1
				m[i] = coins[j]

			}

		}
	}
	if dp[amount] ==  2 << 31 {
		dp[amount] = -1
	}

	if dp[amount] != -1 {
		sum := amount
		for i := dp[amount];i>0 ;i-- {
			fmt.Print(m[sum],"  ")
			sum = sum-m[sum]
		}
	}

	return dp[amount]
}



func main() {

	//fmt.Println(coinChange([]int{1,2,5},100))
	//fmt.Println(coinChange([]int{186,419,83,408},6249))
	fmt.Println(coinChangeDP([]int{186,419,83,408},6249))
	now := time.Now()
	fmt.Println(coinChange([]int{186,419,83,408},6249))
	fmt.Println(time.Since(now).Seconds())
}
