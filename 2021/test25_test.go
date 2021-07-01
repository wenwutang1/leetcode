package _021

import (
	"fmt"
	"testing"
)

func Test26(t *testing.T){

	fmt.Println(findMin([]int{4,5,6,1,2,3}))
}




//水坑数组找最小值问题
//1  中值和数组右边的值进行比较
//2  如果num[mid]<num[len-1],最小值可能是num[mid]或者比num[mid]小
//3  如果num[mid]>=num[len-1],最小值一定在num[mid]右方
func findMin(nums []int) int {
	l,r:=0, len(nums)-1
	for l<r{
		mid:=(l+r)>>1
		if nums[mid]<nums[r]{
			r=mid
		}else{
			l=mid+1
		}
	}
	return nums[l]

}

//存在重复值的情况
func findMin2(nums []int)int{
	l,r:=0, len(nums)-1
	for l<r{
		mid:=(l+r)>>1
		if nums[mid]==nums[r]{
			r--
		}else if nums[mid]<nums[r]{
			r=mid
		}else{
			l=mid+1
		}
	}
	return nums[l]
}
