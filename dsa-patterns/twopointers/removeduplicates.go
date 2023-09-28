package twopointers

/*
Problem Statement #
Given an array of sorted numbers, remove all duplicates from it. You should not use any extra space; after removing the duplicates in-place return the new length of the array.

Example 1:

Input: [2, 3, 3, 3, 6, 9, 9]
Output: 4
Explanation: The first four elements after removing the duplicates will be [2, 3, 6, 9].
Example 2:

Input: [2, 2, 2, 11]
Output: 2
Explanation: The first two elements after removing the duplicates will be [2, 11].
*/

func RemoveDuplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var left int
	var right int = 1
	for right < len(arr) {
		if arr[left] != arr[right] {
			arr[left+1] = arr[right]
			left++
			right++
		} else {
			right++
		}
	}
	return left + 1
}
