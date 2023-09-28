package intervals

import (
	"dsa-patterns/helpers"
	"sort"
)

/*
Problem Statement #
Given a list of intervals, merge all the overlapping intervals to produce a list that has only mutually exclusive intervals.

Example 1:

Intervals: [[1,4], [2,5], [7,9]]
Output: [[1,5], [7,9]]
Explanation: Since the first two intervals [1,4] and [2,5] overlap, we merged them into
one [1,5].
svg viewer

Example 2:
Intervals: [[6,7], [2,4], [5,9]]
Output: [[2,4], [5,9]]
Explanation: Since the intervals [6,7] and [5,9] overlap, we merged them into one [5,9].

Example 3:
Intervals: [[1,4], [2,6], [3,5]]
Output: [[1,6]]
Explanation: Since all the given intervals overlap, we merged them into one.
*/

type Interval struct {
	Start int
	End   int
}

type Intervals []Interval

func NewIntervalList(list [][]int) Intervals {
	var intervals = make(Intervals, 0, len(list))
	for _, v := range list {
		intervals = append(intervals, Interval{Start: v[0], End: v[1]})
	}
	return intervals
}

func NewInterval(start, end int) Interval {
	return Interval{Start: start, End: end}
}

func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	var i int = 1
	var result []Interval
	var max = intervals[0].End
	var min = intervals[0].Start

	for i < len(intervals) {
		if intervals[i].Start <= max {
			max = helpers.Max(max, intervals[i].End)
		} else {
			result = append(result, NewInterval(min, max))
			min = intervals[i].Start
			max = intervals[i].End
		}
		i++
	}

	result = append(result, NewInterval(min, max))
	return result
}

func (intvls Intervals) Sort() {
	sort.Slice(intvls, func(a, b int) bool {
		if intvls[a].Start == intvls[b].Start {
			return intvls[a].End < intvls[b].End
		}
		return intvls[a].Start < intvls[b].Start
	})
}

func (intervals Intervals) SortByEndTime() {
	sort.Slice(intervals, func(a, b int) bool {
		if intervals[a].End == intervals[b].End {
			return intervals[a].Start < intervals[b].Start
		}
		return intervals[a].End < intervals[b].End
	})
}

/*
Problem Statement #
Given a list of non-overlapping intervals sorted by their start time, insert a given interval at the correct position and merge all necessary intervals to produce a list that has only mutually exclusive intervals.

Example 1:

Input: Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,6]
Output: [[1,3], [4,7], [8,12]]
Explanation: After insertion, since [4,6] overlaps with [5,7], we merged them into one [4,7].
Example 2:

Input: Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,10]
Output: [[1,3], [4,12]]
Explanation: After insertion, since [4,10] overlaps with [5,7] & [8,12], we merged them into [4,12].
Example 3:

Input: Intervals=[[2,3],[5,7]], New Interval=[1,4]
Output: [[1,4], [5,7]]
Explanation: After insertion, since [1,4] overlaps with [2,3], we merged them into one [1,4].
*/

// Input: Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,7]
func (itvls Intervals) InsertInterval(interval Interval) Intervals {
	var min = interval.Start
	var max = interval.End
	var result Intervals
	var currInterval = Interval{Start: itvls[0].Start, End: itvls[0].End}
	var i int
	for i < len(itvls) {
		if itvls[i].End >= min && itvls[i].Start <= max {
			min = helpers.Min(min, itvls[i].Start)
			max = helpers.Max(max, itvls[i].End)
			currInterval = Interval{min, max}
		} else {
			result = append(result, currInterval)
			currInterval = itvls[i]
		}
		i++
	}
	result = append(result, currInterval)

	return result
}

/*
Problem Statement #
Given two lists of intervals, find the intersection of these two lists. Each list consists of disjoint intervals sorted on their start time.

Example 1:

Input: arr1=[[1, 3], [5, 6], [7, 9]], arr2=[[2, 3], [5, 7]]
Output: [2, 3], [5, 6], [7, 7]
Explanation: The output list contains the common intervals between the two lists.
Example 2:

Input: arr1=[[1, 3], [5, 7], [9, 12]], arr2=[[5, 10]]
Output: [5, 7], [9, 10]
Explanation: The output list contains the common intervals between the two lists.
*/

// Input: arr1=[[1, 3], [5, 10], [9, 12]], arr2=[[5, 6],[7,8]]
//[[5,6],[7,8]]

// Input: arr1=[[1, 3], [5, 8], [9, 12]], arr2=[[4, 10]]
//[[5,8],[9,10]]

