package _98__House_Robber

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	prev := 0
	curr := 0
	for _, num := range nums {
		temp := curr
		curr = max(curr, prev+num)
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
