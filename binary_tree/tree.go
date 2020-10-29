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

func (n TreeNode) print() {
	fmt.Print(" => ", n.val)
}
