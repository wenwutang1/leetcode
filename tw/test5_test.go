package tw

import (
	"fmt"
	"testing"
)


func Test_Kw(t *testing.T){
	var result [][]int=[][]int{
		{1,1,0,0,0},
		{1,1,1,1,0},
		{1,0,0,0,0},
		{1,1,0,0,0},
		{1,1,1,1,1},
	}
	rows := kWeakestRows(result, 3)
	fmt.Println(rows)
}

func kWeakestRows(mat [][]int, k int) []int {
	type drx struct {
		index int
		value int
	}
	var result =make([]drx,0,k)
	for i:=0;i< len(mat);i++{
		//二分法
		var l,r=0, len(mat[i])-1
		if mat[i][r]==1{
			result=append(result,drx{
				index: i,
				value: r+1,
			})
			continue
		}
		if mat[i][0]==0{
			result=append(result,drx{
				index: i,
				value: 0,
			})
			continue
		}
		for l<=r{
			var mid =int(uint(l+r)>>1)
			if mat[i][mid]>0{
				l=mid+1
			}else {
				r=mid-1
			}
		}
		result=append(result,drx{
			index: i,
			value: l,
		})
	}
	//排序
	//var s []int=make([]int,0,k)
	for i:=0;i< len(result)-1;i++{
		for j:=0;j< len(result)-i-1;j++{
			if result[j].value>result [j+1].value{
				result[j],result[j+1]=result[j+1],result[j]
			}
		}
	}
	var array []int=make([]int,0,k)
	for i:=0;i< k;i++{
		array=append(array,result[i].index)
	}
	return array
}




func sortm(array []int){
	for i:=0;i< len(array)-1;i++{
		for j:=0;j< len(array)-1-i;j++{
			if array[j]>=array[j+1]{
				array[j],array[j+1]=array[j+1],array[j]
			}
		}
	}
	fmt.Println("sort",array)
}