package _094__Car_Pooling

func carPooling(trips [][]int, capacity int) bool {

	// using a difference array
	// https://wcipeg.com/wiki/Prefix_sum_array_and_difference_array
	arr := make([]int, 1001)
	for _, i := range trips {
		people := i[0]
		start := i[1]
		end := i[2]
		// these 2 operations update at index start and end, to update the whole range between start and end
		arr[start] += people
		arr[end] -= people
	}
	curr := 0
	for j := 0; j < 1001; j++ {
		curr += arr[j]
		// exceeded capacity at any point
		if curr > capacity {
			return false
		}
	}
	// never exceeded capacity
	return true

}
