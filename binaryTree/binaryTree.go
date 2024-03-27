package Binarytree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	inorder(root, &res)
	return res
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

// 建立 binaryTree
func CreateBinaryTree(val []int) *TreeNode {
	if len(val) == 0 {
		return nil
	}

	return buildTree(val, 0)
}

func buildTree(val []int, index int) *TreeNode {
	if index >= len(val) || val[index] == math.MinInt32 {
		return nil
	}

	node := &TreeNode{Val: val[index]}

	node.Left = buildTree(val, 2*index+1)
	node.Right = buildTree(val, 2*index+2)

	return node
}
