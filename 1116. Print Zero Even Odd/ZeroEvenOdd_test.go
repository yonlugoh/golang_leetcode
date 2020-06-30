package _116__Print_Zero_Even_Odd

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestZeroEvenOdd(t *testing.T) {
	tests := []int{
		8,
		20,
	}

	results := []string{
		"0102030405060708",
		"010203040506070809010011012013014015016017018019020",
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		ZeroEvenOdd(tests[i]) // stdout from foobar is captured

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		if string(out) != results[i] {
			t.Fatalf("case %d fails: %v\n", i, string(out))
		}
	}
}
