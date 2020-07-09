package _00__Longest_Increasing_Subsequence

// https://leetcode.com/problems/longest-increasing-subsequence/discuss/74824/javapython-binary-search-onlogn-time-with-explanation
func lengthOfLIS(nums []int) int {
	tails := []int{}
	for i := 0; i < len(nums); i++ {
		tlen := len(tails)
		if tlen == 0 {
			tails = []int{nums[i]}
			continue
		}
		if nums[i] > tails[tlen-1] {
			tails = append(tails, nums[i])
			continue
		}
		// binary search
		start, end, mid := 0, tlen-1, 0
		for start != end {
			mid = start + (end-start)/2
			if nums[i] > tails[mid] {
				start = mid + 1
			} else {
				end = mid
			}
		}
		//where to insert nums[i]
		tails[end] = nums[i]
	}
	return len(tails)
}
