package _50__Delete_Node_in_a_BST

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {

	root := TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	tests := []TreeNode{
		root,
	}
	targets := []int{
		3,
	}

	results := [][]int{
		[]int{2, 4, 5, 6, 7},
	}
	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := deleteNode(&tests[i], targets[i]); !reflect.DeepEqual(inorderTraversal(ret), results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}

}

func inorderTraversal(root *TreeNode) []int {

	res := []int{}
	if root == nil {
		return res
	}

	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res

}
