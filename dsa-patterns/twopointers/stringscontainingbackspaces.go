package twopointers

/*
Comparing Strings containing Backspaces (medium) #
Given two strings containing backspaces (identified by the character ‘#’), check if the two strings are equal.

Example 1:

Input: str1="xy#z", str2="xzz#"
Output: true
Explanation: After applying backspaces the strings become "xz" and "xz" respectively.
Example 2:

Input: str1="xy#z", str2="xyz#"
Output: false
Explanation: After applying backspaces the strings become "xz" and "xy" respectively.
Example 3:

Input: str1="xp#", str2="xyz##"
Output: true
Explanation: After applying backspaces the strings become "x" and "x" respectively.
In "xyz##", the first '#' removes the character 'z' and the second '#' removes the character 'y'.
Example 4:

Input: str1="xywrrmp", str2="xywrrmu#p"
Output: true
Explanation: After applying backspaces the strings become "xywrrmp" and "xywrrmp" respectively.

*/

func CompareStringsContainingBackspaces(str1, str2 string) bool {
	var p1 = len(str1) - 1 //To iterate through str1 from the end
	var p2 = len(str2) - 1 //To iterate through str2 from the end
	var skipCntStr1 int    // Keep track of number of characters to be skipped in str1 because of backspace
	var skipCntStr2 int    // Keep track of number of characters to be skipped in str2 because of backspace

	for p1 >= 0 && p2 >= 0 {
		if string(str1[p1]) == "#" {
			skipCntStr1++
			p1--
		} else if skipCntStr1 > 0 {
			skipCntStr1--
			p1--
		}

		if string(str2[p2]) == "#" {
			skipCntStr2++
			p2--
		} else if skipCntStr2 > 0 {
			skipCntStr2--
			p2--
		}

		if skipCntStr1 == 0 && skipCntStr2 == 0 {
			if string(str1[p1]) != string(str2[p2]) { //if at any point, the characters of both strings don't match, then the strings do not match
				return false
			}
			p1--
			p2--
		}

		if p1 < 0 && p2 < 0 { //If both strings have reached their end, then its a match
			return true
		}

		if p1 < 0 || p2 < 0 { //If one of the strings have reached their end, then its not a match
			return false
		}
	}

	return false
}
