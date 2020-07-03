package _44__Binary_Tree_Preorder_Traversal

import (
	"reflect"
	"testing"
)

func TestPreorderIterative(t *testing.T) {

	root := TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 15}

	tests := []TreeNode{
		root,
	}

	results := [][]int{
		[]int{3, 9, 20, 15, 7},
	}
	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := preorderRecursive(&tests[i]); !reflect.DeepEqual(ret, results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}

}
