package intervals

type priorityQueue struct {
	values []intervalWithPriority
}

type intervalWithPriority struct {
	Interval
	priority int
}

func NewPriorityQueue() *priorityQueue {
	return &priorityQueue{}
}

func (h *priorityQueue) Insert(interval Interval, priority int) {
	h.values = append(h.values, intervalWithPriority{Interval: interval, priority: priority})
	h.bubbleUp()
}

func (h *priorityQueue) Peek() (bool, Interval) {
	if len(h.values) < 1 {
		return false, Interval{}
	}
	return true, h.values[0].Interval
}

func (h *priorityQueue) bubbleUp() bool {
	if len(h.values) < 2 {
		return false
	}
	var currIdx = len(h.values) - 1
	var currElem = h.values[currIdx]
	for true {
		parentIdx := int(currIdx - 1/2)
		parentElem := h.values[parentIdx]

		if parentElem.priority <= currElem.priority { //Parent is smaller than current element, which indicates that we are at the right place
			break
		}

		h.values[parentIdx], h.values[currIdx] = h.values[currIdx], h.values[parentIdx]
		currIdx = parentIdx
	}
	return true
}

func (h *priorityQueue) ExtractMin() (bool, Interval) {
	if len(h.values) < 1 {
		return false, Interval{}
	}
	smallestElem := h.values[0]
	h.values[0], h.values[len(h.values)-1] = h.values[len(h.values)-1], h.values[0]
	h.values = h.values[0 : len(h.values)-1]
	h.sinkDown()
	return true, smallestElem.Interval
}

func (h *priorityQueue) sinkDown() bool {
	if len(h.values) <= 1 {
		return false
	}
	var currIdx int
	var currElem = h.values[currIdx]
	for true {
		leftChildIdx := int(2*currIdx + 1)
		rightChildIdx := int(2*currIdx + 2)
		var leftElem intervalWithPriority
		var rightElem intervalWithPriority
		swapIdx := -1
		if leftChildIdx < len(h.values) {
			leftElem = h.values[leftChildIdx]
			if leftElem.priority < currElem.priority {
				swapIdx = leftChildIdx
			}
		}

		if rightChildIdx < len(h.values) {
			rightElem = h.values[rightChildIdx]
			if (swapIdx == -1 && rightElem.priority < currElem.priority) || (swapIdx != -1 && rightElem.priority < leftElem.priority) {
				swapIdx = rightChildIdx
			}
		}

		if swapIdx == -1 {
			break
		}

		h.values[swapIdx], h.values[currIdx] = h.values[currIdx], h.values[swapIdx]
		currIdx = swapIdx
	}
	return true
}

//[10, 3, 2, 6, 7]
//n-1/2
