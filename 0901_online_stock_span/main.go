// Problem: https://leetcode.com/problems/online-stock-span/
// Results: https://leetcode.com/submissions/detail/839791999/
package main

type Stock struct {
	price int
	span  int
}

type StockSpanner struct {
	stack []Stock
}

func Constructor() StockSpanner {
	return StockSpanner{stack: []Stock{}}
}

func (this *StockSpanner) Next(price int) int {
	stock := Stock{price, 1}
	for len(this.stack) > 0 && this.stack[len(this.stack)-1].price <= stock.price {
		stock.span += this.stack[len(this.stack)-1].span
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, stock)
	return stock.span
}
