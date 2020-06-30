package _02__Binary_Tree_Level_Order_Traversal

import "testing"
import "reflect"

func TestLevelOrder(t *testing.T) {

	results := [][][]int{
		{{3}, {9, 20}, {15, 7}},
	}

	root := TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 15}

	tests := []TreeNode{
		root,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := levelOrder(&tests[i]); !reflect.DeepEqual(ret, results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}

}
