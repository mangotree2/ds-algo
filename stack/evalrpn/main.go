package main

import (
	"fmt"
	"strconv"
)

//根据逆波兰表示法，求表达式的值。
//
//有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
//
//说明：
//
//整数除法只保留整数部分。
//给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
//示例 1：
//
//输入: ["2", "1", "+", "3", "*"]
//输出: 9
//解释: ((2 + 1) * 3) = 9
//示例 2：
//
//输入: ["4", "13", "5", "/", "+"]
//输出: 6
//解释: (4 + (13 / 5)) = 6
//示例 3：
//
//输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
//输出: 22
//解释:
//  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
//= ((10 * (6 / (12 * -11))) + 17) + 5
//= ((10 * (6 / -132)) + 17) + 5
//= ((10 * 0) + 17) + 5
//= (0 + 17) + 5
//= 17 + 5
//= 22

func evalRPN(tokens []string) int {

	const subVal = 48

	stack := make([]int64,0,len(tokens))
	var bop,op,ret int64

	for _,v := range tokens {

		if len(v) > 1 {
			num,_ := strconv.ParseInt(v,10,64)
			stack = append(stack,num)
			continue
		}

		switch v[0] {
		default:
			stack = append(stack,int64(v[0]-subVal))
		case '+':
			bop = stack[len(stack) -1]
			op = stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			ret = op +bop
			stack = append(stack,ret)
		case '-':
			bop = stack[len(stack) -1]
			op = stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			ret = op -bop
			stack = append(stack,ret)
		case '*':
			bop = stack[len(stack) -1]
			op = stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			ret = op *bop
			stack = append(stack,ret)
		case '/':
			bop = stack[len(stack) -1]
			op = stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			ret = op /bop
			stack = append(stack,ret)
		}

	}

	return int(stack[0])
}
func main() {
	for _,v := range "=-*/" {
		fmt.Println(v)

	}

}
