// Problem: https://leetcode.com/problems/time-based-key-value-store/
// Results: https://leetcode.com/submissions/detail/816447866/
package main

type ValueNode struct {
	timestamp int
	value     string
}

type TimeMap struct {
	data map[string][]ValueNode
}

func Constructor() TimeMap {
	return TimeMap{data: map[string][]ValueNode{}}
}

func (this *TimeMap) closestIndex(key string, timestamp int) int {
	arr, exists := this.data[key]
	if !exists {
		return -1
	}
	start, end := 0, len(arr)-1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid].timestamp < timestamp {
			start = mid + 1
		} else {
			if arr[mid].timestamp == timestamp {
				return mid
			}
			end = mid - 1
		}
	}
	if start < len(arr)-1 && arr[start].timestamp < timestamp {
		start++
	}
	return start
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	i := this.closestIndex(key, timestamp)
	if i == -1 {
		this.data[key] = append(this.data[key], ValueNode{
			timestamp: timestamp,
			value:     value,
		})
	} else {
		curr := this.data[key]
		if i < len(curr) && curr[i].timestamp == timestamp {
			curr[i].value = value
			return
		}
		curr = append(curr, ValueNode{})
		copy(curr[i+1:], curr[i:len(curr)-1])
		curr[i].timestamp = timestamp
		curr[i].value = value
		this.data[key] = curr
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	i := this.closestIndex(key, timestamp)
	if i == -1 {
		return ""
	}
	curr := this.data[key]
	if i == len(curr) || (i > 0 && curr[i].timestamp > timestamp) {
		i--
	}
	if curr[i].timestamp > timestamp {
		return ""
	}
	return curr[i].value
}
