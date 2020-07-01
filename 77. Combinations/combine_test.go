package _7__Combinations

import (
	"src/github.com/stretchr/testify/assert"
	"testing"
)

func TestCombine(t *testing.T) {

	n := []int{
		4,
		6,
		2,
	}
	k := []int{
		2,
		4,
		1,
	}
	results := [][][]int{
		[][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		[][]int{{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 3, 6}, {1, 2, 4, 5}, {1, 2, 4, 6}, {1, 2, 5, 6}, {1, 3, 4, 5}, {1, 3, 4, 6}, {1, 3, 5, 6},
			{1, 4, 5, 6}, {2, 3, 4, 5}, {2, 3, 4, 6}, {2, 3, 5, 6}, {2, 4, 5, 6}, {3, 4, 5, 6}},
		[][]int{{1}, {2}},
	}

	caseNum := len(results)
	for i := 0; i < caseNum; i++ {
		if ret := combine(n[i], k[i]); !assert.ElementsMatch(t, ret, results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
