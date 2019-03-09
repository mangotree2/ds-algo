package binarytree

import (
	"fmt"
	"github.com/mangotree2/ds-algo/binarytree"
	"testing"
)

func TestBinaryTree_InOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.Left = tree.NewNode(3)
	binaryTree.root.Right = tree.NewNode(4)
	binaryTree.root.Right.Left = tree.NewNode(5)

	binaryTree.InOrderTraverseByStack()
	InOrderTraverse(binaryTree.root)
	fmt.Println()
}

func TestBinaryTree_PreOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.Left = tree.NewNode(3)
	binaryTree.root.Right = tree.NewNode(4)
	binaryTree.root.Right.Left = tree.NewNode(5)

	binaryTree.PreOrderTraverseByStack()
	PreOrderTraverse(binaryTree.root)
	fmt.Println()

}

func TestBinaryTree_PostOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.Left = tree.NewNode(3)
	binaryTree.root.Right = tree.NewNode(4)
	binaryTree.root.Right.Left = tree.NewNode(5)

	binaryTree.PostOrderTraverseByStack()
	PostOrderTraverse(binaryTree.root)
	fmt.Println()

}

func TestBinaryTree_LevelTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.Left = tree.NewNode(3)
	binaryTree.root.Right = tree.NewNode(4)
	binaryTree.root.Right.Left = tree.NewNode(5)

	binaryTree.LevelTraverse()
	fmt.Println()

}
