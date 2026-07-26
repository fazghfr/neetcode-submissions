/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

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

func minMeetingRooms(intervals []Interval) int {
	h := &MinHeap{}
	heap.Init(h)

	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func (i, j int) bool {
		if intervals[i].start == intervals[j].start {
			return intervals[i].end < intervals[j].end
		}
		return intervals[i].start < intervals[j].start
	})

	for i := 0; i < len(intervals); i++ {
		current := intervals[i]
		if i == 0 {
			heap.Push(h, current.end)
			continue
		}

		root := (*h)[0]

		// no conflict -> reuse room
		if current.start >= root {
			heap.Pop(h)
		} 
		heap.Push(h, current.end)
	}

	return len(*h)
}

// use MinHeap
// if it does not conflict with root, no need new room
// if it does, need new room -> push to heap