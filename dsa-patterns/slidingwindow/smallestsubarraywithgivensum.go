package slidingwindow

import (
	"dsa-patterns/helpers"
	"math"
)

/* Given an array of positive numbers and a positive number ‘S’, find the length of the smallest contiguous subarray whose sum is greater than or equal to ‘S’. Return 0, if no such subarray exists.

Example 1:

Input: [2, 1, 5, 2, 3, 2], S=7
Output: 2
Explanation: The smallest subarray with a sum great than or equal to '7' is [5, 2].
Example 2:

Input: [2, 1, 5, 2, 8], S=7
Output: 1
Explanation: The smallest subarray with a sum greater than or equal to '7' is [8].
Example 3:

Input: [3, 4, 1, 1, 6], S=8
Output: 3
Explanation: Smallest subarrays with a sum greater than or equal to '8' are [3, 4, 1] or [1, 1, 6]. */

func SmallestSubArrayWithGivenSum(sum int, arr []int) []int {
	if len(arr) == 0 || sum == 0 {
		return []int{}
	}
	var windowStart, windowEnd int
	var smallWinStart int
	var smallWinLen int = math.MaxInt
	var currSum int

	for windowEnd = 0; windowEnd < len(arr); windowEnd++ {
		currSum += arr[windowEnd]

		for currSum >= sum {
			smallWinLen = helpers.Min(smallWinLen, windowEnd-windowStart+1)
			if windowEnd-windowStart+1 <= smallWinLen {
				smallWinStart = windowStart
			}
			currSum -= arr[windowStart]
			windowStart++
		}
	}
	return arr[smallWinStart : smallWinStart+smallWinLen]
}

//Pseudo Code

//[2, 1, 5, 2, 3, 2, 8],7
/*
Iterate through the list and gradually increase the window size until you find a sum that is equal to or greater than the given sum.
While sum is >= given sum, remove an element from the start of the window.
Have a variable to store the current smallest subarray i.e min(windowEnd-windowStart)
Repeat the above process until you find the smallest sub array with sum >= given sum
*/
