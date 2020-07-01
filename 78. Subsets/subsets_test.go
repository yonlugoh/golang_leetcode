package _8__Subsets

import (
	"src/github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := [][]int{
		[]int{1, 2, 3},
		[]int{},
		[]int{4, 5, 1, 6, 9},
	}

	results := [][][]int{
		[][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		[][]int{[]int{}},
		[][]int{{}, {4}, {5}, {4, 5}, {1}, {4, 1}, {5, 1}, {4, 5, 1}, {6}, {4, 6}, {5, 6}, {4, 5, 6}, {1, 6}, {4, 1, 6}, {5, 1, 6}, {4, 5, 1, 6}, {9},
			{4, 9}, {5, 9}, {4, 5, 9}, {1, 9}, {4, 1, 9}, {5, 1, 9}, {4, 5, 1, 9}, {6, 9}, {4, 6, 9}, {5, 6, 9}, {4, 5, 6, 9}, {1, 6, 9}, {4, 1, 6, 9},
			{5, 1, 6, 9}, {4, 5, 1, 6, 9}},
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := subsets(tests[i]); !assert.ElementsMatch(t, ret, results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
