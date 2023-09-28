package twopointers

import (
	"sort"
)

/*
Problem Statement #
Given an array of unsorted numbers, find all unique triplets in it that add up to zero.

Example 1:

Input: [-3, 0, 1, 2, -1, 1, -2]
Output: [-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]
Explanation: There are four unique triplets whose sum is equal to zero.
Example 2:

Input: [-5, 2, -1, -2, 3]
Output: [[-5, 2, 3], [-2, -1, 3]]
Explanation: There are two unique triplets whose sum is equal to zero.

*/

func TripletSumToZero(arr []int) [][]int {
	if len(arr) == 0 {
		return [][]int{}

	}
	var requiredSum int
	var result [][]int
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		remainingSum := requiredSum - arr[i]
		pairWithSum := pairsWithTargetSum(arr[i+1:], remainingSum)
		if len(pairWithSum) > 0 {
			for j := 0; j < len(pairWithSum); j++ {
				result = append(result, append([]int{arr[i]}, pairWithSum[j]...))
			}
		}
	}
	return result
}

func pairsWithTargetSum(arr []int, target int) [][]int {
	if len(arr) == 0 {
		return [][]int{}
	}
	var left int
	var right int = len(arr) - 1
	var result [][]int
	for left < right {
		currSum := arr[left] + arr[right]
		if currSum == target {
			result = append(result, []int{arr[left], arr[right]})
			left++
			right--
		} else if currSum > target {
			right--
		} else {
			left++
		}
	}
	return result
}
