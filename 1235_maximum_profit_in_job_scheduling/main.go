// Problem: https://leetcode.com/problems/maximum-profit-in-job-scheduling/
// Results: https://leetcode.com/submissions/detail/850127944/
package main

import "sort"

type Job struct {
	start  int
	end    int
	profit int
}

func dfs(cache []int, jobs []Job, i int) int {
	if i == len(jobs) {
		return 0
	}
	if cache[i] > 0 {
		return cache[i]
	}
	next := i + 1
	for next < len(jobs) && jobs[next].start < jobs[i].end {
		next++
	}
	with := jobs[i].profit + dfs(cache, jobs, next)
	skip := dfs(cache, jobs, i+1)
	if skip > with {
		with = skip
	}
	cache[i] = with
	return with
}

func jobScheduling(startTime, endTime, profit []int) int {
	size := len(startTime)
	jobs := make([]Job, size)
	for i := range jobs {
		jobs[i] = Job{
			start:  startTime[i],
			end:    endTime[i],
			profit: profit[i],
		}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].start < jobs[j].start
	})
	return dfs(make([]int, size), jobs, 0)
}
