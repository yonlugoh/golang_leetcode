package _5__3Sum

import "sort"

func threeSum(nums []int) [][]int {

	res := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	length := len(nums)
	for i := 0; i < length-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := length - 1
		for left < right {
			total := nums[i] + nums[left] + nums[right]
			if total < 0 {
				left++
			} else if total > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return res

}
