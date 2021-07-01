package back

import (
	"fmt"
	"testing"
)

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

func Test_t1(t *testing.T) {
	profit := maxProfit([]int{7, 6, 4, 3, 1})
	fmt.Println(profit)
}

//贪心算法
//只要交易日大于前一天的交易日就卖出
//price[i]>price[i-1]
//对于连续的上涨的交易日Pn(连续上涨的最后一天)  pn-p1=(p2-p1)+(p3-p2)+...(pn-pn-1)==pn-p1

func maxProfit(prices []int) int {

	var total int
	for i := 0; i < len(prices); i++ {
		if i > 0 && prices[i] > prices[i-1] {
			total += prices[i] - prices[i-1]
		}
	}
	return total
}

//动态规划
//对于每个数组下标i只有两种状态,要就是买,要就是不买
//设置状态转移方程dp[i][j]表示第i天交易完后手中有没有股票
//dp表示手中的现金的最大值
//dp[i][0]当前没有股票(前一天手中也没有持有股票,):dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i])
//dp[i][0] 其实相当于卖的状态
//dp[i][1]=max(dp[i-1][1],dp[i-1][0]-prices[i])
func maxProfit2(prices []int) int {
	dp := make([][2]int, 0, len(prices))
	//没有卖出收益为0,什么也没干
	dp[0][0] = 0
	//第一天买入收益为0-prices[0]
	dp[0][1] = -prices[0]
	for i := 0; i < prices[i]-1; i++ {
		//当前没有持股票的现金
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		//前一天买了,今天不买,今天买前一天卖出
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(dp)-1][0]
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//递归算法
func maxProfit3(prices []int) int {

	dfs(prices, 0, 0, 0, nil)
	return 0
}

//status=0,表示卖出
//status=1,表示买入
func dfs(prices []int, status int, index int, money int, res *int) {
	if index == len(prices) {
		*res = max(*res, money)
	}
	//这个递归直接向下,表示这天什么也没干
	dfs(prices, status, index+1, money, res)
	if status == 0 {
		//试着买入股票
		dfs(prices, 1, index+1, money-prices[index], res)
	}
	if status == 1 {
		//试着卖出股票
		dfs(prices, 0, index+1, money+prices[index], res)
	}
}
