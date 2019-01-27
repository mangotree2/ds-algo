package main

import "github.com/mangotree2/ds-algo/linkedlist/singlelinkedlist"

func main() {

}

func FindMiddleNode(l *singlelinkedlist.List) *singlelinkedlist.Node {
	if l.Len() == 0 || l.Len() == 1 {
		return nil
	}
	//若无len 字段

	slow,fast := l.Front(),l.Front().Next()

	for nil != fast && nil != fast.Next() {
		slow = slow.Next()
		fast = fast.Next().Next()
	}

	return slow

}
