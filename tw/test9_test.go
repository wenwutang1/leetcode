package tw

import (
	"fmt"
	"sort"
	"testing"
)
//[2,3,1,3,2,4,6,7,9,2,19]
//[2,1,4,3,9,6]
func Test_r1(t *testing.T){
	fmt.Println(relativeSortArray([]int{2,3,1,3,2,4,6,7,9,2,19},[]int{2,1,4,3,9,6}))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	//排序arr1
	sort.Ints(arr1)
	var result []int=make([]int,0, len(arr1))
	//记录arr2中没有的数据
	var flag   []int8=make([]int8, len(arr1))
	//放入的数量
	var c int
	for i:=0;i< len(arr2);i++{
		//二分法
		l,r:=0, len(arr1)-1
		for l<=r{
			mid:=int(uint(l+r)>>1)
			if arr1[mid]==arr2[i]{
				flag[mid]=1
				result=append(result,arr1[mid])
				c++
				//双指针查找
				for in1,in2:=mid-1,mid+1;in1>-1&&arr1[in1]==arr2[i]||in2< len(arr1)&&arr1[in2]==arr2[i];{
					if in1>-1&&arr1[in1]==arr2[i]{
						result=append(result,arr1[mid])
						flag[in1]=1
						c++
					}
					if in2< len(arr1)&&arr1[in2]==arr2[i]{
						result=append(result,arr1[mid])
						flag[in2]=1
						c++
					}
					in1--
					in2++
				}
				break
			}else if arr1[mid]>arr2[i]{
				r=mid-1
			}else {
				l=mid+1
			}
		}
	}
	//遍历arr1,装入剩下的
	if c< len(arr1){
		for i:=0;i< len(flag);i++{
			if flag[i]==0{
				result=append(result,arr1[i])
			}
		}
	}
	return result
}

