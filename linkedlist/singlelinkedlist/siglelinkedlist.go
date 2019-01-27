package singlelinkedlist

type Node struct {
	next *Node

	list *List

	Value interface{}
}

func (n *Node) Next() *Node {
	if p := n.next; n.list != nil && p != &n.list.root {
		return p
	}
	return nil
}

//todo 2个返回值 或者直接panic
func (n *Node) Prev() *Node {
	if n.list == nil {
		return nil
	}

	pre := &n.list.root

	for pre.next != n {
		pre = pre.next
	}

	if pre.next == n && pre != &n.list.root {
		return pre
	}

	return nil

}

type List struct {
	root Node
	len  int
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.len = 0
	return l
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Node {
	if l.len == 0 {
		return nil
	}

	n := l.root.next
	for n.next != &l.root {
		n = n.next
	}

	return n
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) insert(e, at *Node) *Node {
	e.next = at.next
	at.next = e
	e.list = l
	l.len++
	return e
}

func (l *List) insertValue(v interface{}, at *Node) *Node {
	return l.insert(&Node{Value: v}, at)
}

func (l *List) remove(e *Node) *Node {
	pre := &l.root
	for pre.next != e {
		pre = pre.next
	}

	pre.next = e.next
	e.list = nil
	//e.next = nil
	l.len--
	return e
}

func (l *List) Remove(e *Node) interface{} {
	if e == nil || e.list == nil || e.next == nil {
		return false
	}

	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

func (l *List) PushFront(v interface{}) *Node {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

func (l *List) PushBack(v interface{}) *Node {
	l.lazyInit()
	n := l.Back()
	if n == nil {
		return l.insertValue(v, &l.root)
	}
	return l.insertValue(v, n)
}

func (l *List) InsertBefore(v interface{}, mark *Node) *Node {
	if mark == nil || mark.next == nil || mark.list == nil || mark.list != l {
		return nil
	}

	pre := &l.root
	for pre.next != mark {
		pre = pre.next
	}

	return l.insertValue(v, pre)

}

func (l *List) InsertAfter(v interface{}, mark *Node) *Node {
	if mark == nil || mark.next == nil || mark.list == nil || mark.list != l {
		return nil
	}

	return l.insertValue(v, mark)
}

func (l *List) MoveToBack(n *Node) {
	if n == nil || n.next == nil || n.list == nil || n.list != l {
		return
	}

	if n.next == &l.root {
		return
	}
	b := l.Back()
	l.insert(l.remove(n), b)
}

func (l *List) MoveToFront(n *Node) {
	if n == nil || n.next == nil || n.list == nil || n.list != l {
		return
	}

	if l.root.next == n {
		return
	}

	l.insert(l.remove(n), &l.root)
}

func (l *List) MoveBefore(e, mark *Node) {
	if mark == nil || mark.next == nil || mark.list == nil ||
		mark.list != l || e == nil || e.next == nil || e.list == nil || e.list != l {
		return
	}

	if e == mark {
		return
	}

	pre := &l.root
	for pre.next != mark {
		pre = pre.next
	}

	if e == pre {
		return
	}

	l.insert(l.remove(e), pre)
}

func (l *List) MoveAfter(e, mark *Node) {
	if mark == nil || mark.next == nil || mark.list == nil ||
		mark.list != l || e == nil || e.next == nil || e.list == nil || e.list != l {
		return
	}
	if e == mark {
		return
	}

	l.insert(l.remove(e), mark)
}

func (l *List) Reverse() {
	l.lazyInit()
	if l.root.next == &l.root || l.root.next.next == &l.root{
		return
	}

	pre := &l.root
	cur := l.root.next
	for cur != &l.root {
		tmp := cur.next
		cur.next = pre

		pre = cur
		cur = tmp
	}

	l.root.next = pre
}



//func (l *List) PushBackList(other *List) {
//	l.lazyInit()
//	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
//		l.insertValue(e.Value, l.root.prev)
//	}
//}
//
//// PushFrontList inserts a copy of an other list at the front of list l.
//// The lists l and other may be the same. They must not be nil.
//func (l *List) PushFrontList(other *List) {
//	l.lazyInit()
//	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
//		l.insertValue(e.Value, &l.root)
//	}
//}
