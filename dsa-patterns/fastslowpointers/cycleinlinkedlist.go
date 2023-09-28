package fastslowpointers

import (
	"fmt"
)

/*
Problem Statement #
Given the head of a Singly LinkedList, write a function to determine if the LinkedList has a cycle in it or not.

*/

type Node struct {
	Next  *Node
	Value *int
}

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

func NewSinglyLinkedList(val int) *SinglyLinkedList {
	newNode := Node{Value: &val}
	return &SinglyLinkedList{Head: &newNode, Tail: &newNode, Len: 1}
}

func (sll *SinglyLinkedList) Insert(val int) {
	newNode := Node{Value: &val}
	currTail := sll.Tail
	currTail.Next = &newNode
	sll.Tail = &newNode
	sll.Len++
}

func (sll *SinglyLinkedList) Traverse(head ...*Node) {
	var curr *Node
	curr = sll.Head
	if len(head) > 0 {
		curr = head[0]
	}
	fmt.Print(*curr.Value)
	for curr.Next != nil {
		curr = curr.Next
		fmt.Printf("-> %d", *curr.Value)
	}
}

func (sll *SinglyLinkedList) HasCycle() bool {
	cycleExists, _ := sll.checkForCycle()
	if cycleExists {
		return true
	}
	return false
}

// checkForCycle will return if the list has a cycle and pointer to the node where slow & fast pointers meet.
func (sll *SinglyLinkedList) checkForCycle() (bool, *Node) {
	var slow = sll.Head
	var fast = sll.Head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true, slow
		}
	}
	return false, nil
}

/*
Given the head of a LinkedList with a cycle, find the length of the cycle.
1>2>3>4
*/

func (sll *SinglyLinkedList) LengthOfCycle(head *Node) int {
	cycleExists, slowPtr := sll.checkForCycle()
	if cycleExists {
		return sll.calculateCycleLen(slowPtr)
	}
	return -1
}

func (sll *SinglyLinkedList) calculateCycleLen(slow *Node) int {
	var current = slow
	var length int
	for current != nil && current.Next != nil {
		current = current.Next
		length++
		if current == slow {
			break
		}
	}
	return length
}

/* Problem Statement #
Given the head of a Singly LinkedList that contains a cycle, write a function to find the starting node of the cycle. */

// 1>2>3>4>5
// ....|---|
// ......^

// 1>2>3>4>5
// ..|-----|
// ........^
func (sll *SinglyLinkedList) FindFirstNodeOfCycle() *Node {
	lenOfCycle := sll.LengthOfCycle(sll.Head)
	if lenOfCycle != -1 {
		ptr1 := sll.Head
		ptr2 := ptr1
		for lenOfCycle > 0 {
			ptr2 = ptr2.Next
			lenOfCycle -= 1
		}

		for ptr1 != ptr2 {
			ptr1 = ptr1.Next
			ptr2 = ptr2.Next
		}
		return ptr1

	}
	return nil
}

/*
Problem Statement #
Given the head of a Singly LinkedList, write a method to return the middle node of the LinkedList.

If the total number of nodes in the LinkedList is even, return the second middle node.

Example 1:

Input: 1 -> 2 -> 3 -> 4 -> 5 -> null
Output: 3
Example 2:

Input: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null
Output: 4
Example 3:

Input: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> null
Output: 4

Input: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> null
Output: 5
*/

func (sll *SinglyLinkedList) FindMiddleNode() *Node {
	var slow = sll.Head
	var fast = sll.Head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func (sll *SinglyLinkedList) copy(head *Node) *SinglyLinkedList {
	var newSll = NewSinglyLinkedList(*head.Value)
	var curr = head.Next
	for curr != nil {
		newSll.Insert(*curr.Value)
		curr = curr.Next
	}
	return newSll
}

/* Palindrome LinkedList (medium) #
Given the head of a Singly LinkedList, write a method to check if the LinkedList is a palindrome or not.

Your algorithm should use constant space and the input LinkedList should be in the original form once the algorithm is finished. The algorithm should have
�
(
�
)
O(N) time complexity where ‘N’ is the number of nodes in the LinkedList.

Example 1:

Input: 2 -> 4 -> 6 -> 4 -> 2 -> null
Output: true


Example 2:

Input: 2 -> 4 -> 6 -> 4 -> 2 -> 2 -> null
Output: false */

//1>2>3>2>1>1>2>3>2>1

//1>2>3>2>1
//1>2>3

//1>2>3>3>2>1
//3>2>1 -- 1>2>3

func (sll *SinglyLinkedList) IsPalindrome() bool {
	if sll.Head == nil || sll.Tail == nil {
		return false
	}
	var isPalindrome bool
	middleNode := sll.FindMiddleNode()
	copyOfSecondHalf := sll.copy(middleNode)
	copyOfSecondHalf.Reverse(copyOfSecondHalf.Head)
	var headSecondHalf = copyOfSecondHalf.Head
	var headFirstHalf = sll.Head
	for headSecondHalf != nil && headFirstHalf != nil {
		if *headSecondHalf.Value != *headFirstHalf.Value {
			break
		}
		headSecondHalf = headSecondHalf.Next
		headFirstHalf = headFirstHalf.Next
	}

	if headFirstHalf == nil || headSecondHalf == nil {
		isPalindrome = true
	}

	return isPalindrome
}

func (sll *SinglyLinkedList) Reverse(head *Node) *Node {
	curr := head
	head, sll.Tail = sll.Tail, head
	if sll.Head == sll.Tail {
		sll.Head = head
	}
	var next *Node
	var prev *Node

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

/*
Rearrange a LinkedList (medium) #
Given the head of a Singly LinkedList, write a method to modify the LinkedList such that the nodes from the second half of the LinkedList are inserted alternately to the nodes from the first half in reverse order. So if the LinkedList has nodes 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null, your method should return 1 -> 6 -> 2 -> 5 -> 3 -> 4 -> null.

Your algorithm should not use any extra space and the input LinkedList should be modified in-place.

Example 1:

Input: 2 -> 4 -> 6 -> 8 -> 10 -> 12 -> null
Output: 2 -> 12 -> 4 -> 10 -> 6 -> 8 -> null
Example 2:

Input: 2 -> 4 -> 6 -> 8 -> 10 -> null
Output: 2 -> 10 -> 4 -> 8 -> 6 -> null

*/

// figure out how to determine tail
func (sll *SinglyLinkedList) Rearrange() {
	middleNode := sll.FindMiddleNode()

	headSecondHalf := sll.Reverse(middleNode)
	var headFirstHalf = sll.Head
	var oldNext *Node
	var oldSecondHalfNext *Node
	for headFirstHalf != nil && headSecondHalf != nil && headFirstHalf.Next != headSecondHalf {
		oldNext = headFirstHalf.Next
		oldSecondHalfNext = headSecondHalf.Next
		headFirstHalf.Next = headSecondHalf
		headSecondHalf.Next = oldNext

		headFirstHalf = oldNext
		headSecondHalf = oldSecondHalfNext
	}
	sll.Traverse()
}

// 1>2>3>4>5
func (sll *SinglyLinkedList) setTail(tail *Node) *Node {
	if tail != nil {
		sll.Tail = tail
		return sll.Tail
	}
	curr := sll.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	sll.Tail = curr
	return sll.Tail
}