// Input: arr1=[[1, 2], [3, 4], [5, 6], [7,8]], arr2=[[7, 9]]
func IntersectIntervals(intvl1, intvl2 Intervals) Intervals {
	var result Intervals
	var i, j int
	var min, max int
	for i < len(intvl1) && j < len(intvl2) {
		if intvl1[i].End >= intvl2[j].Start && intvl1[i].Start <= intvl2[j].End { //If they intersect
			min = helpers.Max(intvl1[i].Start, intvl2[j].Start)
			max = helpers.Min(intvl1[i].End, intvl2[j].End)
			result = append(result, NewInterval(min, max))
		}
		if intvl1[i].End == intvl2[j].End {
			i++
			j++
		} else if intvl1[i].End < intvl2[j].End {
			i++
		} else {
			j++
		}
	}

	return result
}

/*
Problem Statement #

Given an array of intervals representing ‘N’ appointments, find out if a person can attend all the appointments.

Example 1:
Appointments: [[1,4], [2,5], [7,9]]
Output: false
Explanation: Since [1,4] and [2,5] overlap, a person cannot attend both of these appointments.

Example 2:
Appointments: [[6,7], [2,4], [8,12]]
Output: true
Explanation: None of the appointments overlap, therefore a person can attend all of them.

Example 3:
Appointments: [[4,5], [2,3], [3,6]]
Output: false
Explanation: Since [4,5] and [3,6] overlap, a person cannot attend both of these appointments.
*/
func AreAppointmentsConflicting(intervals Intervals) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	var max = intervals[0].End
	var i int = 1

	for i < len(intervals) {
		if intervals[i].Start < max {
			return true
		} else {
			max = intervals[i].End
		}
		i++
	}

	return false
}

/*
Minimum Meeting Rooms (hard) #
Given a list of intervals representing the start and end time of ‘N’ meetings, find the minimum number of rooms required to hold all the meetings.

Example 1:

Meetings: [[1,4], [2,5], [7,9]]
Output: 2
Explanation: Since [1,4] and [2,5] overlap, we need two rooms to hold these two meetings. [7,9] can
occur in any of the two rooms later.
Example 2:

Meetings: [[6,7], [2,4], [8,12]]
Output: 1
Explanation: None of the meetings overlap, therefore we only need one room to hold all meetings.
Example 3:

Meetings: [[1,4], [2,3], [3,6]]
Output:2
Explanation: Since [1,4] overlaps with the other two meetings [2,3] and [3,6], we need two rooms to
hold all the meetings.

Example 4:

Meetings: [[4,5], [2,3], [2,4], [3,5]]
Output: 2
Explanation: We will need one room for [2,3] and [3,5], and another room for [2,4] and [4,5].

*/

func MinMeetingRooms(intervals Intervals) int {
	if len(intervals) < 1 {
		return 0
	}
	intervals.Sort()
	var maxRoomCnt int
	var priorityQ = NewPriorityQueue()
	for i := 0; i < len(intervals); i++ {
		anyMeetingExists, meeting := priorityQ.Peek()
		for anyMeetingExists && intervals[i].Start >= meeting.End { //checking and removing other meetings which don't conflict with the current meeting. At any point, priority queue will have the list of all active meetings which overlap each other.
			priorityQ.ExtractMin()
			anyMeetingExists, meeting = priorityQ.Peek()
		}
		priorityQ.Insert(intervals[i], intervals[i].End)            //Setting the end time as priority
		maxRoomCnt = helpers.Max(maxRoomCnt, len(priorityQ.values)) //Room count is the max of all concurrent meetings
	}

	return maxRoomCnt
}

/* Below is an alternate approach without using priority queue which doesn't result in best fit as this approach looks for the meeting with the smallest end time before allocating a new meeting */
/* It won't work for this input - [[2, 8], [4 ,16], [10 ,20], [5 ,21], [15 ,21], [23 ,25], [3 ,27], [17, 29]] */
/* For the above scenario, the optimal output is 5 but the below method returns 6 */
/* Below is the explanation of the allocation strategy which causes 6 rooms to return instead of 5 :
[[{2,8},{10,20}],[{4,16},{23,25}],[{5,21},{}],[{15,21},{}],[{3,27},{}],[{{17,29}},{}]] */
// -----Room1----   -----Room2----  ---Room3--- ---Room4---  ---Room5---  ----Room6----
func MinMeetingRoomsAlternate(intervals Intervals) int {
	if len(intervals) < 1 {
		return 0
	}
	intervals.SortByEndTime()
	var vacantIdx int
	var roomCnt int = 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < intervals[vacantIdx].End {
			roomCnt++
		} else {
			vacantIdx++
		}
	}

	return roomCnt
}
