package tree

import (
	"fmt"
	"testing"
)

func Test_03(t *testing.T){

	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(tree)
}


func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)<1{
		return nil
	}
	//左子树的长度
	leftLen:=0
	for k,v:=range inorder{
		leftLen=k
		if v==preorder[0]{
			break
		}
	}
	//创建根节点
	root := new(TreeNode)
	root.Val=preorder[0]
	root.Left=buildTree(preorder[1:leftLen+1],inorder[:leftLen])
	root.Right=buildTree(preorder[leftLen+1:],inorder[leftLen+1:])
	return root
}




