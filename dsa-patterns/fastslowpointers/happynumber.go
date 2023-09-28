package fastslowpointers

/*
Problem Statement #
Any number will be called a happy number if, after repeatedly replacing it with a number equal to the sum of the square of all of its digits, leads us to number ‘1’. All other (not-happy) numbers will never reach ‘1’. Instead, they will be stuck in a cycle of numbers which does not include ‘1’.

Example 1:

Input: 23
Output: true (23 is a happy number)
Explanations: Here are the steps to find out that 23 is a happy number:

2^2 + 3^2 = 4 + 9 = 13
1^2 + 3^  = 1 + 9 = 10
1^2 + 0^2 = 1 + 0 = 1


*/

func IsHappyNumber(num int) bool {
	var slow int = num
	var fast int = num
	for true {
		slow = getDigit(slow)
		fast = getDigit(getDigit(fast))
		if slow == fast {
			break
		}
	}
	return slow == 1
}

func getDigit(num int) int {
	var sum int
	for num > 0 {
		digit := num % 10
		num = num / 10
		sum += digit * digit
	}
	return sum
}
