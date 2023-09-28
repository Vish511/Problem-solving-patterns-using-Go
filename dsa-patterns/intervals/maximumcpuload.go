package intervals

import (
	"dsa-patterns/helpers"
	"fmt"
	"sort"
)

/* Maximum CPU Load (hard) #
We are given a list of Jobs. Each job has a Start time, an End time, and a CPU load when it is running. Our goal is to find the maximum CPU load at any time if all the jobs are running on the same machine.

Example 1:

Jobs: [[1,4,3], [2,5,4], [7,9,6]]
Output: 7
Explanation: Since [1,4,3] and [2,5,4] overlap, their maximum CPU load (3+4=7) will be when both the
jobs are running at the same time i.e., during the time interval (2,4).
Example 2:

Jobs: [[6,7,10], [2,4,11], [8,12,15]]
Output: 15
Explanation: None of the jobs overlap, therefore we will take the maximum load of any job which is 15.
Example 3:

Jobs: [[1,4,2], [2,4,1], [3,6,5]]
Output: 8
Explanation: Maximum CPU load will be 8 as all jobs overlap during the time interval [3,4].  */

type CPUJob struct {
	Interval
	Load int
}

type CPUJobList []CPUJob

func NewCPUJobList(jobs [][]int) CPUJobList {
	var cpuJobList []CPUJob
	for i := 0; i < len(jobs); i++ {
		cpuJobList = append(cpuJobList, CPUJob{Interval: Interval{Start: jobs[i][0], End: jobs[i][1]}, Load: jobs[i][2]})
	}
	return cpuJobList
}

func (jobs CPUJobList) MaximumCPULoad() int {
	if len(jobs) < 1 {
		return 0
	}
	sort.SliceStable(jobs, func(a, b int) bool {
		return jobs[a].Start < jobs[b].Start
	})
	fmt.Println(jobs)
	var jobLoadMap = jobs.createJobLoadMap() //Created a map to lookup the load for a given interval. This has been used inorder to avoid reimplementing priority queue for the 'CPUJob' type. Also, this won't work for duplicate jobs (jobs having same start & end time)
	var priorityQ = NewPriorityQueue()
	var maxCPULoad int
	var currCPULoad int
	for i := 0; i < len(jobs); i++ {
		prevJobFound, prevJob := priorityQ.Peek()
		for prevJobFound && prevJob.End <= jobs[i].Start {
			currCPULoad -= jobLoadMap[prevJob]
			priorityQ.ExtractMin()
			prevJobFound, prevJob = priorityQ.Peek()
		}
		priorityQ.Insert(jobs[i].Interval, jobs[i].End)
		currCPULoad += jobs[i].Load
		fmt.Println(currCPULoad, priorityQ.values)
		maxCPULoad = helpers.Max(maxCPULoad, currCPULoad)
	}
	return maxCPULoad
}

func (jobs CPUJobList) createJobLoadMap() map[Interval]int {
	var jobLoadMap = make(map[Interval]int, len(jobs))
	for i := 0; i < len(jobs); i++ {
		jobLoadMap[jobs[i].Interval] = jobs[i].Load
	}
	return jobLoadMap
}
