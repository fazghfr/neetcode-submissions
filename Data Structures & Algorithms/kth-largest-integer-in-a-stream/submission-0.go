type KthLargest struct {
	k int
	h MinHeap
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[:n-1]

	return item
}

func Constructor(k int, nums []int) KthLargest {
	h := MinHeap{}
	heap.Init(&h)

	obj := KthLargest{
		k: k,
		h: h,
	}

	for _, v := range nums {
		obj.Add(v)
	}

	return obj
}

func (this *KthLargest) Add(val int) int {
	if this.h.Len() == this.k && val <= this.h[0] {
		return this.h[0]
	}

	heap.Push(&this.h, val)

	if this.h.Len() > this.k {
		heap.Pop(&this.h)
	}

	return this.h[0]
}


// idea

// minheap, but we only push the top k elements to the heap
// so that minheap root already is a klargest elements in the bigger picture


// example
// 1 2 3 4 5 6 7  k = 3

// on construction 
// heap -> 5 6 9 (root 5)

// lets say we push 4. 4 does not fit into the heap since it is less than current root. this aligns,
// since 4 itself is not the amongst 3 largest element.

// if we push 7, then we push to the heap
// so heap -> 6 7 9. 6 is now the klargest element