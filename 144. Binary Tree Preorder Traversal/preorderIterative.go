package _44__Binary_Tree_Preorder_Traversal

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}
	return res
}
