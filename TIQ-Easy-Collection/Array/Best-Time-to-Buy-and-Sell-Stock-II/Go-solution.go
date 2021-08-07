package Best_Time_to_Buy_and_Sell_Stock_II

func maxProfit(prices []int) int {
	var stonks int = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			stonks += prices[i] - prices[i-1]
		}
	}
	return stonks
}
