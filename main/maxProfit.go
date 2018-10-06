package main

import "fmt"

/**

Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.

 */

func maxProfit(prices []int) int {
	if len(prices)==0 {
		return 0
	}
	type data struct {
		value int
		index int
	}
	min := &data{prices[0],0}
	max := &data{prices[0],0}
	l:=len(prices)
	for i:=0;i<l;i++{
		if min.value>prices[i] {
			min.value=prices[i]
			min.index=i
			continue
		}
		if max.value<prices[i] {
			max.value=prices[i]
			max.index=i
			continue
		}
	}
	if min.index<max.index {
		return max.value-min.value
	}
	return 0
}

func main()  {
	fmt.Println(maxProfit([]int{2,4,1}))
}
