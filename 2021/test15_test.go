package _021

import (
	"fmt"
	"testing"
)

func Test15(t *testing.T){

	//fmt.Printf("%b",reverseBits(0b11010))
	//fmt.Printf("%b\n",0b001+0b001)
	//fmt.Printf("%b\n",0b001|0b001)
	//00000010100101000001111010011100
	fmt.Println(reverseBits(0b00000010100101000001111010011100))
}

//把一个十进制数字进行翻转
func reverse(x int) int {


	var res int64
	for x!=0{
		if res>(1<<31-1)||res<(-1<<31){
			return 0
		}
		res=res*10+int64(x%10)
		x=x/10
	}
	return int(res)
}


func reverseBits(num uint32) uint32 {
	var res uint32
	for i:=0;i<32;i++{
		res=res<<1+num&1
		num>>=1
	}
	return res
}
