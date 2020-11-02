package binary_tree

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   interface{}
}

// 前序遍历
func PreorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	node.print()
	PreorderTraversal(node.left)
	PreorderTraversal(node.right)
}

// 中序遍历
func InorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	InorderTraversal(node.left)
	node.print()
	InorderTraversal(node.right)
}

// 后续遍历
func PostorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	PostorderTraversal(node.left)
	PostorderTraversal(node.right)
	node.print()
}

func (n TreeNode) print() {
	fmt.Print(" => ", n.val)
}
