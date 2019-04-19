package binarysearchtree

import (
	"fmt"
	"github.com/mangotree2/ds-algo/binarytree/binarytree"
	"github.com/mangotree2/ds-algo/binarytree"

)


type BinarySearchTree struct {
	*binarytree.BinaryTree
	
	compareFunc func(v,nodeV interface{}) int
}

//this not support repeat value ,todo support it
func NewBinarySearchTree(rootV interface{},compareFunc func(v,nodeV interface{}) int) *BinarySearchTree {
	if nil == compareFunc {
		return nil
	}

	return &BinarySearchTree{BinaryTree: binarytree.NewBinaryTree(rootV), compareFunc: compareFunc}
}

func (t *BinarySearchTree) Find(v interface{}) *tree.Node {
	cur := t.BinaryTree.Root()

	for nil != cur {
		compareRet := t.compareFunc(v,cur.Value)
		if compareRet == 0 {
			return cur
		} else if compareRet > 0 {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}

	return nil
}

func (t *BinarySearchTree) Insert(v interface{}) bool {

	cur := t.Root()

	for nil != cur {
		compareRet := t.compareFunc(v,cur.Value)
		if compareRet == 0 { //todo support
			return false
		} else if compareRet > 0 {
			if cur.Right == nil {
				cur.Right = tree.NewNode(v)
				break
			}
			cur = cur.Right
		} else{
			if cur.Left == nil {
				cur.Left = tree.NewNode(v)
				break
			}
			cur = cur.Left
		}

	}

	return true

}

func (t *BinarySearchTree) Delete(v interface{}) bool {

	var parent *tree.Node
	cur := t.Root()

	for nil != cur {
		compareRet := t.compareFunc(v,cur.Value)
		if compareRet > 0 {
			parent = cur
			cur = cur.Right
		} else if compareRet < 0 {
			parent = cur
			cur = cur.Left
		} else {
			break
		}

	}

	if cur == nil {
		return false
	} else if nil != cur.Right && nil != cur.Left{
		//查找右子树最小节点
		min := cur.Right
		minP := cur

		for nil != min.Left {
			minP = min
			min = min.Left
		}
		//右子树最小值替代当前值
		cur.Value = min.Value
		cur = min
		parent = minP

	}

	//利用有左节点此节点就不是最小的节点，出来结果有2种，
	// 1：在右子树找到有左节点则此时cur为叶子节点，
	// 2: 在右子树没找到左节点，则此时cur 为删除节点的右节点，此时parent 为要删除节点的指针

	//删除节点是叶子的节点或者仅有一个节点
	var child *tree.Node
	if cur.Left != nil {
		child = cur.Left
	} else if cur.Right != nil {
		child = cur.Right
	}


	if parent == nil {
		//t.Root() = child
		t.SetRoot(child)
	} else if parent.Left == cur {
		parent.Left = child
	} else {
		parent.Right = child
	}

	return true
}

func (t *BinarySearchTree) Min() *tree.Node {
	cur := t.Root()
	for nil != cur.Left {
		cur = cur.Left
	}
	return  cur
}

func (t *BinarySearchTree) Max() *tree.Node {
	cur := t.Root()
	for nil != cur.Right {
		cur = cur.Right
	}
	return cur
}

func SortNodesByInOrder(node *tree.Node) {
	if node == nil {
		return
	}
	SortNodesByInOrder(node.Left)
	fmt.Println(node.Value)
	SortNodesByInOrder(node.Right)
	return
}