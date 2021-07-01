package tree

import (
	"fmt"
	"testing"
)

func Test_08(t *testing.T){

	root:=&TreeNode{
		Val:3,
		Right: &TreeNode{
			Val:3,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{Val: 5},
			},
		},
	}

	depth := minDepth(root)
	fmt.Println(depth)
}

//给定一个二叉树，找出其最小深度。
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//说明：叶子节点是指没有子节点的节点
func minDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	if root.Left==nil&&root.Right==nil{
		return 1
	}
	left,right:=0,0
	if root.Left!=nil{
		left=minDepth(root.Left)+1
	}
	if root.Right!=nil{
		right=minDepth(root.Right)+1
	}
	return mintree(left,right)
}


func mintree(a,b int)int{
	if a<b{
		return a
	}
	return b
}
