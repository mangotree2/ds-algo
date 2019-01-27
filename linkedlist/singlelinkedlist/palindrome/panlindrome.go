package main

import (
	"fmt"
	"github.com/mangotree2/ds-algo/linkedlist/singlelinkedlist"
)

func isPalindrome1(l *singlelinkedlist.List) bool {
	llen := l.Len()
	fmt.Println("len" ,llen)
	if llen == 0 {
		return false
	}

	if llen == 1 {
		return true
	}


	s := make([]string,0,llen/2)
	cur := l.Front()
	for i := 1; i <= llen; i++ {
		if llen%2 != 0 && i == (llen/2 +1) {
			cur = cur.Next()

			continue
		}

		if i <= llen/2 {
			s = append(s,cur.Value.(string))
		} else {
			fmt.Println("s: " ,s, "cur :",cur.Value.(string))
			if s[llen-i] != cur.Value.(string) {
				return false
			}
		}

		cur = cur.Next()


	}

	return true
}



func main() {

	l := singlelinkedlist.New()

	l.PushBack("a")
	//l.PushBack("b")
	//l.PushBack("c")
	//l.PushBack("b")
	//l.PushBack("a")

	fmt.Println(isPalindrome1(l))
}