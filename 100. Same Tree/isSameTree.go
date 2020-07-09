package _00__Same_Tree

import "container/list"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tuple struct {
	P *TreeNode
	Q *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	queue := list.New()
	queue.PushBack(Tuple{p, q})
	for queue.Len() > 0 {
		curr := queue.Front().Value.(Tuple)

		p, q = curr.P, curr.Q
		queue.Remove(queue.Front())
		if p == nil && q == nil {
			continue
		} else if p == nil || q == nil || (p.Val != q.Val) {
			return false
		}
		queue.PushBack(Tuple{p.Left, q.Left})
		queue.PushBack(Tuple{p.Right, q.Right})

	}
	return true

}
