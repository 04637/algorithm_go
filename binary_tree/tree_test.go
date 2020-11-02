package binary_tree

import (
	"fmt"
	"testing"
)

func TestTraversal(t *testing.T) {
	tree := generateTree()
	fmt.Println("preorder: ")
	PreorderTraversal(tree)
	fmt.Println("\n--------------------------------")
	fmt.Println("inorder: ")
	InorderTraversal(tree)
	fmt.Println("\n--------------------------------")
	fmt.Println("postorder: ")
	PostorderTraversal(tree)
	fmt.Println("\n--------------------------------")
	fmt.Println()
}

func generateTree() *TreeNode {
	root := &TreeNode{}
	root.val = 5
	root.left = &TreeNode{
		val: 3,
	}
	root.right = &TreeNode{
		val: 8,
	}
	// root 3
	curRoot := root.left
	curRoot.left = &TreeNode{
		val: 1,
	}
	curRoot.right = &TreeNode{
		val: 4,
	}
	// root 1
	curRoot = curRoot.left
	curRoot.left = &TreeNode{
		val: 0,
	}
	curRoot.right = &TreeNode{
		val: 2,
	}
	// root 8
	curRoot = root.right
	curRoot.left = &TreeNode{
		val: 6,
	}
	curRoot.right = &TreeNode{
		val: 10,
	}
	// root 6
	curLeft := curRoot.left
	curLeft.right = &TreeNode{
		val: 7,
	}
	// root 10
	curRight := curRoot.right
	curRight.left = &TreeNode{
		val: 9,
	}
	return root
}
