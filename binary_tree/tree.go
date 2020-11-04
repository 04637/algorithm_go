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

// 前序遍历 非递归
// todo 与网上解法不一致, 通过 LeetCode 测试, 正确性及性能有待验证
func PreorderTraversal2(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	if root == nil {
		return
	}

	for root != nil || len(stack) > 0 {
		if root != nil {
			root.print()
			stack = append(stack, root)
			root = root.left
		} else {
			parent := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = parent.right
		}
	}
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

// 中序遍历 非递归
func InorderTraversal2(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.left
		} else {
			parent := stack[len(stack)-1]
			parent.print()
			stack = stack[:len(stack)-1]
			root = parent.right
		}
	}
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

func PostorderTraversal2(node *TreeNode) {

}

func (n TreeNode) print() {
	fmt.Print(" => ", n.val)
}
