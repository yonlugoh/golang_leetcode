package _4__Binary_Tree_Inorder_Traversal

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {

	root := TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	tests := []TreeNode{
		root,
	}

	results := [][]int{
		[]int{2, 3, 4, 5, 6, 7},
	}
	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := inorderTraversal(&tests[i]); !reflect.DeepEqual(ret, results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}

}
