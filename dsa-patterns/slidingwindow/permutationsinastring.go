package slidingwindow

/*
Permutation in a String (hard) #
Given a string and a pattern, find out if the string contains any permutation of the pattern.

Permutation is defined as the re-arranging of the characters of the string. For example, “abc” has the following six permutations:

abc
acb
bac
bca
cab
cba
If a string has ‘n’ distinct characters it will have
n!
n! permutations.

Example 1:

Input: String="oidbcaf", Pattern="abc"
Output: true
Explanation: The string contains "bca" which is a permutation of the given pattern.
Example 2:

Input: String="odicf", Pattern="dc"
Output: false
Explanation: No permutation of the pattern is present in the given string as a substring.
Example 3:

Input: String="bcdxabcdy", Pattern="bcdyabcdx"
Output: true
Explanation: Both the string and the pattern are a permutation of each other.
Example 4:

Input: String="aaacb", Pattern="abc"
Output: true
Explanation: The string contains "acb" which is a permutation of the given pattern.
*/

func PermutationsInAString(str, pattern string) bool {
	if str == "" || pattern == "" {
		return false
	}
	var patternCharMap = make(map[string]int)
	var windowStart int
	var matched int
	for _, c := range pattern {
		char := string(c)
		patternCharMap[char]++
	}

	for windowEnd, c := range str {
		charAtWindowEnd := string(c)

		_, ok := patternCharMap[charAtWindowEnd]
		if ok {
			patternCharMap[charAtWindowEnd]--
			if patternCharMap[charAtWindowEnd] == 0 {
				matched++
			}
			if matched == len(patternCharMap) {
				return true
			}
		}

		if windowEnd >= len(pattern)-1 {
			charAtWindowStart := string(str[windowStart])
			windowStart++
			if v, leftCharFound := patternCharMap[charAtWindowStart]; leftCharFound {
				if v == 0 {
					matched--
				}
				patternCharMap[charAtWindowStart]++

			}
		}
	}

	return false
}

/*
ododicf {d:1,c:1}
o


ocbbbac   abc

*/

/*
Iterate through the pattern and store the characters along with their frequency in a map.
Iterate through the string and if a character is found which is part of the pattern, reduce the count.If we encouter a character with a frequency higher than the one in the pattern, then shrink the window.
If we encounter a character which is not part of the pattern, shrink the start of the window in a loop and add the count of the characters going out of the window
Add a check to see if the window length equals the length of the pattern and if it does then return true, else false
*/
