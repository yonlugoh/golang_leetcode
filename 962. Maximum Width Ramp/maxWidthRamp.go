package _62__Maximum_Width_Ramp

func maxWidthRamp(A []int) int {

	n := len(A)
	rMax := make([]int, n)
	rMax[n-1] = A[n-1]

	for i := n - 2; i >= 0; i-- {
		rMax[i] = max(rMax[i+1], A[i])
	}

	left := 0
	right := 0
	ans := 0

	for right < n {
		for left < right && A[left] > rMax[right] {
			left++
		}

		ans = max(ans, right-left)
		right++

	}
	return ans

}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
