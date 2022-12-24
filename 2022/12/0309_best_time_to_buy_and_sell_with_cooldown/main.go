// Problem: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// Results: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/submissions/864070970/
package main

const (
	Buy = iota
	Sell
	Cooldown
)

type Key struct {
	idx    int
	action int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func seek(cache map[Key]int, prices []int, i, prev int) int {
	if i == len(prices) {
		return 0
	}
	k := Key{i, prev}
	if n, cached := cache[k]; cached {
		return n
	}
	if prev == Sell {
		cache[k] = seek(cache, prices, i+1, Cooldown)
	} else if prev == Buy {
		skip := seek(cache, prices, i+1, Buy)
		sell := seek(cache, prices, i+1, Sell) + prices[i]
		cache[k] = max(skip, sell)
	} else {
		skip := seek(cache, prices, i+1, Cooldown)
		buy := seek(cache, prices, i+1, Buy) - prices[i]
		cache[k] = max(skip, buy)
	}
	return cache[k]
}

func maxProfit(prices []int) int {
	cache := map[Key]int{}
	res := seek(cache, prices, 0, Cooldown)
	return res
}
