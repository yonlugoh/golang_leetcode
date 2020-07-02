package _13__House_Robber_II

func robber(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		return nums[0]
	} else {
		// perform the helper function
		return max(helper(nums[1:]), helper(nums[:length-1]))
	}
}

// using 2 pointers, return the max amount robbed in this slice
func helper(nums []int) int {
	prevMax := 0
	curMax := 0
	for _, i := range nums {
		temp := curMax
		// prevMax+i => this means to rob current element i, added with the prevMax (not robbing prev element)
		// curMax => do not rob current element i, as robbing previous element contributed to curMax
		curMax = max(prevMax+i, curMax)
		prevMax = temp

	}

	return curMax
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
