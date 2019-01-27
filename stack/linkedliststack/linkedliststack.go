package linkedliststack

type node struct {
	next *node
	value interface{}
}

type LinkedListStack struct {
	topNode *node
}

func (s *LinkedListStack) Push(v interface{}) {
	s.topNode = &node{next:s.topNode, value:v}
}

func (s *LinkedListStack) Pop() interface{} {
	if s.topNode == nil {
		return nil
	}
	v := s.topNode.value
	s.topNode = s.topNode.next
	return v
}

func (s *LinkedListStack) Empty() bool {
	if s.topNode == nil {
		return true
	}

	return false
}

func (s *LinkedListStack) Top() interface{} {
	if s.topNode != nil {
		return s.topNode.value
	}
	return nil
}

func (s *LinkedListStack) Flush() {
	s.topNode = nil
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}


