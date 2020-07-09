package _438__Longest_Continuous_Subarray_With_Absolute_Diff

import "container/list"

func longestSubarray(nums []int, limit int) int {
	// Video explanation: https://youtu.be/LDFZm4iB7tA

	res := 0
	n := len(nums)
	minQueue := list.New()
	maxQueue := list.New()

	l := 0
	r := 0
	for r < n {
		cur := nums[r]
		for minQueue.Len() > 0 && nums[minQueue.Back().Value.(int)] >= cur {
			minQueue.Remove(minQueue.Back())
		}
		minQueue.PushBack(r)
		for maxQueue.Len() > 0 && nums[maxQueue.Back().Value.(int)] <= cur {
			maxQueue.Remove(maxQueue.Back())
		}

		maxQueue.PushBack(r)

		mini := nums[minQueue.Front().Value.(int)]
		maxi := nums[maxQueue.Front().Value.(int)]

		if maxi-mini > limit {
			l++
			if l > minQueue.Front().Value.(int) {
				minQueue.Remove(minQueue.Front())
			}
			if l > maxQueue.Front().Value.(int) {
				maxQueue.Remove(maxQueue.Front())
			}
		} else {
			if r-l+1 > res {
				res = r - l + 1
			}
			r++
		}

	}

	return res
}
