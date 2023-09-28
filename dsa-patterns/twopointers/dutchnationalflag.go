package twopointers

/*

Problem Statement #
Given an array containing 0s, 1s and 2s, sort the array in-place. You should treat numbers of the array as objects, hence, we can’t count 0s, 1s, and 2s to recreate the array.

The flag of the Netherlands consists of three colors: red, white and blue; and since our input array also consists of three different numbers that is why it is called Dutch National Flag problem.

Example 1:

Input: [1, 0, 2, 1, 0]
Output: [0 0 1 1 2]
Example 2:

Input: [2, 2, 0, 1, 2, 0]
Output: [0 0 1 2 2 2 ]

*/

func DutchNationalFlag(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	var low int
	var high int = len(arr) - 1
	var i int

	for i <= high {
		if arr[i] == 0 {
			arr[i], arr[low] = arr[low], arr[i]
			i++
			low++
		} else if arr[i] == 1 {
			i++
		} else {
			arr[i], arr[high] = arr[high], arr[i]
			high--
		}
	}
	return arr
}
