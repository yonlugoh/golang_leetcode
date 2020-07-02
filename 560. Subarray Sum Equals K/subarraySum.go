package _60__Subarray_Sum_Equals_K

func subarraySum(nums []int, k int) int {

	sumMap := make(map[int]int)
	sumMap[0] = 1
	count := 0
	cumSum := 0
	for _, num := range nums {
		cumSum += num

		if prevSum, ok := sumMap[cumSum-k]; ok {
			count += prevSum
		}

		if curSum, ok := sumMap[cumSum]; ok {
			sumMap[cumSum] = curSum + 1
		} else {
			sumMap[cumSum] = 1
		}
	}
	return count

}
