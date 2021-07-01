package _021

import (
	"fmt"
	"testing"
)

func Test38(t *testing.T){

	var res []int=[]int{1,2,3,0,0,0}
	merge(res,3,[]int{2,5,6},3)
	fmt.Println(res)
}



func merge(nums1 []int, m int, nums2 []int, n int)  {
	if n==0{
		return
	}
	var i =0
	var result =make([]int, len(nums1))
	var index1 ,index2 =0,0
	for index2<n{
		if index1<m&&nums1[index1]<=nums2[index2]{
			result[i]=nums1[index1]
			index1++
		}else{
			result[i]=nums2[index2]
			index2++
		}
		i++
	}
	copy(nums1,result)
}
