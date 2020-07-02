package _17__Merge_Two_Binary_Trees

import (
	"reflect"
	"testing"
)

func TestMergeTrees(t *testing.T) {

	tree1 := TreeNode{1, nil, nil}
	tree1.Left = &TreeNode{3, nil, nil}
	tree1.Left.Left = &TreeNode{5, nil, nil}
	tree1.Right = &TreeNode{2, nil, nil}

	tree2 := TreeNode{2, nil, nil}
	tree2.Left = &TreeNode{1, nil, nil}
	tree2.Left.Right = &TreeNode{4, nil, nil}
	tree2.Right = &TreeNode{3, nil, nil}
	tree2.Right.Right = &TreeNode{7, nil, nil}

	tree3 := mergeTrees(&tree1, &tree2)
	results := [][][]int{
		{{3}, {4, 5}, {5, 4, 7}},
	}
	caseNum := len(results)
	for i := 0; i < caseNum; i++ {
		if ret := levelOrder(tree3); !reflect.DeepEqual(ret, results[i]) {
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
