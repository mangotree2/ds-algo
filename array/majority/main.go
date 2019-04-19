package main

import "fmt"

//给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在众数。
//
//示例 1:
//
//输入: [3,2,3]
//输出: 3
//示例 2:
//
//输入: [2,2,1,1,1,2,2]
//输出: 2
//m没考虑出现相同次数的











func majorityElement(nums []int) int {
	m := make(map[int]int)
	s := make([]int,0,len(nums)+1)

	for _,v := range nums {
		m[v]++
		if len(s) < m[v] {
			s = s[:m[v]]
			s[m[v]-1]=v
		}
		fmt.Println("v: ",v ,"cnt: ",m[v],"len s :",len(s))
		fmt.Println(s)


	}


	return s[len(s)-1]

}

func main() {
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
}
