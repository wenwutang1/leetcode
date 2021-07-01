package _021

import (
	"fmt"
	"testing"
)

func Test9(t *testing.T){

	i := fib2(5)
	fmt.Println(i)
}

func fib(n int) int {
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}
	dp:=make([]int,n+1)
	dp[0]=0
	dp[1]=1
	for i:=2;i<=n;i++{
		dp[i]=(dp[i-1]+dp[i-2])%(1e9+7)
	}
	return dp[n]
}


func fib2(n int)int{
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}
	dp:=make([]int,2)
	dp[0]=0
	dp[1]=1
	for i:=2;i<=n+1;i++{
		dp[0],dp[1]=dp[1],(dp[0]+dp[1])%(1e9+7)
	}
	return dp[1]
}
