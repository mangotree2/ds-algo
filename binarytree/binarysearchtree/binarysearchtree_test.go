package binarysearchtree


import "testing"

var compareFunc = func(v, nodeV interface{}) int {
	vInt := v.(int)
	nodeVInt := nodeV.(int)

	if vInt > nodeVInt {
		return 1
	} else if vInt < nodeVInt {
		return -1
	}
	return 0
}

func TestBST_Find(t *testing.T) {
	bst := NewBinarySearchTree(1, compareFunc)

	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	t.Log(bst.Find(2))
}

func TestBST_Insert(t *testing.T) {
	bst := NewBinarySearchTree(1, compareFunc)

	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	//bst.InOrderTraverse()
	SortNodesByInOrder(bst.Root())
}

func TestBST_Min(t *testing.T) {
	bst := NewBinarySearchTree(1, compareFunc)

	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	t.Log(bst.Min())
}

func TestBST_Max(t *testing.T) {
	bst := NewBinarySearchTree(1, compareFunc)

	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	t.Log(bst.Max())
}

func TestBST_DeleteA(t *testing.T) {
	bst := NewBinarySearchTree(1, compareFunc)

	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	t.Log(bst.Delete(7))

	//bst.InOrderTraverse()
	SortNodesByInOrder(bst.Root())

}

func TestBST_DeleteB(t *testing.T) {
	bst := NewBinarySearchTree(1, compareFunc)

	t.Log(bst.Delete(1))
	t.Log(bst.Root())

	//bst.InOrderTraverse()
	SortNodesByInOrder(bst.Root())

}

func TestBST_DeleteC(t *testing.T) {
	bst := NewBinarySearchTree(1, compareFunc)

	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	t.Log(bst.Delete(1))

	//bst.InOrderTraverse()
	SortNodesByInOrder(bst.Root())

}

func TestBST_DeleteD(t *testing.T) {
	bst := NewBinarySearchTree(1, compareFunc)

	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(5)

	t.Log(bst.Delete(1))

	//bst.InOrderTraverse()
	SortNodesByInOrder(bst.Root())

}
func TestBST_DeleteE(t *testing.T) {
	bst := NewBinarySearchTree(5, compareFunc)

	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(1)

	t.Log(bst.Delete(5))
	//bst.InOrderTraverse()
	SortNodesByInOrder(bst.Root())

}

func TestBST_DeleteF(t *testing.T) {
	bst := NewBinarySearchTree(5, compareFunc)

	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(1)

	t.Log(bst.Delete(2))
	//bst.InOrderTraverse()
	SortNodesByInOrder(bst.Root())

}

func TestBST_DeleteG(t *testing.T) {
	bst := NewBinarySearchTree(5, compareFunc)

	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(1)

	t.Log(bst.Delete(1))
	//bst.InOrderTraverse()
	SortNodesByInOrder(bst.Root())

}