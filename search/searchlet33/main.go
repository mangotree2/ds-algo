package main

import "fmt"

//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
//( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//
//搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//
//你可以假设数组中不存在重复的元素。
//
//你的算法时间复杂度必须是 O(log n) 级别。
//
//示例 1:
//
//输入: nums = [4,5,6,7,0,1,2], target = 0
//输出: 4
//示例 2:
//
//输入: nums = [4,5,6,7,0,1,2], target = 3
//输出: -1





func search(nums []int, target int) int {

	hi := len(nums) - 1
	low := 0


	mid := 0
	for {
		if low > hi {
			break
		}

		mid = (hi+low)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[hi] {
			if nums[mid] < target && nums[hi] >= target {
				low = mid +1
			} else {
				hi = mid-1
			}
		} else {
			if nums[mid] > target && nums[low] <= target {
				hi = mid -1
			} else {
				low = mid + 1
			}

		}



	}



	return -1

}

func main() {
	//fmt.Println(search([]int{4,5,6,7,0,1,2},3))
	fmt.Println(search([]int{3,5,1},3))

}


