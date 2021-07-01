package tw

import (
	"fmt"
	"testing"
)

func Test_t1(t *testing.T){
	array := peakIndexInMountainArray([]int{0,2,1,0})
	fmt.Println(array)
}

//我们把符合下列属性的数组A称作山脉：
//A.length >= 3
//存在 0 < i< A.length - 1 使得A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
//给定一个确定为山脉的数组，返回任何满足A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]的 i的值
func peakIndexInMountainArray(arr []int) int {
	if len(arr)<1{
		return -1
	}
	return peakIndexInMountainArray2(arr,0, len(arr))
}


func peakIndexInMountainArray2(arr []int,l,r int)int{
	if l==r{
		return -1
	}
	var mid int=l+(r-l)>>1
	if arr[mid]>arr[mid-1]&&arr[mid]>arr[mid+1]{
		return mid
	}
	//左边
	if arr[mid]>arr[mid-1]{
		l=mid+1
	}else if arr[mid]>arr[mid+1]{ //右边
		r=mid
	}
	return peakIndexInMountainArray2(arr,l,r)
}
