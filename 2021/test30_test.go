package _021

import (
	"fmt"
	"testing"
)

func Test30(t *testing.T){
	//reverse123(-123)
	//fmt.Println(123/100)
	//fmt.Println(isPalindrome(101))
	//fmt.Println(romanToInt2("IV"))
	fmt.Println("13:00"<"12:00")
}

//123
func reverse123(x int) int {
	var result int
	for x!=0{
		if result< (-1<<31/10) ||result>(1<<31-1)/10{
			return 0
		}
		result=result*10+(x%10)
		x=x/10
	}
	return result
}


func isPalindrome(x int) bool {
	if x<0{
		return false
	}
	var old =x
	var nums int
	for x!=0{
		nums=nums*10+x%10
		x=x/10
	}
	return nums==old
}

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
func romanToInt(s string) int {
	var mmp=map[byte]int{
		'I':1,
		'V':5,
		'X':10,
		'L':50,
		'C':100,
		'D':500,
		'M':1000,
	}
	var result int
	for i:=0;i< len(s)-1;i++{
		if mmp[s[i]]<mmp[s[i+1]]{
			result+=-mmp[s[i]]
		}else{
			result+=mmp[s[i]]
		}
	}
	result+=mmp[s[len(s)-1]]
	return result
}


func romanToInt2(s string)int{
	var result int
	for i:=0;i< len(s);i++{
		if i< len(s)-1&&Index(s[i])<Index(s[i+1]){
			result+=-Index(s[i])
		}else{
			result+=Index(s[i])
		}
	}
	return result
}

func Index(i byte )int{
	switch  i{
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 100
	default:
		return 0
	}
}