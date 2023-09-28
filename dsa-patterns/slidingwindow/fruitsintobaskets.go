package slidingwindow

import "dsa-patterns/helpers"

/* Problem Statement #
Given an array of characters where each character represents a fruit tree, you are given two baskets and your goal is to put maximum number of fruits in each basket. The only restriction is that each basket can have only one type of fruit.

You can start with any tree, but once you have started you can’t skip a tree. You will pick one fruit from each tree until you cannot, i.e., you will stop when you have to pick from a third fruit type.

Write a function to return the maximum number of fruits in both the baskets.

Example 1:

Input: Fruit=['A', 'B', 'C', 'A', 'C']
Output: 3
Explanation: We can put 2 'C' in one basket and one 'A' in the other from the subarray ['C', 'A', 'C']
Example 2:

Input: Fruit=['A', 'B', 'C', 'B', 'B', 'C']
Output: 5
Explanation: We can put 3 'B' in one basket and two 'C' in the other basket.
This can be done if we start with the second letter: ['B', 'C', 'B', 'B', 'C']
*/

func FruitsIntoBaskets(fruits []string) ([]string, int) {
	if len(fruits) == 0 {
		return []string{}, 0
	}
	var basketLimit int = 2 //Max number of baskets that can hold fruits
	var filledBaskets = make(map[string]int, len(fruits))
	var windowStart int
	var maxFruitsLen int
	var resultWinStartIdx int
	for windowEnd, fruitTree := range fruits {
		fruitTree = string(fruitTree)
		filledBaskets[fruitTree]++

		if len(filledBaskets) == basketLimit {
			if windowEnd-windowStart+1 > maxFruitsLen {
				resultWinStartIdx = windowStart
			}
			maxFruitsLen = helpers.Max(maxFruitsLen, windowEnd-windowStart+1)
		}

		for len(filledBaskets) > basketLimit {
			fruitAtWinStart := string(fruits[windowStart])
			filledBaskets[fruitAtWinStart]--

			if filledBaskets[fruitAtWinStart] == 0 {
				delete(filledBaskets, fruitAtWinStart)
			}
			windowStart++
		}
	}

	return fruits[resultWinStartIdx : resultWinStartIdx+maxFruitsLen], maxFruitsLen
}

/*
 Iterate through the fruit trees and gradually add fruits to the baskets.
 Once the basket limits are full, calculate the max number of fruits added thus far and store it in a variable.
 When the basket limit is greater than allowed limit, keep removing fruit trees which we visited first until there is space for adding more fruit trees.
 Once there is space in the basket, keep adding new fruit trees in the basket and repeat the above process.
 Return the max number of fruit trees along with their details.
*/
