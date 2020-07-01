package _5__3Sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := [][]int{
		[]int{-1, 0, 1, 2, -1, -4},
		[]int{},
		[]int{-4, -7, 4, 5, 2, 1, 3, 4, -1, -10, 6},
	}

	results := [][][]int{
		[][]int{{-1, 0, 1}, {-1, -1, 2}},
		[][]int{},
		[][]int{{-10, 4, 6}, {-7, 1, 6}, {-7, 2, 5}, {-7, 3, 4}, {-4, -1, 5}, {-4, 1, 3}},
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := threeSum(tests[i]); !assert.ElementsMatch(t, ret, results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
