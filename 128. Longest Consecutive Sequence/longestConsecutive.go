package _28__Longest_Consecutive_Sequence

func longestConsecutive(nums []int) int {

	set := make(map[int]int)
	for _, num := range nums {
		set[num] = 0
	}
	if len(set) <= 1 {
		return len(set)
	}
	best := 1
	for _, x := range nums {
		_, found := set[x-1]
		// if x-1 is not found, this means we can start from x and increment, to see if it forms the longest streak
		if !found {
			streak := 1
			y := x + 1
			_, found = set[y]
			// keep incrementing to extend streak
			for found {
				streak++
				y++
				best = max(best, streak)
				_, found = set[y]
			}
		}
	}
	return best
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
