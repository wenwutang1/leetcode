package _021

import (
	"fmt"
	"testing"
)

func Test44(t *testing.T){

	var root =&TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{Val: 21},
			Right: &TreeNode{
				Val: 20,
				Right: &TreeNode{Val: 888},
			},
		},
	}
	fmt.Println(countTree(root))
}

func maxDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	return 1+max123(maxDepth(root.Left),maxDepth(root.Right))
}

func max123(x,y int)int{
	if x>y{
		return x
	}
	return y
}

func countTree(root *TreeNode)int{
	if root==nil{
		return 0
	}
	return 1+countTree(root.Left)+countTree(root.Right)
}
