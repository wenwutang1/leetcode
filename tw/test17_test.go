package tw

import (
	"fmt"
	"sort"
	"testing"
)

func Test_ff(t *testing.T) {


}

/////////////////hg/////////////////////
func detectCapitalUse(word string) bool {
	//记录大写字母的个数
	var num = 0
	for i := 0; i < len(word); i++ {
		if 'A' <= word[i] && word[i] <= 'Z' {
			num++
		}
	}
	if !(num == len(word) || num == 1 || num == 0) {
		return false
	}
	if num != 0 && !('A' <= word[0] && word[0] <= 'Z') {
		return false
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	invertTree2(root)
	return root
}

func invertTree2(root *TreeNode) {
	if root==nil{
		return
	}
	temp:=root.Right
	root.Right=root.Left
	root.Left=temp
	invertTree2(root.Left)
	invertTree2(root.Right)
}




func distributeCandies(candyType []int) int {
	set:=make(map[int]struct{})
	for i:=0;i< len(candyType);i++{
		set[candyType[i]]= struct{}{}
	}
	if len(set)>= len(candyType)/2{
		return len(candyType)/2
	}
	return len(set)
}

func findErrorNums(nums []int) []int {
	var result =make([]int,2)
	var res=make(map[int]int)
	for i:=0;i< len(nums);i++{
		res[nums[i]]++
	}
	for i:=1;i<=len(nums);i++{
		if res[i]==2{
			result[0]=i
		}
		if res[i]==0{
			result[1]=i
		}
	}
	return result
}



func largestPerimeter(A []int) int {

	sort.Slice(A, func(i, j int) bool {
		return A[i]>A[j]
	})
	fmt.Println(A)
	for i:=0;i< len(A)-2;i++{
		if A[i+2]+A[i+1]>A[i]{
			return A[i]+A[i+1]+A[i+2]
		}
	}
	return 0
}

//"3876620623801494171"
//"6529364523802684779"