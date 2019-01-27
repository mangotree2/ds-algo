package main

import "github.com/mangotree2/ds-algo/linkedlist/singlelinkedlist"

func main() {


}

func DeleteBottomN(l *singlelinkedlist.List,n int) {
	if n < 0 {
		return
	}

	fast := l.Front()
	for i := 1; i < n && fast != nil ;i++ {
		fast = fast.Next()
	}

	if fast == nil {
		return
	}


	slow := l.Front()

	for nil != fast.Next() {
		slow = slow.Next()
		fast = fast.Next()
	}

	slow.Next() = slow.Next().Next()
}
