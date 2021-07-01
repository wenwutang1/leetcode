package _021

import (
	"fmt"
	"testing"
)

func Test13(t *testing.T) {

	root:=&TreeNode{
		Val: 2,
		Left: &TreeNode{Val: 1},
		Right:  &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(isValidBST(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Integer struct {
	value int
	flag  bool
}


//第一种简单的判断方法
//递归判断当前节点的值和其左右节点的关系是不正确的
//   5
// 4   6
//    2  7
//列如这个数2 这个节点满足当前关系但是不满足比5大
//     5
//  3    8
// 1 4  7  9
//
//观察这个树 ,左子树的的取值范围在[-00,当前节点),右子树的取值范围在(当前节点,+00)
//
func isValidBST(root *TreeNode) bool {
	return isValidBST2(root,nil,nil)
}

func isValidBST2(root *TreeNode,low *Integer,high *Integer)bool{
	//空子树也是一个搜索树
	if root==nil{
		return true
	}
	if (low!=nil &&root.Val<=low.value) ||(high!=nil&&root.Val>=high.value) {
		return false
	}
	return isValidBST2(root.Left,low,&Integer{value: root.Val}) &&isValidBST2(root.Right,&Integer{value: root.Val},high)
}



