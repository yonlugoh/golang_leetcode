package _30__Surrounded_Regions

import (
	"reflect"
	"testing"
)

func TestSolveBoard(t *testing.T) {

	tests := [][][]byte{
		[][]byte{{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'}},
		[][]byte{{'X', 'O', 'X', 'O', 'X', 'O', 'O', 'O', 'X', 'O'},
			{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'O', 'O', 'X'},
			{'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'X'},
			{'O', 'O', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'X'},
			{'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O', 'O'},
			{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'O'},
			{'X', 'O', 'X', 'O', 'O', 'X', 'X', 'O', 'X', 'O'},
			{'X', 'X', 'O', 'X', 'X', 'O', 'X', 'O', 'O', 'X'},
			{'O', 'O', 'O', 'O', 'X', 'O', 'X', 'O', 'X', 'O'},
			{'X', 'X', 'O', 'X', 'X', 'X', 'X', 'O', 'O', 'O'}},
		[][]byte{{'O', 'X', 'X', 'O', 'X'},
			{'X', 'O', 'O', 'X', 'O'},
			{'X', 'O', 'X', 'O', 'X'},
			{'O', 'X', 'O', 'O', 'O'},
			{'X', 'X', 'O', 'X', 'O'}},
	}

	results := [][][]byte{
		[][]byte{{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'X'}},
		[][]byte{{'X', 'O', 'X', 'O', 'X', 'O', 'O', 'O', 'X', 'O'},
			{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'O', 'O', 'X'},
			{'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'X'},
			{'O', 'O', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'X'},
			{'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O', 'O'},
			{'X', 'O', 'O', 'X', 'X', 'X', 'X', 'X', 'X', 'O'},
			{'X', 'O', 'X', 'X', 'X', 'X', 'X', 'O', 'X', 'O'},
			{'X', 'X', 'O', 'X', 'X', 'X', 'X', 'O', 'O', 'X'},
			{'O', 'O', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O'},
			{'X', 'X', 'O', 'X', 'X', 'X', 'X', 'O', 'O', 'O'}},
		[][]byte{{'O', 'X', 'X', 'O', 'X'},
			{'X', 'X', 'X', 'X', 'O'},
			{'X', 'X', 'X', 'O', 'X'},
			{'O', 'X', 'O', 'O', 'O'},
			{'X', 'X', 'O', 'X', 'O'}},
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if solveBoard(tests[i]); !reflect.DeepEqual(tests[i], results[i]) {
			t.Fatalf("case %d fails: %v\n", i, tests[i])
		}
	}
}
