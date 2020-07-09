package _26__Invert_Binary_Tree

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {

	results := [][][]int{
		{{4}, {7, 2}, {9, 6, 3, 1}},
	}

	root := TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 6}

	root.Left.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 1}

	tests := []TreeNode{
		root,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := invertTree(&tests[i]); !reflect.DeepEqual(levelOrder(ret), results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}

}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	level := []*TreeNode{root}

	for root != nil && len(level) > 0 {
		currentVals := []int{}
		nextLevel := []*TreeNode{}
		for _, node := range level {
			currentVals = append(currentVals, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}

		}
		res = append(res, currentVals)
		level = nextLevel

	}

	return res

}
