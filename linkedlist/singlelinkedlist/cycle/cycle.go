package main

import (
	"container/list"
	"fmt"
	"github.com/mangotree2/ds-algo/linkedlist/singlelinkedlist"
)

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