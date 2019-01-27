package stack

type Stack interface {
	Push (v interface{})
	Pop() interface{}
	Empty() bool
	Top() interface{}
	Flush()
}

