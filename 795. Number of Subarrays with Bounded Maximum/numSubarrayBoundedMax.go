package _95__Number_of_Subarrays_with_Bounded_Maximum

func numSubarrayBoundedMax(A []int, L int, R int) int {
	res := 0
	lptr := 0
	rptr := 0
	m := 0
	for rptr < len(A) {
		if A[rptr] <= R && A[rptr] >= L {
			m = rptr - lptr + 1

		} else if A[rptr] > R {
			m = 0
			lptr = rptr + 1
		}
		res += m
		rptr++
	}

	return res
}
