package _021

import (
	"fmt"
	"testing"
)

////给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
////设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//列如:7, 6, 4, 3, 1.你只能不能交易因为他是不能赚钱的


//贪心算法
//只能交易前一日的股票小于交易后一日的股票就卖出
//

func Test_test5(t *testing.T){
	sum:=0
	caobi2([]int{7,8,5,4,3,2,1},0,0,0,&sum)
	fmt.Println(sum)
}


func caobi(arr[]int)int{
	var total int
	for i:=0;i< len(arr)-1;i++{
		if arr[i]<arr[i+1]{
			total=arr[i+1]-arr[i]
		}
	}
	return total
}


//递归,回溯法
//用一个状态stat表示当前入或者卖出
func caobi2(arr []int,stat int,index int,money int,sum *int){
	if index== len(arr){
		if *sum<money{
			*sum=money
		}
		return
	}
	//什么都不做
	caobi2(arr,stat,index+1,money,sum)
	if stat==0{
		//买入股票
		caobi2(arr,1,index+1,money-arr[index],sum)
	}
	if stat==1{
		//卖出
		caobi2(arr,0,index+1,money+arr[index],sum)
	}
}
