package main

import "fmt"

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	head :=&ListNode{}
	l3 := head
	preAdd := 0

	for  {


		sum := l1.Val + l2.Val
		if preAdd > 0 {
			sum += preAdd

		}

		preAdd = sum /10
		sum = sum %10

		l3.Val = sum




		l1 = l1.Next
		l2 = l2.Next
		if l1 != nil && l2 != nil {
			next := &ListNode{}
			l3.Next = next
			l3 = next
		} else {
			break
		}

	}

	if l1 == nil && l2 == nil && preAdd >0 {
		next := &ListNode{}
		l3.Next = next
		l3 = next
		l3.Val = preAdd
		return head

	}


	if l1 != nil {
		next := &ListNode{}
		l3.Next = next
		l3 = next
		for p := l1;p != nil ;p = p.Next {

			if preAdd > 0 {
				p.Val += preAdd

			}

			preAdd = p.Val /10
			p.Val = p.Val %10

			l3.Val = p.Val
			if p.Next != nil {
				next := &ListNode{}
				l3.Next = next
				l3 = next
			}
		}
	}

	if l2 != nil {
		next := &ListNode{}
		l3.Next = next
		l3 = next
		for p := l2;p != nil ;p = p.Next {
			if preAdd > 0 {
				p.Val += preAdd
			}

			preAdd = p.Val /10
			p.Val = p.Val %10

			l3.Val = p.Val
			if p.Next != nil {
				next := &ListNode{}
				l3.Next = next
				l3 = next
			}
		}
	}

	if preAdd > 0 {
		next := &ListNode{}
		l3.Next = next
		l3 = next
		l3.Val = preAdd
	}
	//for p:= head;p!=nil;p= p.Next{
	//	fmt.Println(p.Val)
	//}

	return head



}

func main() {

	l1 := &ListNode{
		Val:2,
		Next:&ListNode{
			Val:4,
			Next:&ListNode{
				Val:3,
			},
		},
	}

	l2 := &ListNode{
		Val:5,
		Next:&ListNode{
			Val:6,
			Next:&ListNode{
				Val:4,
			},
		},
	}
	l3 := addTwoNumbers(l1,l2)

	for p:= l3;p!=nil;p= p.Next{
		fmt.Println(p.Val)
	}
}
