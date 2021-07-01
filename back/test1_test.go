package back

import (
	"fmt"
	"testing"
)

func Test_back1(t *testing.T){
	Back1([]int{1,2,3,4,5},9)
	fmt.Println(result)

}

var result=make([][]int,0)
//{1,2,3,4,5}
//9

func Back1(num []int,target int)[][]int{
	var road =make([]int,0, len(num))
	//back(num,target,road,0,0)
	flag:=make([]int,len(num))
	back1(num,road,flag,target,0)
	return nil
}

func back(num []int,target int,road []int,index int,sum int){
	for i:=index;i< len(num);i++{
		if sum+num[i]<target{
			road=append(road,num[i])
			back(num,target,road,i+1,sum+num[i])
			road=road[:len(road)-1]
		}
		//剪枝
		if sum+num[i]==target{
			temp:=make([]int, len(road)+1)
			copy(temp,append(road,num[i]))
			result=append(result,temp)
			//road=road[:len(road)-1]
			//return
		}
	}
}

//////////////////////////全搜索//////////////////////
func back1(num []int,road []int,flag []int,target,sum int){
	for i:=0;i< len(num);i++{
		if flag[i]==1{
			continue
		}
		if sum+num[i]<target{
			road=append(road,num[i])
			//设置走过的节点
			flag[i]=1
			back1(num,road,flag,target,sum+num[i])
			road=road[:len(road)-1]
			flag[i]=0
		}
		//剪枝?
		if sum+num[i]==target{
			t := make([]int, len(road)+1)
			copy(t,append(road,num[i]))
			result=append(result,t)
		}
	}
}



//.给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
//
//例如，给出 n = 3，生成结果为
//回溯算法2
//[
//
//  "((()))",
//
//  "(()())",
//
//  "(())()",
//
//  "()(())",
//
//  "()()()"
//
//]
func Test_back2(t *testing.T){
	back3(3)
}

func back3(n int){

	var result []string=make([]string,0)
	answer(n*2,"",&result)
	fmt.Println(result)
	return
}

func answer(n int,road string,result *[]string){
	if len(road)<n{
		//减枝
		//右括号开头的显然不能构成
		if len(road)>0&&road[0]==41{
			return
		}
		answer(n,road+"(",result)
		answer(n,road+")",result)
	}else{
		if isValid(road){
			*result=append(*result,road)
		}
	}
}

//判断是否是有效的括号
func isValid(str string)bool{
	//遇到左括号入栈,右括号出栈
	var stack=make([]uint8,0, len(str))
	for i:=0;i< len(str);i++{
		if str[i]==40{
			stack=append(stack,str[i])
		}else if len(stack)>0&&str[i]==41{
			stack=stack[:len(stack)-1]
		}else{
			return false
		}
	}
	if len(stack)!=0{
		return false
	}
	return true
}









