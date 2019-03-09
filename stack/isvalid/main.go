package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//
//示例 1:
//
//输入: "()"
//输出: true
//示例 2:
//
//输入: "()[]{}"
//输出: true
//示例 3:
//
//输入: "(]"
//输出: false
//示例 4:
//
//输入: "([)]"
//输出: false
//示例 5:
//
//输入: "{[]}"
//输出: true

func isValid(s string) bool {

	stack := []int32{}

	sz := rune('(')
	sy := rune(')')

	zz := rune('[')
	zy := rune(']')

	dz := rune('{')
	dy := rune('}')

	for _,v := range s {
		switch v {
		case sy:
			if len(stack) <= 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pop != sz {
				return false
			}
		case zy:
			if len(stack) <= 0 {
				return false

			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pop != zz {
				return false
			}
		case dy:
			if len(stack) <= 0 {
				return false

			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pop != dz {
				return false
			}
		case sz,zz,dz:
			stack = append(stack,v)

		default:
			continue

		}

	}

	if len(stack)>0 {
		return false
	}

	return true


}

func main() {

	fmt.Println(isValid("(]"))
}
