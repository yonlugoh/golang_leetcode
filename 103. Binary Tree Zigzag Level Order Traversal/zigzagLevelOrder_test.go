package _03__Binary_Tree_Zigzag_Level_Order_Traversal

import (
	"reflect"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {

	results := [][][]int{
		{{3}, {20, 9}, {15, 7}, {4, 5}},
	}

	root := TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Left.Left = &TreeNode{Val: 5}
	root.Right.Left.Right = &TreeNode{Val: 4}

	tests := []TreeNode{
		root,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := zigzagLevelOrder(&tests[i]); !reflect.DeepEqual(ret, results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}

}
