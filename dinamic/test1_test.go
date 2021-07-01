package dinamic

import (
	"fmt"
	"sort"
	"testing"
)



type A struct {
	Name string
}

func (this A)SSS(){}

type B struct {
	*A
}

func Test_d1(t *testing.T){

	b:=&B{}
	b.SSS()
	fmt.Println(b.Name)
}

//给定不同面额的硬币 coins 和一个总金额 amount。
//编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
//如果没有任何一种硬币组合能组成总金额，返回-1。
//你可以认为每种硬币的数量是无限的。
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1

//dp[i][j] 当前硬币数对应的最少的组合
//dp[i][j]=min(dp[i-1][j]-cur([arr[i]]),dp[i-1][j])
func maximumSum(arr []int) int {
	return 0
}



//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
func moveZeroes(nums []int)  {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]<nums[j]
	})
	fmt.Println("sort",nums)
	var j=0
	for i:=0;i< len(nums)&&nums[i]==0;i++{
		j++
	}
	//将0后移
	if j==0{
		return
	}
	//数组位置前移动
	for i:=j;i< len(nums);i++{
		nums[i-j]=nums[i]
	}
	for i:= len(nums)-j;i< len(nums);i++{
		nums[i]=0
	}
}

// s = "anagram", t = "nagaram"
func isAnagram(s string, t string) bool {
	if len(s)!= len(t){
		return false
	}
	var nums1=make([]int,26)
	for i:=0;i< len(s);i++{
		nums1[s[i]-'a']++
	}
	for i:=0;i< len(t);i++{
		index:=t[i]-'a'
		nums1[index]--
		if nums1[index]<0{
			return false
		}
	}
	return true
}


//羅馬數字
func romanToInt(s string) int {

	var sum =0
	for i:=0;i< len(s);i++{
		if i< len(s)-1&&Index(s[i])<Index(s[i+1]){
			sum=sum-Index(s[i])
		}else{
			sum=sum+Index(s[i])
		}
	}
	return sum
}

func Index(a byte)int{
	switch a {
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
		return 1000
	}
	return 0
}
