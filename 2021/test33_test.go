package _021

import (
	"fmt"
	"testing"
)

func Test34(t *testing.T){

	//fmt.Println(lengthOfLastWord("hello world 1golangttt go "))
	fmt.Println(plusOne([]int{9,9}))
}

func lengthOfLastWord(s string) int {
	if len(s)<1{
		return 0
	}
	var count int
	for j:= len(s)-1;j>-1;j--{
		if s[j]==' '{
			if count==0{
				continue
			}else{
				break
			}
		}
		count++
	}
	return count
}



func plusOne(digits []int) []int{
	return plusOne111(digits,1)
}

func plusOne111(digits []int,creb int) []int {

	var flag =0
	for j:= len(digits)-1;j>-1;j--{
		if j== len(digits)-1{
			digits[j]=digits[j]+creb
		}else{
			digits[j]=digits[j]+flag
		}
		if digits[j]>=10{
			flag=digits[j]/10
			digits[j]=digits[j]%10
		}else{
			flag=0
		}
	}
	//顶位需要进位,数组长度需要扩容
	if flag>0{
		digits=append([]int{flag},digits...)
	}
	return digits
}
