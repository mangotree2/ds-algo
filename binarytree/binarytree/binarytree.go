package binarytree

import "github.com/mangotree2/ds-algo/binarytree"

type BinaryTree struct {
	root *tree.Node
}

//todo insert delete find order 前中后+层 求高度

func NewBinaryTree(data interface{}) *BinaryTree {
	return &BinaryTree{root: tree.NewNode(data)}
}

func (t *BinaryTree) Root() *tree.Node {
	return t.root
}

func (t *BinaryTree) SetRoot(node *tree.Node) {
	t.root= node
}

func (t *BinaryTree)Insert(data interface{}) {
	if t.root.Value == nil {
		t.root.Value = data
	}

	cur := t.root
	for nil != cur {
		cur = cur.Left
	}


}