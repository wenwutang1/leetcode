package back

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T){
	subsetsWithDup([]int{1,2,2})
}


//给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）
//不能包含重复的子集
//输入: [1,2,2]
//输出:
//[
//  [2],
//  [1],
//  [1,2,2],
//  [2,2],
//  [1,2],
//  []
//]

func subsetsWithDup(nums []int) [][]int {
	var _=make([]int, len(nums))
	var result=make([][]int,0)
	back12(nums,[]int{},&result,0)
	fmt.Println(result)
	return nil
}


func back12(nums []int,road []int,result *[][]int,index int){
	if len(road)== len(nums){
		return
	}
	//不包含重复元素
	for i:=index;i< len(nums);i++{
		//重复数字去除
		road=append(road,nums[i])
		var temp=make([]int, len(road))
		copy(temp,road)
		*result=append(*result,temp)
		back12(nums,road,result,index+1)
		road=road[:len(road)-1]
	}
	road=[]int{}
}