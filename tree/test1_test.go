package tree

import "testing"

//Definition for a binary tree node
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}


func Test_01(t *testing.T) {

}

func mirrorTree(root *TreeNode) *TreeNode {
	if root==nil{
		return root
	}
	mirrorTree2(root)
	return root
}


func mirrorTree2(root *TreeNode){
	if root==nil{
		return
	}
	//左子树交换
	root.Left,root.Right=root.Right,root.Left
	mirrorTree2(root.Left)
	mirrorTree2(root.Right)
}



