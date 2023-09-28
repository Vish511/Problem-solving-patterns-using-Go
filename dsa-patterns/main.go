package main

import (
	"dsa-patterns/intervals"
	"fmt"
)

func main() {
	fmt.Println("main")
	// fmt.Println(slidingwindow.MaxSubArraySum(3, []int{2, 1, 5, 1, 3, 2}))
	// fmt.Println(slidingwindow.SmallestSubArrayWithGivenSum(7, []int{2, 1, 5, 2, 8}))
	// fmt.Println(slidingwindow.LongestSubStringWithKDistinctChars(1, "araaci"))
	// fmt.Println(slidingwindow.FruitsIntoBaskets([]string{"A", "B", "C", "B", "B", "C"}))
	// fmt.Println(slidingwindow.NonRepeatingSubString("aabccbb"))
	// fmt.Println(slidingwindow.LongestSubStringWithSameLettersAfterReplacementBruteForce("abazbb", 2))
	// fmt.Println(slidingwindow.LongestSubStringWithSameLettersAfterReplacement("aabcfbb", 2))
	// fmt.Println(slidingwindow.LongestSubArrayWithOnesAfterReplacement([]int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}, 2))
	// fmt.Println(slidingwindow.PermutationsInAString("obbac", "abc"))
	// fmt.Println(slidingwindow.StringAnagrams("ppqp", "pq"))
	// fmt.Println(slidingwindow.SmallestWindowContainingSubstring("AYZABOBECODXBANC", "ABC"))
	// fmt.Println(slidingwindow.WordsConcatenation("catratcatcatcatfoxfoxcat", []string{"cat", "fox"}))
	// fmt.Println(twopointers.PairWithTargetSum([]int{2, 5, 9, 11}, 11))
	// fmt.Println(twopointers.RemoveDuplicates([]int{1, 2, 2, 2, 2, 2, 3, 3, 4444}))
	// fmt.Println(twopointers.RemoveAllInstancesOfKey([]int{3, 2, 3, 6, 3, 10, 9, 3}, 3))
	// fmt.Println(twopointers.SquaringASortedArray([]int{-3, -1, 0, 0}))
	// fmt.Println(twopointers.TripletSumToZero([]int{-5, 2, -1, -2, 3}))
	// fmt.Println(twopointers.PairWithSumCloseToTarget([]int{1, 1, 1, 2, 3, 4, 4, 5}, 10))
	// fmt.Println(twopointers.TripletSumCloseToTarget([]int{1, 0, 1, 1}, 2))
	// fmt.Println(twopointers.TripletSumLessThanTarget([]int{-1, 4, 2, 1, 3}, 5))
	// fmt.Println(twopointers.DutchNationalFlag([]int{2, 2, 0, 1, 2, 0}))
	// fmt.Println(twopointers.QuadrupleSumToTarget([]int{4, 1, 2, -1, 1, -3}, 1))
	// fmt.Println(twopointers.CompareStringsContainingBackspaces("xy#z", "xzz#"))
	// fmt.Println(twopointers.MinWindowSort([]int{1, 3, 2, 0, -1, 7, 10}))
	// fmt.Println(twopointers.NewSinglyLinkedList(23))
	// sll := fastslowpointers.NewSinglyLinkedList(2)

	// sll.Insert(4)
	// sll.Insert(6)
	// sll.Insert(4)
	// sll.Insert(2)
	// // sll.Insert(12)
	// // sll.Insert(1)

	// // sll.Tail.Next = sll.Head.Next
	// //HasCycle
	// fmt.Println(sll.HasCycle())
	// //Determine Length Of Cycle
	// fmt.Println(sll.LengthOfCycle(sll.Head))
	// //Find first node of cycle
	// firstNode := sll.FindFirstNodeOfCycle()
	// fmt.Println(firstNode)
	// // fmt.Println(fastslowpointers.IsHappyNumber(29))
	// //Find middle node of list
	// middleNode := sll.FindMiddleNode()
	// fmt.Println(middleNode)

	// //Reverse list
	// sll.Reverse(sll.Head)
	// fmt.Println("Traverse after reverse")
	// sll.Traverse()

	// //Check if list is a palindrome
	// fmt.Println("is Palindrome", sll.IsPalindrome())

	// //Rearrange list
	// sll.Rearrange()
	// fmt.Println(*sll.Head.Value, *sll.Tail.Value)

	// fmt.Println(fastslowpointers.CircularArrayLoopExists([]int{2, 1, -1, -2}))
	// //[2,2,2,-1]
	// // fmt.Println(fastslowpointers.MoveForward(3, -1, 4))

	// fmt.Println(helpers.MergeHelper([]int{1, 3, 5, 7}, []int{2, 4, 5, 7, 9}))
	// fmt.Println(helpers.MergeSort([]int{4, 3, 2, 1, 5, 6}))
	// fmt.Println(mergeintervals.MergeIntervals([]mergeintervals.Interval{mergeintervals.NewInterval(1, 4), mergeintervals.NewInterval(2, 6), mergeintervals.NewInterval(7, 9), mergeintervals.NewInterval(10, 12)}))
	// intervalList := intervals.NewIntervalList([][]int{{1, 3}, {5, 7}, {8, 12}})
	// fmt.Println(intervalList.InsertInterval(intervals.NewInterval(4, 6)))

	intvl1 := intervals.NewIntervalList([][]int{{1, 3}, {5, 7}, {9, 12}})
	intvl2 := intervals.NewIntervalList([][]int{{5, 10}})
	fmt.Println(intervals.IntersectIntervals(intvl1, intvl2))

	conflictingIntervals := intervals.NewIntervalList([][]int{{4, 5}, {2, 3}, {3, 6}})

	fmt.Println(conflictingIntervals)
	intvlSort := intervals.NewIntervalList([][]int{{3, 27},
		{5, 21},
		{2, 8},
		{4, 16},
		{15, 21},
		{10, 20},
		{17, 29},
		{23, 25}})
	fmt.Println(intervals.MinMeetingRooms(intvlSort))

	var cpuJobList = intervals.NewCPUJobList([][]int{{1, 4, 3}, {2, 5, 4}, {7, 9, 6}})
	var cpuJobList2 = intervals.NewCPUJobList([][]int{{6, 7, 10}, {2, 4, 11}, {8, 12, 15}, {13, 15, 10}, {16, 20, 5}, {18, 23, 10}, {21, 29, 15}, {27, 31, 5}})
	var cpuJobList3 = intervals.NewCPUJobList([][]int{{1, 4, 2}, {2, 4, 1}, {3, 6, 5}})

	fmt.Println(cpuJobList.MaximumCPULoad())
	fmt.Println(cpuJobList2.MaximumCPULoad())
	fmt.Println(cpuJobList3.MaximumCPULoad())
	// intervals := mergeintervals.Intervals{}
}
