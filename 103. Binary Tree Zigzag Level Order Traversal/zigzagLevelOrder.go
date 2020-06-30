package _03__Binary_Tree_Zigzag_Level_Order_Traversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	stack := Stack{}
	stack.Push(root)
	level := 0
	for len(stack) > 0 {
		// level order traversal
		cur := []int{}
		levelNodes := []*TreeNode{}
		for len(stack) > 0 {
			n := stack.Pop()
			cur = append(cur, n.Val)
			levelNodes = append(levelNodes, n)
		}
		for i := 0; i < len(levelNodes); i++ {
			n := levelNodes[i]
			// push to stack from left to right when level is a even number
			if level%2 == 0 {
				if n.Left != nil {
					stack.Push(n.Left)
				}
				if n.Right != nil {
					stack.Push(n.Right)
				}
			} else {
				// push to stack from right to left when level is an odd number
				if n.Right != nil {
					stack.Push(n.Right)
				}
				if n.Left != nil {
					stack.Push(n.Left)
				}
			}
		}
		res = append(res, cur)
		level++
	}
	return res
}

type Stack []*TreeNode

func (s *Stack) Push(x *TreeNode) {
	*s = append(*s, x)
}

func (s *Stack) Pop() *TreeNode {
	old := *s
	if len(old) == 0 {
		return nil
	}
	tail := old[len(old)-1]
	*s = old[:len(old)-1]
	return tail
}
