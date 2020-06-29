package _5__Jump_Game_II

func jump2(nums []int) int {
	jumps := 0
	curEnd := 0
	curFarthest := 0
	for i := 0; i < len(nums)-1; i++ {
		curFarthest = max(curFarthest, i+nums[i])
		if i == curEnd {
			jumps++
			curEnd = curFarthest
		}
	}
	return jumps

}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
