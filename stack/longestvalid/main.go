package main

import "fmt"

//给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
//
//示例 1:
//
//输入: "(()"
//输出: 2
//解释: 最长有效括号子串为 "()"
//示例 2:
//
//输入: ")()())"
//输出: 4
//解释: 最长有效括号子串为 "()()"

func longestValidParentheses(s string) int {

	stack := make([]int,0,len(s))
	index :=  make([]int,len(s))
	var max ,temp int

	for i ,v := range s {
		if v == '(' {
			stack = append(stack,i)
		} else if len(stack) > 0 {
			index[stack[len(stack)-1]],index[i] = 1,1
			stack = stack[:len(stack)-1]
		}
	}


	for _,v := range index {
		if v > 0{
			temp += 1
		} else {
			if max < temp {
				max = temp
			}
			temp = 0
		}
	}


	if max < temp {
		max = temp
	}
	return max


}


func main() {
	fmt.Println(longestValidParentheses(")()())"))
}
