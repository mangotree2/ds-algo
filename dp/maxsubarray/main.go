package main


import "fmt"

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//进阶:
//
//如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。




//我的思路 当前位置的最大和
// max[index]= max(a[index],a[index]+max[index-1])
func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	maxs := make([]int,len(nums))
	maxs[0] = nums[0]
	ret := maxs[0]
	for i := 1;i<len(nums);i ++ {
		maxs[i] = max(nums[i],nums[i]+maxs[i-1])
		if ret < maxs[i] {
			ret = maxs[i]

		}
	}

	return ret

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b

}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}
