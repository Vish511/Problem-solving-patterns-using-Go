package slidingwindow

import "dsa-patterns/helpers"

//Problem Statement
/*
Given a string, find the length of the longest substring in it with no more than K distinct characters.

Example 1:

Input: String="araaci", K=2
Output: 4
Explanation: The longest substring with no more than '2' distinct characters is "araa".
Example 2:

Input: String="araaci", K=1
Output: 2
Explanation: The longest substring with no more than '1' distinct characters is "aa".
Example 3:

Input: String="cbbebi", K=3
Output: 5
Explanation: The longest substrings with no more than '3' distinct characters are "cbbeb" & "bbebi".
*/

func LongestSubStringWithKDistinctChars(k int, str string) (string, int) {
	if str == "" || k == 0 {
		return "", 0
	}
	var windowStart int
	var maxSubStrLen int
	var targetWinStartIdx int
	var distinctChars = make(map[string]int)

	for windowEnd, char := range str {
		charAtWinEnd := string(char)
		distinctChars[charAtWinEnd]++

		if len(distinctChars) == k {
			if windowEnd-windowStart+1 > maxSubStrLen {
				targetWinStartIdx = windowStart
			}
			maxSubStrLen = helpers.Max(maxSubStrLen, windowEnd-windowStart+1)
		}

		for len(distinctChars) > k {
			charAtWinStart := string(str[windowStart])
			distinctChars[charAtWinStart]--
			if distinctChars[charAtWinStart] == 0 {
				delete(distinctChars, charAtWinStart)
			}
			windowStart++
		}
	}

	return str[targetWinStartIdx : targetWinStartIdx+maxSubStrLen], maxSubStrLen
}

/* Pseudo code

Iterate through the array by starting a window with one element and gradually adding more elements.
Maintain a variable to keep track of distinct characters visited so far.
Check if distinct characters visited so far is equal to k, if it is, then remove one element from the start of the window and add an element to the end of the window.
Declare a variable to keep track of the longest sub string length along with the start index of the longest sub string.
Update this variable on every iteration.
Return the first longest substr visited along with the length



*/
