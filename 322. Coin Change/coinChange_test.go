package _22__Coin_Change

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := [][]int{
		[]int{1, 2, 5},
		[]int{2},
		[]int{1, 5, 20, 100},
	}
	targets := []int{
		11,
		3,
		313,
	}
	results := []int{
		3,
		-1,
		8,
	}
	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := coinChange(tests[i], targets[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
