package _00__Same_Tree

import (
	"testing"
)

func TestIsSameTree(t *testing.T) {

	p1 := TreeNode{Val: 1}
	p1.Left = &TreeNode{Val: 2}
	p1.Right = &TreeNode{Val: 3}
	q1 := TreeNode{Val: 1}
	q1.Left = &TreeNode{Val: 2}
	q1.Right = &TreeNode{Val: 3}

	p2 := TreeNode{Val: 1}
	p2.Left = &TreeNode{Val: 2}
	q2 := TreeNode{Val: 1}
	q2.Right = &TreeNode{Val: 2}

	p3 := TreeNode{Val: 1}
	p3.Left = &TreeNode{Val: 2}
	p3.Right = &TreeNode{Val: 1}
	q3 := TreeNode{Val: 1}
	q3.Left = &TreeNode{Val: 1}
	q3.Right = &TreeNode{Val: 2}

	p := []TreeNode{
		p1,
		p2,
		p3,
	}
	q := []TreeNode{
		q1,
		q2,
		q3,
	}

	results := []bool{

		true,
		false,
		false,
	}

	caseNum := len(results)
	for i := 0; i < caseNum; i++ {
		if ret := isSameTree(&p[i], &q[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}

}
