package slidingwindow

/*
String Anagrams (hard) #
Given a string and a pattern, find all anagrams of the pattern in the given string.

Anagram is actually a Permutation of a string. For example, “abc” has the following six anagrams:

abc
acb
bac
bca
cab
cba
Write a function to return a list of starting indices of the anagrams of the pattern in the given string.

Example 1:

Input: String="ppqp", Pattern="pq"
Output: [1, 2]
Explanation: The two anagrams of the pattern in the given string are "pq" and "qp".
Example 2:

Input: String="abbcabc", Pattern="abc"
Output: [2, 3, 4]
Explanation: The three anagrams of the pattern in the given string are "bca", "cab", and "abc".
*/

func StringAnagrams(str, pattern string) []int {
	if str == "" || pattern == "" {
		return []int{}
	}
	var patternFreqMap = make(map[string]int)
	var resultIndices []int
	var matchedCnt int
	var windowStart int
	for _, c := range pattern {
		char := string(c)
		patternFreqMap[char]++
	}

	for windowEnd, c := range str {
		charAtWinEnd := string(c)

		if _, ok := patternFreqMap[charAtWinEnd]; ok {
			patternFreqMap[charAtWinEnd]--
			if patternFreqMap[charAtWinEnd] == 0 {
				matchedCnt++
			}

			if matchedCnt == len(patternFreqMap) {
				resultIndices = append(resultIndices, windowStart)
			}
		}

		if windowEnd-windowStart+1 >= len(pattern) {
			charAtWinStart := string(str[windowStart])
			if v, leftCharFound := patternFreqMap[charAtWinStart]; leftCharFound {
				if v == 0 {
					matchedCnt--
				}
				patternFreqMap[charAtWinStart]++
			}
			windowStart++
		}

	}

	return resultIndices
}

//abbc
//b

/*
Iterate over the pattern and store the characters of the pattern in a map with the count of each character as the value.
Iterate over the string and if the character exists in the pattern then decrement its count. Have a variable to track matched characters.
If the count is 0, then increment the matched variable.
Have a variable to store the max count of repeating character in current window. If the window has more characters than K, then gradually shrink the window from the front, and add the count back in the char frequency map.
If we find an anagram, store the windowstart in an array. Return all the starting indices of anagram.
*/
