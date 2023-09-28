package twopointers

import (
	"math"
	"sort"
)

/*
Problem Statement #
Given an array of unsorted numbers and a target number, find a triplet in the array whose sum is as close to the target number as possible, return the sum of the triplet. If there are more than one such triplet, return the sum of the triplet with the smallest sum.

Example 1:

Input: [-2, 0, 1, 2], target=2
Output: 1
Explanation: The triplet [-2, 1, 2] has the closest sum to the target.
Example 2:

Input: [-3, -1, 1, 2], target=1
Output: 0
Explanation: The triplet [-3, 1, 2] has the closest sum to the target.
Example 3:

Input: [1, 0, 1, 1], target=100
Output: 3
Explanation: The triplet [1, 1, 1] has the closest sum to the target.
*/

func TripletSumCloseToTarget(arr []int, target int) []int {
	if len(arr) == 0 {
		return []int{}

	}
	var result []int
	var closestSum int = math.MaxInt

	sort.Ints(arr)
	for i := 0; i < len(arr)-2; i++ {
		var left = i + 1
		var right int = len(arr) - 1
		for left < right {
			currSum := arr[i] + arr[left] + arr[right]
			absoluteDiff := int(math.Abs(float64(currSum - target)))
			if absoluteDiff < closestSum {
				closestSum = absoluteDiff
				result = append([]int{}, arr[i], arr[left], arr[right])
			}
			if currSum == target {
				result = append([]int{}, arr[i], arr[left], arr[right])
				left++
				right--
				break
			} else if currSum > target {
				right--
			} else {
				left++
			}
		}

	}
	return result
}
