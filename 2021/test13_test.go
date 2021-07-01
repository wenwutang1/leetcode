package _021

import (
	"fmt"
	"testing"
)

func Test14(t *testing.T){

	var array= []int{1,2,3,4,5}
	array=append(array,9999)
	fmt.Println(len(array),cap(array)) //6,10
	AppendArr(array)
	fmt.Println(array)//1,2,3,4,5,9999
	a:=array[0:7] //1,2,3,4,5,9999,6
	fmt.Println(a)
	fmt.Printf("%p\n",array)
}


func AppendArr(arr []int){
	//arr[0]=100
	arr=append(arr,6)
	fmt.Printf("%p\n",arr)//7,10
	fmt.Println(len(arr), cap(arr))
	//fmt.Println(arr)
}
