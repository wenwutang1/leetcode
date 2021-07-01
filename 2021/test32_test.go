package _021

import (
	"fmt"
	"testing"
)

func Test33(t *testing.T){
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}


// 连续最大和
// [-2,1,-3,4,-1,2,1,-5,4]
func maxSubArray(nums []int) int {
	if len(nums)<1{
		return 0
	}
	var dp =make([]int, len(nums))
	dp[0]=nums[0]
	var max =dp[0]
	for i:=1;i< len(nums);i++{
		dp[i]=max1(dp[i-1]+nums[i],nums[i])
		if max<dp[i]{
			max=dp[i]
		}
	}
	return max
}

func max1(a,b int)int{
	if a<b{
		return b
	}
	return a
}
