package _52__Maximum_Product_Subarray

import (
	"math"
)

func maxProduct(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	// Global maximum variable
	// Space complexity: O(1)
	maxglobal := -math.MaxInt32
	maxlocal := 1
	minlocal := 1

	// Time complexity: O(N)
	for _, num := range nums {
		// two products necessary as we may need to swap maxlocal and minlocal if num is negative
		prod1 := maxlocal * num
		prod2 := minlocal * num
		maxlocal = max(num, max(prod1, prod2))
		minlocal = min(num, min(prod1, prod2))
		maxglobal = max(maxglobal, maxlocal)

	}

	return maxglobal

}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
