package _021

import (
	"fmt"
	"testing"
)

func Test42(t *testing.T){
	p:=&TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	q:=&TreeNode{
		Val:  1,
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isSameTree(p,q))
}


func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p==nil||q==nil{
		return p==q
	}

	if p.Val!=q.Val{
		return false
	}

	return isSameTree(p.Left,q.Left)&&isSameTree(p.Right,q.Right)
}
