package main

import (
	"fmt"
	"math"
)

//给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
//
//示例 1:
//
//输入: [1,2,0]
//输出: 3
//示例 2:
//
//输入: [3,4,-1,1]
//输出: 2
//示例 3:
//
//输入: [7,8,9,11,12]
//输出: 1
//说明:
//
//你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。

func firstMissingPositive(nums []int) int {

	min := math.MaxInt64
	nums = append(nums,-1)
	l := len(nums)
	for i := 0;i < l;i++ {
		if nums[i] > 0 {
			if nums[i] < min {
				min = nums[i]
			}
		}
	}


	if min >0 && min != 1 {
		return 1
	}


	for i := 0; i < l; i++ {
		for nums[i] > 0 && nums[i] <= l-1 && nums[nums[i]] != nums[i] {
			nums[nums[i]],nums[i] = nums[i],nums[nums[i]]
		}
	}
	fmt.Println(nums)
	for i := 1; i < l ; i++ {
		if nums[i] != i {
			return i
		}
	}


	return l
}

func main() {
	fmt.Println(firstMissingPositive([]int{1}))
}
