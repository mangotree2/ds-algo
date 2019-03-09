package main

import "fmt"


//合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//
//示例:
//
//输入:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//输出: 1->1->2->3->4->4->5->6
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergetwolist(l1, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}


	root := new(ListNode)
	start := root
	for {

		if l1.Val <= l2.Val {
			root.Next = l1
			l1 = l1.Next

		} else {
			root.Next = l2
			l2 = l2.Next
		}
		root = root.Next

		if l1 == nil || l2 == nil {
			break
		}

	}
	if  l1== nil {
		root.Next =l2
	}
	if l2 == nil {
		root.Next = l1
	}
	//for l := start; l != nil ; l = l.Next{
	//	fmt.Println(l.Val)
	//}
	return start.Next

}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil

	}

	if len(lists) == 1 {
		return lists[0]
	}

	if len(lists) == 2 {
		return mergetwolist(lists[0], lists[1])
	}

	lists = append(lists, mergetwolist(lists[0], lists[1]))
	lists = lists[2:]


	return mergeKLists(lists)

}

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l3 := &ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  6,
			Next:nil,
		},
	}

	for l := mergeKLists([]*ListNode{l1,l2,l3}); l != nil ;l = l.Next {
		fmt.Println(l.Val)
	}
}
