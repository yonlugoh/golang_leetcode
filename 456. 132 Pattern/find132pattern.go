package _56__132_Pattern

type Pair struct {
	Max int
	Min int
}

/*
The idea is that we can use a stack to keep track of previous min-max intervals.

Here is the principle to maintain the stack:
For each number num in the array
If stack is empty:
    push a new Pair of num into stack
If stack is not empty:
    if num < stack.peek().min, push a new Pair of num into stack
    if num >= stack.peek().min, we first pop() out the peek element, denoted as last
        if num < last.max, we are done, return true;
        if num >= last.max, we merge num into last, which means last.max = num.
        Once we update last, if stack is empty, we just push back last.
        However, the crucial part is:
        If stack is not empty, the updated last might:
            Entirely covered stack.peek(), i.e. last.min < stack.peek().min (which is always true) && last.max >= stack.peek().max,
in which case we keep popping out stack.peek().
            Form a 1-3-2 pattern, we are done ,return true

So at any time in the stack, non-overlapping Pairs are formed in descending order by their min value,
which means the min value of peek element in the stack is always the min value globally.
*/
func find132pattern(nums []int) bool {
	stack := []Pair{}
	for _, n := range nums {
		if len(stack) == 0 || n < stack[len(stack)-1].Min {
			stack = append(stack, Pair{n, n})
		} else if n > stack[len(stack)-1].Min {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if n < last.Max {
				return true
			} else {
				last.Max = n
				for len(stack) > 0 && n >= stack[len(stack)-1].Max {
					stack = stack[:len(stack)-1]
				}
				if len(stack) > 0 && stack[len(stack)-1].Min < n {
					return true
				}
				stack = append(stack, last)
			}

		}
	}
	return false
}
