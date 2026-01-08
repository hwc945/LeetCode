package lc

func MaxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	//记录当前最大利润和最小买入价
	pro, minNum := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if pro < prices[i]-minNum {
			pro = prices[i] - minNum
		}
		if minNum > prices[i] {
			minNum = prices[i]
		}
	}
	return pro
}
func maxProfit(prices []int) int {
	// 1. 别写那堆没用的 len 检查了
	// 2. 初始化：既然要找最小值，就设为第一天价格；利润设为0
	minPrice := 0 // 或者直接用 prices[0]，但注意判空
	if len(prices) > 0 {
		minPrice = prices[0]
	}
	maxProfit := 0

	for _, price := range prices {
		// 更新历史最低价
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			// 只有当前价格比最低价高时，才可能产生更大利润
			// 用 else if 减少一次无意义的比较
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
