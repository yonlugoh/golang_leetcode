package _54__Maximum_Binary_Tree

import (
	"reflect"
	"testing"
)

func TestConstructMaximumBinaryTree(t *testing.T) {

	tests := [][]int{
		[]int{3, 2, 1, 6, 0, 5},
	}

	results := [][][]int{
		{{6}, {3, 5}, {2, 0}, {1}},
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := constructMaximumBinaryTree(tests[i]); !reflect.DeepEqual(levelOrder(ret), results[i]) {
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
