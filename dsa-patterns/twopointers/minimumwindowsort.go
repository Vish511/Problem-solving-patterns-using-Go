package twopointers

import (
	"dsa-patterns/helpers"
	"math"
)

/*
Minimum Window Sort (medium) #
Given an array, find the length of the smallest subarray in it which when sorted will sort the whole array.

Example 1:

Input: [1, 2, 5, 3, 7, 10, 9, 12]
Output: 5
Explanation: We need to sort only the subarray [5, 3, 7, 10, 9] to make the whole array sorted
Example 2:

Input: [1, 3, 2, 0, -1, 7, 10]
Output: 5
Explanation: We need to sort only the subarray [1, 3, 2, 0, -1] to make the whole array sorted
Example 3:

Input: [1, 2, 3]
Output: 0
Explanation: The array is already sorted
Example 4:

Input: [3, 2, 1]
Output: 3
Explanation: The whole array needs to be sorted.
*/

func MinWindowSort(arr []int) []int {
	if len(arr) < 2 {
		return []int{}
	}
	var low int
	var high int = len(arr) - 1
	var lowestInWindow int = math.MaxInt
	var highestInWindow int = -1
	for low < len(arr)-1 && arr[low] < arr[low+1] {
		low++
	}

	if low == len(arr)-1 {
		return []int{}
	}

	for high > 0 && arr[high] >= arr[high-1] {
		high--
	}
	for i := low; i <= high; i++ {
		lowestInWindow = helpers.Min(lowestInWindow, arr[i])
		highestInWindow = helpers.Max(highestInWindow, arr[i])
	}

	for low > 0 && arr[low-1] > lowestInWindow {
		low--
	}
	for high < len(arr) && arr[high] > highestInWindow {
		high++
	}

	return arr[low : high+1]
}
