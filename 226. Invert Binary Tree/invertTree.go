package _26__Invert_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			temp := node.Left
			node.Left = node.Right
			node.Right = temp
			stack = append(stack, node.Left)
			stack = append(stack, node.Right)
		}

	}
	return root
}
