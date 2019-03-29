package main

import (
	"github.com/mangotree2/ds-algo/binarytree"
)

//给定一个二叉树，找出其最大深度。
//
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回它的最大深度 3 。









func maxDepth(root *tree.Node) int {
	if root == nil {
		return 0
	}
	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)
	if lDepth > rDepth {
		return lDepth +1
	} else {
		return rDepth +1
	}

}

func main() {

}
