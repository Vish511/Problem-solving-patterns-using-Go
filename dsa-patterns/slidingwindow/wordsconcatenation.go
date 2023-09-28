package slidingwindow

/*
Words Concatenation (hard) #
Given a string and a list of words, find all the starting indices of substrings in the given string that are a concatenation of all the given words exactly once without any overlapping of words. It is given that all words are of the same length.

Example 1:

Input: String="catfoxcat", Words=["cat", "fox"]
Output: [0, 3]
Explanation: The two substring containing both the words are "catfox" & "foxcat".
Example 2:

Input: String="catcatfoxfox", Words=["cat", "fox"]
Output: [3]
Explanation: The only substring containing both the words is "catfox".
*/

func WordsConcatenation(str string, words []string) []int {
	if str == "" || len(words) == 0 {
		return []int{}
	}
	var windowStart int
	var matchedCnt int
	var result []int
	var wordFreqMap = make(map[string]int)
	var wordLen int = len(words[0])
	var concatWordLen int = wordLen * len(words)

	for i := 0; i < len(words); i++ {
		wordFreqMap[words[i]]++
		//Check to ensure length is the same for all words
		if len(words[i]) != wordLen {
			return []int{}
		}
	}

	for windowEnd, _ := range str {
		if windowEnd >= wordLen-1 {
			wordToBeMatched := str[windowStart : windowEnd+1]
			if _, matchFound := wordFreqMap[wordToBeMatched]; matchFound {
				wordFreqMap[wordToBeMatched]--
				if wordFreqMap[wordToBeMatched] == 0 {
					matchedCnt++
				}
			}
			//If window size exceeds the total length of concatenated words, begin to shrink the window.
			if windowEnd >= concatWordLen-1 {
				if matchedCnt == len(wordFreqMap) {
					result = append(result, windowEnd-concatWordLen+1)
				}
				wordToBeRemoved := str[windowEnd-concatWordLen+1 : windowEnd-concatWordLen+wordLen+1]

				if _, found := wordFreqMap[wordToBeRemoved]; found {
					if wordFreqMap[wordToBeRemoved] == 0 {
						matchedCnt--
					}
					wordFreqMap[wordToBeRemoved]++
				}
			}
			windowStart++
		}
	}
	return result
}
