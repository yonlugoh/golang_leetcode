package _30__Course_Schedule_III

import (
	"container/heap"
	"sort"
)

// An IntHeap is a min-heap of ints.
// https://golang.org/pkg/container/heap/
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func scheduleCourse3(courses [][]int) int {

	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	current := 0
	maxheap := &IntHeap{}
	heap.Init(maxheap)
	for _, i := range courses {
		time := i[0]
		end := i[1]
		current += time
		// push negative values for maxheap
		heap.Push(maxheap, -time)
		// greedy approach, pop from maxheap if we can't accomodate this course
		if current > end {
			nt := (*maxheap)[0]
			// nt is a negative value, which is then added on to current time.
			// (ie, we remove the course that contributed the max duration to our current time)
			current += nt
			heap.Pop(maxheap)
		}
	}
	return maxheap.Len()
}
