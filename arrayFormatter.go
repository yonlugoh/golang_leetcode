package golang_leetcode

import (
	"fmt"
)

// format python's arrays to golang slice, in string form
func arrayFormatter(s string) {
	res := ""
	for _, c := range s {
		if string(c) == "[" {
			res = res + "{"
		} else if string(c) == "]" {
			res = res + "}"
		} else {
			res = res + string(c)
		}

	}
	fmt.Print(res)
}
