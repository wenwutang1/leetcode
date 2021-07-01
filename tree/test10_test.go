package tree

import (
	"fmt"
	"testing"
)

//      1
//   2    10
// 5    11  12
func Test_10(t *testing.T){

	root:=&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 10,
			Left: &TreeNode{Val: 11},
			Right: &TreeNode{Val: 12},
		},
	}
	fmt.Println(postorderTraversal(root))

}

//后序遍历方式
func postorderTraversal(root *TreeNode) []int {

	var result []int=make([]int,0)
	postorderTraversal2(root,&result)
	return result
}

func postorderTraversal2(root *TreeNode,result *[]int)  {

	if root==nil{
		return
	}
	postorderTraversal2(root.Left,result)
	postorderTraversal2(root.Right,result)
	*result=append(*result,root.Val)
}