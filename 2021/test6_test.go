package _021

import (
	"fmt"
	"testing"
)

func Test_tt(t *testing.T){

	fmt.Println(generateParenthesis(10))

}



var result1 =make([]string,0)

func generateParenthesis(n int) []string {

	generatePath(n,0,"")
	return result1
}

func generatePath(n int,index int,road string){
	if index<2*n{
		if road!=""&&road[0]==')'{
			return
		}
		generatePath(n,index+1,road+"(")
		generatePath(n,index+1,road+")")
	}else{
		//判断是否是一个括号
		if judgePath(road){
			result1=append(result1,road)
		}
	}
}

func judgePath(str string)bool{
	stack:=make([]byte,0, len(str))
	for i:=range str{
		if str[i]==40{
			stack=append(stack,40)
		}else if str[i]==41{
			if len(stack)<1{
				return false
			}
			stack=stack[:len(stack)-1]
		}
	}
	return len(stack)<=0
}