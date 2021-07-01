package _021

import (
	"fmt"
	"testing"
)

func Test19(t *testing.T){
	fmt.Println(subsets([]int{1,2}))
}

func subsets(nums []int) [][]int {

	res:=make([][]int,0)
	subsets2(nums,[]int{},0,&res)
	return res
}


//road 路径
//index 索引
//result结果集
func subsets2(nums []int,road []int,index int,result *[][]int){

	if index> len(nums){
		return
	}
	temp:=make([]int, len(road))
	copy(temp,road)
	*result=append(*result,temp)
	for i:=index;i< len(nums);i++{
		subsets2(nums,append(road,nums[i]),i+1,result)
	}
}