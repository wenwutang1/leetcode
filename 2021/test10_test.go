package _021

import (
	"fmt"
	"testing"
)

//[58,48,64,36,19,19,67,13,32,2,59,50,29,68,50,0,69,31,54,20,22,43,30,9,68,71,20,22,48,74,2,65,27,54,30,5,66,24,64,68,9,31,50,59,15,72,6,49,11,71,12,61,5,66,30,1,2,39,59,35,53,21,76,17,71,40,68,57,64,53,70,21,50,49,25,63,35]
//46

func Test11(t *testing.T){

	fmt.Println(canReach([]int{58,48,64,36,19,19,67,13,32,2,59,50,29,68,50,0,69,31,54,20,22,43,30,9,68,71,20,22,48,74,2,65,27,54,30,5,66,24,64,68,9,31,50,59,15,72,6,49,11,71,12,61,5,66,30,1,2,39,59,35,53,21,76,17,71,40,68,57,64,53,70,21,50,49,25,63,35},46))
}
//跳跃数组
//只有两种调法+-
//start==arr.len return
//start==0       return
//只要往前或者往后一种情况达到目的即可完成
//怎么避免来回递归的问题 start=2 递归又回到start=2
func canReach(arr []int, start int) bool {

	var flag =make([]int, len(arr))
	return canReach2(arr,start,flag)
}

func  canReach2(arr []int, start int,flag []int) bool {
	if start>= len(arr){
		return false
	}
	if start<0{
		return false
	}
	//走了一圈又回来了
	if flag[start]==1{
		return false
	}
	//找到了
	if arr[start]==0{
		return true
	}
	flag[start]=1
	return canReach2(arr,start+arr[start],flag) ||canReach2(arr,start-arr[start],flag)
}


