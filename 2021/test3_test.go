package _021

import (
	"fmt"
	"testing"
)

//{1,2,3,4,5}
//9
var res =make([][]int,0)
func Test_back1(t *testing.T){
	var arr=[]int{1,2,3,4,5}
	back(arr,[]int{},0,9,0)
	fmt.Println(res)
}

func back(arr ,road []int,index,target,sum int){

	for i:=index;i< len(arr);i++{
		if sum+arr[i]<target{
			back(arr,append(road,arr[i]),i+1,target,sum+arr[i])
		}
		if sum+arr[i]==target{
			temp:=make([]int, len(road)+1)
			copy(temp,append(road,arr[i]))
			res=append(res,temp)
		}
	}
}


