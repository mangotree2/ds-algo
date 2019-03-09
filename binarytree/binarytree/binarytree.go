package binarytree

import (
	"fmt"
	"github.com/mangotree2/ds-algo/binarytree"
	"github.com/mangotree2/ds-algo/stack/arraystack"
)

type BinaryTree struct {
	root *tree.Node
}

//todo insert delete find order 前中后+层 求高度

func NewBinaryTree(data interface{}) *BinaryTree {
	return &BinaryTree{root: tree.NewNode(data)}
}

func (bt *BinaryTree)Root()*tree.Node{
	return bt.root
}

func (bt *BinaryTree) SetRoot(r *tree.Node) {
	bt.root = r
}

func InOrderTraverse(n *tree.Node) {
	if n == nil {
		return
	}

	InOrderTraverse(n.Left)
	fmt.Printf("%+v",n.Value)
	InOrderTraverse(n.Right)

}

func (bt *BinaryTree) InOrderTraverseByStack() {
	p := bt.root
	s := arraystack.NewArrayStack()

	for !s.Empty() || nil != p {
		if nil != p {
			s.Push(p)
			p = p.Left
		} else {
			tmp := s.Pop().(*tree.Node)
			fmt.Printf("%+v",tmp.Value)
			p = tmp.Right
		}
	}

	fmt.Println()
}

func PreOrderTraverse(n *tree.Node) {
	if n == nil {
		return
	}

	fmt.Printf("%+v",n.Value)
	PreOrderTraverse(n.Left)
	PreOrderTraverse(n.Right)
}

func (bt *BinaryTree) PreOrderTraverseByStack() {
	p := bt.root
	s := arraystack.NewArrayStack()

	for !s.Empty() || nil != p {
		if nil != p {
			fmt.Printf("%+v",p.Value)
			s.Push(p)
			p = p.Left
		} else {
			p = s.Pop().(*tree.Node).Right
		}
	}

	fmt.Println()

}

func PostOrderTraverse(n *tree.Node) {
	if n == nil {
		return
	}

	PreOrderTraverse(n.Left)
	PreOrderTraverse(n.Right)
	fmt.Printf("%+v",n.Value)

}

func (bt *BinaryTree) PostOrderTraverseByStack() {
	s1 := arraystack.NewArrayStack()
	s2 := arraystack.NewArrayStack()
	s1.Push(bt.root)
	for !s1.Empty() {
		p := s1.Pop().(*tree.Node)
		s2.Push(p)
		if nil != p.Left {
			s1.Push(p.Left)
		}
		if nil != p.Right {
			s1.Push(p.Right)
		}
	}

	for !s2.Empty() {
		fmt.Printf("%+v", s2.Pop().(*tree.Node).Value	)
	}
}

func (bt *BinaryTree) LevelTraverse() {
	if bt.root == nil {
		return
	}
	queue := []*tree.Node{}
	queue = append(queue,bt.root)
	last := bt.root //上一层最右节点
	nlast := bt.root//当前层最右节点

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		fmt.Printf("%+v",q.Value)

		if q.Left != nil {
			queue = append(queue,q.Left)
			nlast = q.Left
		}

		if q.Right != nil {
			queue = append(queue,q.Right)
			nlast = q.Right
		}


		if last == q {
			last = nlast
			fmt.Println()
		}

	}

}

