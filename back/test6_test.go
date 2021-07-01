package back

import (
	"fmt"
	"sort"
	"testing"
)

func Test_next1(t *testing.T) {

	var result =[]int{4,2,0,2,3,2,0} //4203022
	nextPermutation(result)
	fmt.Println(result)
}

//实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//必须原地修改，只允许使用额外常数空间。
//以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
//1,2,3 → 1,3,2
//3,2,1 → 1,2,3
//1,1,5 → 1,5,1
func nextPermutation(nums []int) {
	//倒过来寻找
	if len(nums)==1{
		return
	}
	var i = len(nums) - 1
	for ; i > 0; i-- {
		var j=i-1
		for ;j>-1;j--{
			if nums[i]>nums[j]{
				//交换
				swap(nums,i,j)
				//排序,找到比j大的第一个数
				//排序后面的数
				sort1(nums,j+1)
				return
			}
		}
	}
	if i==0{
		//交换
		revert(nums)
	}
}


func sort1(nums []int,index int){
	temp:=nums[index:]
	sort.Slice(temp, func(i, j int) bool {
		return temp[i]<temp[j]
	})
}

func swap(nums []int,i,j int){
	nums[i],nums[j]=nums[j],nums[i]
}
func revert(nums []int){
	for i,j:= len(nums)-1,0;i>j;i,j=i-1,j+1{
		nums[i],nums[j]=nums[j],nums[i]

	}
}