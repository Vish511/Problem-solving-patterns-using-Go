package twopointers

/*
Problem Statement #
Given a sorted array, create a new array containing squares of all the number of the input array in the sorted order.

Example 1:

Input: [-2, -1, 0, 2, 3]
Output: [0, 1, 4, 4, 9]
Example 2:

Input: [-3, -1, 0, 1, 2]
Output: [0 1 1 4 9]
*/

func SquaringASortedArray(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	var left int
	var right int = len(arr) - 1
	var result []int

	for i := 0; i < len(arr); i++ {
		var leftNumSq = arr[left] * arr[left]
		var rightNumSq = arr[right] * arr[right]
		if rightNumSq >= leftNumSq {
			result = append([]int{rightNumSq}, result...)
			right--
		} else {
			result = append([]int{leftNumSq}, result...)
			left++
		}
	}

	return result
}
