package _44__Binary_Tree_Preorder_Traversal

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderRecursive(root *TreeNode) []int {
	res := []int{}

	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		dfs(root.Left, res)
		dfs(root.Right, res)

	}

}
