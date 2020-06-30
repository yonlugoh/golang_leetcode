package _95__Find_Median_from_Data_Stream

import "container/heap"

type MedianFinder struct {
	Maxheap IntHeap
	Minheap IntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	s := &IntHeap{}
	l := &IntHeap{}
	heap.Init(s)
	heap.Init(l)
	return MedianFinder{*s, *l}

}

func (this *MedianFinder) AddNum(num int) {
	if (len(this.Maxheap)) == 0 {
		heap.Push(&this.Maxheap, -num)
		return
	}
	if num <= -this.Maxheap[0] {
		heap.Push(&this.Maxheap, -num)
	} else {
		heap.Push(&this.Minheap, num)
	}

	if len(this.Maxheap)-len(this.Minheap) == 2 {
		toPush := this.Maxheap[0]
		heap.Pop(&this.Maxheap)
		heap.Push(&this.Minheap, -toPush)
	} else if len(this.Minheap)-len(this.Maxheap) == 2 {
		toPush := this.Minheap[0]
		heap.Pop(&this.Minheap)
		heap.Push(&this.Maxheap, -toPush)

	}

}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.Maxheap) == len(this.Minheap) {
		return float64(this.Minheap[0]-this.Maxheap[0]) / 2.0
	} else if len(this.Maxheap) < len(this.Minheap) {
		return float64(this.Minheap[0])
	} else {
		return float64(-this.Maxheap[0])
	}

}

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

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
