package main

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


func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	pNow := head
	var pPre *ListNode
	var pNext *ListNode
	var tail *ListNode


	for pNow != nil {

		pNext = pNow.Next

		if pNext == nil {
			tail = pNow
		}

		pNow.Next= pPre

		pPre = pNow

		pNow = pNext

	}

	return tail


}

func reverseListByRecursive(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	} else {

		tail := reverseListByRecursive(head.Next)

		head.Next.Next = head

		head.Next = nil
		return tail

	}

}

func main() {

}
