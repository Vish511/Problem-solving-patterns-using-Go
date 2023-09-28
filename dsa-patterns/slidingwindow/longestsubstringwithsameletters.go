package slidingwindow

import (
	"dsa-patterns/helpers"
)

/*
Problem Statement #
Given a string with lowercase letters only, if you are allowed to replace no more than ‘k’ letters with any letter, find the length of the longest substring having the same letters after replacement.

Example 1:

Input: String="aabccbb", k=2
Output: 5
Explanation: Replace the two 'c' with 'b' to have a longest repeating substring "bbbbb".
Example 2:

Input: String="abbcb", k=1
Output: 4
Explanation: Replace the 'c' with 'b' to have a longest repeating substring "bbbb".
Example 3:

Input: String="abccde", k=1
Output: 3
Explanation: Replace the 'b' or 'd' with 'c' to have the longest repeating substring "ccc".
*/

func LongestSubStringWithSameLettersAfterReplacement(str string, k int) (string, int) {
	if str == "" {
		return "", 0
	}
	var maxLen int
	var maxRepeatingChar int
	var charFrequency = make(map[string]int)
	var windowStart int
	for windowEnd, c := range str {
		char := string(c)

		charFrequency[char]++

		maxRepeatingChar = helpers.Max(maxRepeatingChar, charFrequency[char])

		for windowEnd-windowStart+1-maxRepeatingChar > k {
			leftChar := string(str[windowStart])
			charFrequency[leftChar]--
			windowStart += 1
		}
		maxLen = helpers.Max(maxLen, windowEnd-windowStart+1)

	}
	return "", maxLen
}

func LongestSubStringWithSameLettersAfterReplacementBruteForce(str string, k int) (string, int) {
	if str == "" || k < 0 {
		return "", 0
	}
	var maxLen int
	var maxRepeatingChar string
	for windowEnd, c := range str {
		charAtWindowEnd := string(c)
		usedK := k

		for windowStart := windowEnd - 1; windowStart >= 0; windowStart-- {
			charAtWinStart := string(str[windowStart])
			if charAtWinStart != charAtWindowEnd {
				usedK--
			}
			if usedK >= 0 {
				if windowEnd-windowStart+1 > maxLen {
					maxLen = windowEnd - windowStart
					maxRepeatingChar = charAtWindowEnd
				}
				maxLen = helpers.Max(maxLen, windowEnd-windowStart+1)
			}
		}
	}
	return maxRepeatingChar, maxLen
}

/*

aabcbab 2

abcdefgh



"baabb" 3-4
"abb"



k is greater
k is lesser


a: 0,0-2
b:0,1



Declare a map of visited characters along with their last seen index and iterate through the characters of the string and check the difference between window end and windowstart store the max(windowend-windowstart,computedmax) in a variable.
If the new character visited is different than the current repeating character, decrement k and compute the length.
If the new character visited is different than the current repeating character and k == 0, then move the window start to the position of the character replaced with k
If new the character visited is not in the map


vars needed
longestLength
firstReplacedIdx
windowend,windowStart

*/
