package main

import (
	"container/list"
	"fmt"
	"github.com/mangotree2/ds-algo/linkedlist/singlelinkedlist"
)

//给定一个链表，判断链表中是否有环。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
//
func main() {

	l1 := list.New()

	l1.PushBack(1)
	l1.PushBack(1)

	fmt.Println(HasCycle1(l1))

	l := singlelinkedlist.New()
	l.PushBack(1)
	l.PushBack(1)

	fmt.Println(HasCycle(l))


}

func HasCycle(l *singlelinkedlist.List) bool {

	fast := l.Front().Next()
	slow := l.Front()

	for nil != fast {
		if slow == fast {
			return true
		}

		fast = fast.Next().Next()
		slow = slow.Next()
	}


	return false
}

func HasCycle1(l *list.List) bool {

	fast := l.Front().Next()
	slow := l.Front()

	for nil != fast {
		if slow == fast {
			return true
		}

		fast = fast.Next().Next()
		slow = slow.Next()
	}


	return false
}