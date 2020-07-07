package _13__Arithmetic_Slices

func arithmeticSlices(A []int) int {
	curr := 0
	sum := 0

	for i := 2; i < len(A); i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			curr++
			sum += curr
		} else {
			curr = 0
		}

	}
	return sum
}
