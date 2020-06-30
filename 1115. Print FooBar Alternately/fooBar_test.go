package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFooBar(t *testing.T) {
	tests := []int{
		5,
		10,
	}

	results := []string{
		"foobarfoobarfoobarfoobarfoobar",
		"foobarfoobarfoobarfoobarfoobarfoobarfoobarfoobarfoobarfoobar",
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		fooBar(tests[i]) // stdout from foobar is captured

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		if string(out) != results[i] {
			t.Fatalf("case %d fails: %v\n", i, string(out))
		}
	}
}
