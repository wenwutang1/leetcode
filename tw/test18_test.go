package tw

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_can(t *testing.T){


	fmt.Println(fizzBuzz(15))

}

func canConstruct(ransomNote string, magazine string) bool {

	var mmp=make([]int,26)
	for i:=0;i< len(magazine);i++{
		mmp[magazine[i]-'a']++
	}

	for i:=0;i< len(ransomNote);i++{
		mmp[ransomNote[i]-'a']--
		if mmp[ransomNote[i]-'a']<0{
			return false
		}
	}
	return true
}

func arrangeCoins(n int) int {
	return 0
}

func fizzBuzz(n int) []string {
	var result =make([]string,0,n)
	type F func(int)
	var f F
	f=func (a int){
		if a>n{
			return
		}
		if a%3==0&&a%5==0{
			result=append(result,"FizzBuzz")
		}else if a%3==0{
			result=append(result,"Fizz")
		}else if a%5==0{
			result=append(result,"Buzz")
		}else{
			result=append(result,strconv.Itoa(a))
		}
		a++
		f(a)
	}
	f(1)
	return result
}
