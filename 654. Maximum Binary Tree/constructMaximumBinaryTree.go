package _54__Maximum_Binary_Tree

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := nums[0]
	index := 0
	for i := range nums {
		if max < nums[i] {
			max = nums[i]
			index = i
		}
	}
	left := nums[:index]
	right := nums[index+1:]
	root := TreeNode{max, constructMaximumBinaryTree(left), constructMaximumBinaryTree(right)}
	return &root
}
