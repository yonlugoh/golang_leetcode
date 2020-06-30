package _00__Number_of_Islands

import "testing"

func TestNumIslands(t *testing.T) {

	tests := [][][]byte{
		[][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}},
		[][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}},
	}

	results := []int{
		1,
		3,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := numIslands(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
