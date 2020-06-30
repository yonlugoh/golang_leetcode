package _02__Binary_Tree_Level_Order_Traversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	level := []*TreeNode{root}

	for root != nil && len(level) > 0 {
		currentVals := []int{}
		nextLevel := []*TreeNode{}
		for _, node := range level {
			currentVals = append(currentVals, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}

		}
		res = append(res, currentVals)
		level = nextLevel

	}

	return res

}
