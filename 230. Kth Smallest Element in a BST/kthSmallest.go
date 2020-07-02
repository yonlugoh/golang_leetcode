package _30__Kth_Smallest_Element_in_a_BST

/*
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {

	stack := []TreeNode{}

	for true {
		for root != nil {
			stack = append(stack, *root)
			root = root.Left

		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return node.Val
		}
		root = node.Right

	}
	return 0
}
