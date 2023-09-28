package slidingwindow

import "dsa-patterns/helpers"

/*
Problem Statement #
Given an array of positive numbers and a positive number ‘k’, find the maximum sum of any contiguous subarray of size ‘k’.

Example 1:

Input: [2, 1, 5, 1, 3, 2], k=3
Output: 9
Explanation: Subarray with maximum sum is [5, 1, 3].
Example 2:

Input: [2, 3, 4, 1, 5], k=2
Output: 7
Explanation: Subarray with maximum sum is [3, 4].
*/

func MaxSubArraySumAlternateImpl(k int, arr []int) int {
	initSum := windowSum(k, arr)
	var maxSum int = initSum
	for i := k; i < len(arr); i++ {
		initSum = initSum - arr[i-k]
		initSum = initSum + arr[i]
		maxSum = helpers.Max(maxSum, initSum)
	}
	return maxSum
}

func MaxSubArraySum(k int, arr []int) int {
	var windowStart, windowEnd int
	var maxSum, windowSum int
	for windowEnd = 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd]

		if windowEnd >= k-1 {
			maxSum = helpers.Max(maxSum, windowSum)
			windowSum -= arr[windowStart]
			windowStart += 1
		}
	}
	return maxSum
}

func windowSum(k int, arr []int) int {
	var sum int
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	return sum
}

//Pseudo code
// 1. Calculate initial sum upto k
// 2. Start a loop from k to len of arr
// 3. Subtract the element that went out of the window and add the new element that was added to the window
