package tw

import (
	"fmt"
	"testing"
)

func Test_flog(t *testing.T) {
	//[23,2,6,4,7]
	//0
	sum := checkSubarraySum([]int{0, 1, 0}, 0)
	fmt.Println(sum)

}

/////////////////////////////青蛙//////////////////////////
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	dp := make([]int, len(nums))
	if k == 0 {
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums)-i; j++ {
				dp[j] = dp[j] + nums[j+i]
				if i != 0 && dp[j] == 0 {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			dp[j] = (dp[j] + nums[j+i]) % k
			if i != 0 && dp[j] == 0 {
				return true
			}
		}
	}
	return false

}
