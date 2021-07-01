package tw

import (
	"fmt"
	"testing"
)

func TestY(t *testing.T) {

	ways := findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)
	fmt.Println(ways)
}

func findTargetSumWays(nums []int, S int) int {
	var num = 0
	findTargetSumWays2(nums, S, 0, 0, &num)
	return num
}

func findTargetSumWays2(nums []int, S int, index int, sum int, result *int) {
	if index == len(nums) {
		if sum == S {
			*result++
		}
		return
	}
	//+
	findTargetSumWays2(nums, S, index+1, sum+nums[index], result)
	//-
	findTargetSumWays2(nums, S, index+1, sum-nums[index], result)
	//只需要循环最外层的数字
}
