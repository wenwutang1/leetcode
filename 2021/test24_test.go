package _021

import (
	"fmt"
	"testing"
)

func Test25(t *testing.T){
	stackUp([]int{10,15,30,70,50,60,30,11})
}



//单调栈
//单调递增栈，为了维护单调性[10,30,40,50]
// [10]
// [20]
// [30]
// [40]
// [40,50]
//维护一个单调栈
func stackUp(array []int){
	stack:=make([]int,0, len(array))
	for i:=0;i< len(array);i++{
		for len(stack)>0&&stack[len(stack)-1]<array[i]{
			stack=stack[:len(stack)-1]
		}
		stack=append(stack,array[i])
	}
	fmt.Println(stack)
}

//应用找到数组右边第一个比他小的数
func findFirstLow(array []int)[]int{
	stack:=make([]int,0, len(array))
	result:=make([]int,0, len(array))
	for i:=0;i< len(array);i++{
		for len(stack)>0&&stack[len(stack)-1]<array[i]{
			stack=stack[:len(stack)-1]
		}
		stack=append(stack,array[i])
	}

	return result
}