package arraystack

type ArrayStack struct {
	data []interface{}
	top int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data:[]interface{}{},
		top :-1,
	}
}

func (a *ArrayStack) Push(v interface{}) {
	if a.top < 0 {
		a.top = 0
	} else {
		a.top++
	}

	if a.top > len(a.data) {
		a.data = append(a.data,v)
	} else {
		a.data[a.top] = v
	}

}

func (a *ArrayStack) Pop() interface{} {
	if a.top < 0 {
		return nil
	}

	v := a.data[a.top]
	a.top--
	return v
}

func (a *ArrayStack) Empty() bool {
	if a.top < 0 {
		return true
	}
	return false
}

func (a *ArrayStack) Top() interface{} {
	if a.top < 0 {
		return nil
	}

	return a.data[a.top]
}

func (a *ArrayStack) Flush() {
	a.top = -1
}




