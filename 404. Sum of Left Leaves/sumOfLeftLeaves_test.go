package _04__Sum_of_Left_Leaves

import (
	"reflect"
	"testing"
)

func TestSumOfLeftLeaves(t *testing.T) {

	results := []int{
		24,
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
		if ret := sumOfLeftLeaves(&tests[i]); !reflect.DeepEqual(ret, results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}

}
