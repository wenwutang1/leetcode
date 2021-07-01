package _021

import (
	"fmt"
	"testing"
)

func Test35(t *testing.T){
	fmt.Println(climbStairs2(3))

}


func addBinary(a string, b string) string {

	var result =make([]byte,0)
	var index1, index2= len(a)-1, len(b)-1
	for index1>-1&&index2>-1{
		if a[index1]==b[index2]&&a[index1]=='1'{
			result=append(result,'0')
		}else if a[index1]==b[index2]&&a[index1]=='0'{
			result=append(result,'0')
		}else{
			result=append(result,'1')
		}
		index1--
		index2--
	}

	return ""
}



//0 1 2
func climbStairs(n int) int {
	if n==1{
		return 1
	}
	if n==2{
		return 2
	}
	var dp []int=make([]int,2)
	dp[0]=1
	dp[1]=2
	for i:=3;i<=n;i++{
		dp[0],dp[1]=dp[1],dp[1]+dp[0]
	}
	return dp[1]
}

func climbStairs2(n int) int {
	if n==1{
		return 1
	}
	if n==2{
		return 2
	}
	var dp []int=make([]int,n)
	dp[0]=1
	dp[1]=2
	for i:=2;i<n;i++{
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[len(dp)-1]
}