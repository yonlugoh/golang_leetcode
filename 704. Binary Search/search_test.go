package _04__Binary_Search

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := [][]int{
		[]int{-1, 0, 3, 5, 9, 12},
		[]int{-1, 0, 3, 5, 9, 12},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99},
	}

	targets := []int{
		9,
		2,
		39,
	}

	results := []int{
		4,
		-1,
		39,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := search(tests[i], targets[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
