package slidingwindow

import (
	"dsa-patterns/helpers"
	"math"
)

/*
Smallest Window containing Substring (hard) #
Given a string and a pattern, find the smallest substring in the given string which has all the characters of the given pattern.

Example 1:

Input: String="aabdec", Pattern="abc"
Output: "abdec"
Explanation: The smallest substring having all characters of the pattern is "abdec"
Example 2:

Input: String="abdabca", Pattern="abc"
Output: "abc"
Explanation: The smallest substring having all characters of the pattern is "abc".
Example 3:

Input: String="adcad", Pattern="abc"
Output: ""
Explanation: No substring in the given string has all characters of the pattern.
*/

func SmallestWindowContainingSubstring(str, pattern string) string {
	var resultStartIdx int
	var windowStart int
	var minLen int = math.MaxInt
	var matchedCnt int
	var charFrequencyMap = make(map[string]int)
	for _, c := range pattern {
		char := string(c)
		charFrequencyMap[char]++
	}

	for windowEnd, c := range str {
		charAtWinEnd := string(c)
		if _, rightCharFound := charFrequencyMap[charAtWinEnd]; rightCharFound {
			charFrequencyMap[charAtWinEnd]--
			if charFrequencyMap[charAtWinEnd] == 0 {
				matchedCnt++
			}

			for matchedCnt == len(charFrequencyMap) {
				minLen = helpers.Min(minLen, windowEnd-windowStart+1)
				resultStartIdx = windowStart
				charAtWinStart := string(str[windowStart])
				if _, leftCharFound := charFrequencyMap[charAtWinStart]; leftCharFound {
					charFrequencyMap[charAtWinStart]++
					if charFrequencyMap[charAtWinStart] > 0 {
						matchedCnt--
					}
				}
				windowStart++
			}
		}
	}
	if minLen == math.MaxInt {
		return ""
	}
	return str[resultStartIdx : resultStartIdx+minLen]
}

/*
aabdec abc
abdabca abc

*/
