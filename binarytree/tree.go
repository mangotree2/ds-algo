package tree

import "fmt"

type Node struct {
	Value interface{}
	Left *Node
	Right *Node
}

func NewNode(v interface{}) *Node {
	return &Node{Value:v}
}

func (n *Node)String() string {
	return fmt.Sprintf("(data:%v), (left:%+v), (rigth:%+v)",n.Value,n.Left,n.Right)
}
