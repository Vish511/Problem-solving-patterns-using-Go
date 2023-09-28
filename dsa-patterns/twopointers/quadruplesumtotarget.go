package twopointers

import (
	"sort"
)

/*
Quadruple Sum to Target (medium) #
Given an array of unsorted numbers and a target number, find all unique quadruplets in it, whose sum is equal to the target number.

Example 1:

Input: [4, 1, 2, -1, 1, -3], target=1
Output: [-3, -1, 1, 4], [-3, 1, 1, 2]
Explanation: Both the quadruplets add up to the target.
Example 2:

Input: [2, 0, -1, 1, -2, 2], target=2
Output: [-2, 0, 2, 2], [-1, 0, 1, 2]
Explanation: Both the quadruplets add up to the target
*/

func QuadrupleSumToTarget(arr []int, target int) [][]int {
	if len(arr) == 0 {
		return [][]int{}
	}
	sort.Ints(arr)
	var left, right int
	var result = [][]int{}
	for i := 0; i < len(arr)-4; i++ {
		for j := i + 1; j < len(arr)-3; j++ {
			left = j + 1
			right = len(arr) - 1
			for left < right {
				currSum := arr[i] + arr[j] + arr[left] + arr[right]
				if currSum > target {
					right--
				} else if currSum < target {
					left++
				} else {
					result = append(result, []int{arr[i], arr[j], arr[left], arr[right]})
					left++
					right--
				}
			}
		}

	}

	return result
}
