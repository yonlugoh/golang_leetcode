package _30__Kth_Smallest_Element_in_a_BST

import (
	"testing"
)

func TestKthSmallest(t *testing.T) {

	results := []int{
		3,
		4,
		1,
	}

	root := TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Left.Left = &TreeNode{Val: 1}

	tests := []TreeNode{
		root,
		root,
		root,
	}

	kth := []int{
		3,
		4,
		1,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := kthSmallest(&tests[i], kth[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}

}
