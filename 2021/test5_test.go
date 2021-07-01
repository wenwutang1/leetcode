package _021

import (
	"fmt"
	"testing"
)

func Test5(t *testing.T){
	//全排列
	num:=[]int{1,2,3}
	res := sortTotal4(num)
	fmt.Println(res)
}


var result=make([][]int,0)
func totalSort(num []int)[][]int{
	flag:=make([]int, len(num))
	totalsort2(num,flag,[]int{})
	return result
}


func totalsort2(num []int,flag []int,road []int){
	if len(road)== len(num){
		src:=make([]int, len(road))
		copy(src,road)
		result=append(result,src)
		return
	}
	for i:=0;i< len(num);i++{
		if flag[i]==1{
			continue
		}
		road=append(road,num[i])
		flag[i]=1
		totalsort2(num,flag,road)
		//把上一个标记为
		flag[i]=0
		road=road[:len(road)-1]
	}
}


//组合
//[1,2,3] [1],[2],[3]

func sortTotal4(num []int)[][]int{
	road:=make([]int,0)
	sortTotal42(num,road,0)
	return result

}

func sortTotal42(num []int,road []int,index int){
	if len(road)> len(num) {
		return
	}
	for i:=index;i< len(num);i++{
		road=append(road,num[i])
		temp:=make([]int, len(road))
		copy(temp,road)
		result=append(result,temp)
		sortTotal42(num,road,i+1)
		road=road[:len(road)-1]
	}
}