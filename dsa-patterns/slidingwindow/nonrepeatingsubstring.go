package slidingwindow

/* Problem Statement #
Given a string, find the length of the longest substring which has no repeating characters.

Example 1:

Input: String="aabccbb"
Output: 3
Explanation: The longest substring without any repeating characters is "abc".
Example 2:

Input: String="abbbb"
Output: 2
Explanation: The longest substring without any repeating characters is "ab".
Example 3:

Input: String="abcbde"
Output: 3
Explanation: Longest substrings without any repeating characters are "abc" & "cde". */

func NonRepeatingSubString(str string) int {
	if str == "" {
		return 0
	}
	var windowStart int
	var visitedChars = make(map[string]int, len(str))
	var resultLen int
	// var resultWinStartIdx int

	for windowEnd, r := range str {
		char := string(r)

		if _, ok := visitedChars[char]; !ok {
			visitedChars[char] = windowEnd
			if windowEnd-windowStart+1 > resultLen {
				// resultWinStartIdx = windowStart
				resultLen = windowEnd - windowStart + 1
			}
		} else {
			windowStart = visitedChars[char] + 1
			visitedChars[char] = windowEnd
		}

	}
	return resultLen

}

/*
Iterate through the string and keep track of visited chars in a map.
If the character hasn't already been visited, add it to the visited chars and keep track of max length of non repeating characters in a variable.
If a character has been visited already, move the window start to index of repeating character previously visited + 1. Delete the repeating character from the visited chars map
{a:0,b:1,c:2}
*/
