package _8__Subsets

func subsets(nums []int) [][]int {
	length := len(nums)

	res := [][]int{[]int{}}

	for i := 0; i < length; i++ {
		for _, r := range res {
			//create an additional space to store nums[i]
			list := make([]int, len(r)+1)
			// copies existing result r into list, saving last index
			copy(list, r)
			//store nums[i] in last slot
			list[len(r)] = nums[i]
			// add this to the power set
			res = append(res, list)
		}

	}
	return res
}
