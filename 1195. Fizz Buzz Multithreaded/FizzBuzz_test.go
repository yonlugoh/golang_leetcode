package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFizzBuzzMultithreaded(t *testing.T) {
	tests := []int{
		15,
		30,
	}

	results := []string{
		"1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz ",
		"1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz 16 17 fizz 19 buzz fizz 22 23 fizz buzz 26 fizz 28 29 fizzbuzz ",
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		FizzBuzzMultithreaded(tests[i]) // stdout from foobar is captured

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		if string(out) != results[i] {
			t.Fatalf("case %d fails: %v\n", i, string(out))
		}
	}
}
