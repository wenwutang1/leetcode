package tw

import (
	"testing"
)

func Test_ms1(tt *testing.T){


}

//一个长度为n-1的递增排序数组中的所有数字都是唯一的,并且每个数字都在范围0～n-1之内.
//在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字.

//分治算法
//最后分成两组
//[0,1,2,3,4,5,6,7,9]
//8
func Test_rr(t *testing.T){
	//fmt.Println(arrangeCoins(21))
}
func missingNumber(nums []int) int {
	l,r:=0, len(nums)-1
	for l<=r{
		mid:=int(uint(l+r)>>1)
		if nums[mid]==mid{
			l=mid+1
		}else{
			r=mid-1
		}
	}
	return l
}



//标准二分模板
func search1(nums []int, target int) int {
	l,r:=0, len(nums)-1
	for l<=r{
		mid:=int(uint(r+r)>>1)
		if nums[mid]==target{
			return mid
		} else if nums[mid]>target{
			l=mid-1
		}else{
			r=mid+1
		}
	}
	return -1
}

