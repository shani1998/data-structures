package questions

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
*/

// return diff of the local_max - local_min
func maxProfit(prices []int) int {
	totalProfit, totalDays := 0, len(prices)
	maxSoFar := make([]int, totalDays)

	//populate max right of each element
	currMax := -1
	for i := totalDays - 1; i >= 0; i-- {
		if prices[i] > currMax {
			currMax = prices[i]
		}
		maxSoFar[i] = currMax
	}

	// find profit for each day buy
	for i, price := range prices {
		currProfit := maxSoFar[i] - price
		if currProfit > totalProfit {
			totalProfit = currProfit
		}
	}

	return totalProfit
}
