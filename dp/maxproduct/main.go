package main

//输入: [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//示例 2:
//
//输入: [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。


func maxProduct(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	maxEnd := nums[0]
	minEnd := nums[0]
	ret := nums[0]
	for i:=1;i<len(nums);i++{
		maxTemp := maxEnd
		minTemp := minEnd

		maxEnd = max(nums[i],max(nums[i]*maxTemp,nums[i]*minTemp))
		minEnd = min(nums[i],min(nums[i]*maxTemp,nums[i]*minTemp))

		ret = max(maxEnd,ret)
	}
	return ret
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {

}
