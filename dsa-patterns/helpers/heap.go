package helpers

import "math"

type minHeap struct {
	values []int
}

func NewHeap() *minHeap {
	return &minHeap{}
}

func (h *minHeap) Insert(val int) {
	h.values = append(h.values, val)
	h.bubbleUp()
}

func (h *minHeap) bubbleUp() []int {
	var currIdx = len(h.values) - 1
	var currElem = h.values[currIdx]
	for true {
		parentIdx := int(currIdx - 1/2)
		parentElem := h.values[parentIdx]

		if parentElem < currElem { //Parent is smaller than current element, which indicates that we are at the right place
			break
		}

		h.values[parentIdx], h.values[currIdx] = h.values[currIdx], h.values[parentIdx]
		currIdx = parentIdx
	}
	return h.values
}

func (h *minHeap) ExtractMin() (bool, int) {
	if len(h.values) < 1 {
		return false, math.MinInt
	}
	smallestElem := h.values[0]
	h.values[0], h.values[len(h.values)-1] = h.values[len(h.values)-1], h.values[0]
	h.values = h.values[0 : len(h.values)-1]
	h.sinkDown()
	return true, smallestElem
}

func (h *minHeap) sinkDown() bool {
	if len(h.values) <= 1 {
		return false
	}
	var currIdx int
	var currElem = h.values[currIdx]
	for true {
		leftChildIdx := int(2*currIdx + 1)
		rightChildIdx := int(2*currIdx + 2)
		var leftElem int
		var rightElem int
		swapIdx := -1
		if leftChildIdx < len(h.values) {
			leftElem = h.values[leftChildIdx]
			if leftElem < currElem {
				swapIdx = leftChildIdx
			}
		}

		if rightChildIdx < len(h.values) {
			rightElem = h.values[rightChildIdx]
			if (swapIdx == -1 && rightElem > currElem) || (swapIdx != -1 && rightElem > leftElem) {
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
