package _40__Delete_and_Earn

func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	points := make([]int, 10001)
	for i := 0; i < 10001; i++ {
		points[i] = 0
	}
	for _, num := range nums {
		points[num] += num
	}
	prev := 0
	curr := 0
	for _, point := range points {
		temp := curr
		curr = max(curr, prev+point)
		prev = temp
	}

	return curr
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
