package Level_Order_Bottom

import (
	"reflect"
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {

	results := [][][]int{
		{{15, 7}, {9, 20}, {3}},
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
		if ret := levelOrderBottom(&tests[i]); !reflect.DeepEqual(ret, results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}

}
