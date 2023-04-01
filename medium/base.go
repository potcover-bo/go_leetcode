package main

import "potcover.cc/go_leetcode/utils"

/*
*122. 买卖股票的最佳时机 II
给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。
*/
func maxProfit(prices []int) int {
	dp_i_1_0 := 0
	dp_i_1_1 := -prices[0]

	dp_i_0 := 0
	dp_i_1 := 0

	for i := 1; i < len(prices); i++ {
		dp_i_1 = utils.Max(dp_i_1_1, dp_i_1_0-prices[i])
		dp_i_0 = utils.Max(dp_i_1_0, dp_i_1_1+prices[i])

		dp_i_1_0 = dp_i_0
		dp_i_1_1 = dp_i_1
	}

	return dp_i_0
}

/*
*714. 买卖股票的最佳时机含手续费
给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。

你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

返回获得利润的最大值。

注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
*/
func maxProfitFee(prices []int, fee int) int {
	//TODO
	return 0
}

/*
*309. 最佳买卖股票时机含冷冻期
给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/
func maxProfitFreeze(prices []int) int {
	//TODO

	return 0

}
