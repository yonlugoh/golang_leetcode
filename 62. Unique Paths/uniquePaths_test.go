package _2__Unique_Paths

import "testing"

func TestUniquePaths(t *testing.T) {
	m := []int{
		1,
		3,
		7,
		20,
	}
	n := []int{
		1,
		2,
		3,
		11,
	}

	results := []int{
		1,
		3,
		28,
		20030010,
	}

	caseNum := len(m)
	for i := 0; i < caseNum; i++ {
		if ret := uniquePaths(m[i], n[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
