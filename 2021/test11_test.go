package _021

import (
	"fmt"
	"testing"
)

func Test12(t *testing.T){
	var div=0.0
	fmt.Println(0.0/div) //NaN

	var div1=0.0
	fmt.Println(1/div1) //+Inf

}

func divingBoard(shorter int, longer int, k int) []int {
	var set=make(map[int]byte)
	divingBoard2(shorter,longer,k,0,set)
	fmt.Println(set)
	return nil
}

func divingBoard2(shorter int, longer int, k int,sum int,set map[int]byte){
	if k==0{
		set[sum]=1
		return
	}
	divingBoard2(shorter,longer,k-1,sum+shorter,set)
	divingBoard2(shorter,longer,k-1,sum+longer,set)
}




