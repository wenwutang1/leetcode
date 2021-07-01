package tw

import (
	"fmt"
	"testing"
)

func Test_four(t *testing.T){
//A = [ 1, 2]
	//B = [-2,-1]
	//C = [-1, 2]
	//D = [ 0, 2]
	fmt.Println(string(findTheDifference("ae","aea")))
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	var res int=0
	var mmp =make(map[int]int, len(A)*len(B))
	for i:=0;i< len(A);i++{
		for j:=0;j< len(B);j++{
			mmp[A[i]+B[j]]++
		}
	}
	for i:=0;i< len(C);i++{
		for j:=0;j< len(D);j++{
			res+=mmp[-C[i]-D[j]]
		}
	}
	return res
}


func isHappy(n int) bool {
	slow,fast:=n,deal(n)
	for fast!=1&&slow!=fast{
		slow=deal(slow)
		fast=deal(deal(fast))
	}
	return fast==1
}

func deal(n int)int{
	sum := 0
	for n > 0 {
		sum += (n%10) * (n%10)
		n = n/10
	}
	return sum
}


func findTheDifference(s string, t string) byte {

	var nums=make([]int,26)
	for i:=0;i< len(s);i++{
		nums[s[i]-'a']++
	}
	var temp=make([]int,26)
	for i:=0;i< len(t);i++{
		temp[t[i]-'a']++
	}
	for i:=0;i< len(temp);i++{
		if temp[i]!=nums[i]{
			return byte('a'+i)
		}
	}
	return 0
}


