package twopointers

import (
	"sort"
)

/*
Problem Statement #
Given an array arr of unsorted numbers and a target sum, count all triplets in it such that arr[i] + arr[j] + arr[k] < target where i, j, and k are three different indices. Write a function to return the count of such triplets.

Example 1:

Input: [-1, 0, 2, 3], target=3
Output: 2
Explanation: There are two triplets whose sum is less than the target: [-1, 0, 3], [-1, 0, 2]
Example 2:

Input: [-1, 4, 2, 1, 3], target=5
Output: 4
Explanation: There are four triplets whose sum is less than the target:
   [-1, 1, 4], [-1, 1, 3], [-1, 1, 2], [-1, 2, 3]
*/

func TripletSumLessThanTarget(arr []int, target int) [][]int {
	if len(arr) == 0 {
		return [][]int{}
	}
	var left int
	var right int = len(arr) - 1
	var result [][]int
	sort.Ints(arr)
	for i := 0; i < len(arr)-2; i++ {
		left = i + 1
		for left < right {
			currSum := arr[i] + arr[left] + arr[right]
			if currSum < target {
				result = append(result, []int{arr[i], arr[left], arr[right]})
			}
			if right == left+1 {
				left++
				right = len(arr) - 1
			}
			right--

		}
	}
	return result
}

// [-1,1,2,3,4]
