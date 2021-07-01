package back

import (
	"fmt"
	"sort"
	"testing"
)

func Test_binary(t *testing.T){

	array:=make([]int,100)
	for i:=0;i< len(array);i++{
		array[i]=i+1
	}
	index:=BinarySearch2(array,65,0, len(array)-1)
	fmt.Println(array[index])

}

//基础二分查找法
func BinarySearch(array []int,target int)int{
	sort.Ints(array)
	l,r:=0, len(array)-1
	for l<=r{
		mid:=l+(r-l)>>1
		if array[mid]==target{
			return mid
		}else if array[mid]>target{
			r=mid-1
		}else {
			l=mid+1
		}
	}
	return -1
}

func BinarySearch2(array []int,target int,l,r int)int{
	if l>=r{
		return -1
	}
	mid:=l+(r-l)>>1
	if array[mid]==target{
		return mid
	}
	if array[mid]>target{
		return BinarySearch2(array,target,l,mid-1)
	}else{
		return BinarySearch2(array,target,mid+1,r)
	}
}
