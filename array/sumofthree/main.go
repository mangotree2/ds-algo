package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/3sum/
// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]

// 非最优， 最优是排序后，左右指针查询
//这里其实也是求2数之和，两数之和就是一一对应的Map 映射


func threeSum(nums []int) [][]int {
	l := len(nums)
	m := make(map[int]int,l)
	mm := make(map[[3]int]struct{})
	ret := [][]int{}

	var at [3]int
	for i:=0; i < l; i++ {
		target := -nums[i]
		for j := i+1; j < l; j++{
			temp := target-nums[j]
			if v ,ok := m[temp]; ok && v != j && i != v{

				st := []int{nums[i],nums[v],nums[j]}
				sort.Ints(st)
				at = [3]int{st[0],st[1],st[2]}

				if _,ok := mm[at];!ok {
					mm[at] = struct{}{}
					ret = append(ret,st)
				}
			}
		}
		m[nums[i]] = i
	}

	return ret

}

func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
}
