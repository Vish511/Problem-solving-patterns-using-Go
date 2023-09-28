package slidingwindow

import "dsa-patterns/helpers"

/*
Problem Statement #
Given an array containing 0s and 1s, if you are allowed to replace no more than ‘k’ 0s with 1s, find the length of the longest contiguous subarray having all 1s.

Example 1:

Input: Array=[0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1], k=2
Output: 6
Explanation: Replace the '0' at index 5 and 8 to have the longest contiguous subarray of 1s having length 6.
Example 2:

Input: Array=[0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1], k=3
Output: 9
Explanation: Replace the '0' at index 6, 9, and 10 to have the longest contiguous subarray of 1s having length 9.
*/

//[1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1]

func LongestSubArrayWithOnesAfterReplacement(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}

	var windowStart int
	var maxLen int
	var maxOnesInWindow int

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		if arr[windowEnd] == 1 {
			maxOnesInWindow++
		}

		for windowEnd-windowStart+1-maxOnesInWindow > k {
			if arr[windowStart] == 1 {
				maxOnesInWindow--
			}
			windowStart++
		}

		maxLen = helpers.Max(maxLen, windowEnd-windowStart+1)
	}

	return maxLen
}
