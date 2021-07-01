package _021

import (
	"container/list"
	"fmt"
	"testing"
)

func Test40(t *testing.T) {

	root:=&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 10},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: nil,
		},
	}

	result := inorderTraversal3(root)
	fmt.Println(result)
}


func inorderTraversal(root *TreeNode) []int {
	type f  func(*TreeNode)
	var f1 f
	var result=make([]int,0)
	f1= func(node *TreeNode) {
		if node==nil{
			return
		}
		f1(node.Left)
		result=append(result,node.Val)
		f1(node.Right)
	}
	f1(root)
	return result
}

func inorderTraversal2(root *TreeNode)[]int{
	if root==nil{
		return []int{}
	}
	var result =make([]int,0)
	if root.Left!=nil{
		result=append(result,inorderTraversal2(root.Left)...)
	}
	result=append(result,root.Val)
	if root.Right!=nil{
		result=append(result,inorderTraversal2(root.Right)...)
	}
	return result
}

func inorderTraversal3(root *TreeNode)[]int{
	if root==nil{
		return []int{}
	}
	var result =make([]int,0)
	var cur =root
	var stack=list.New()
	for cur!=nil||stack.Len()>0{

		if cur!=nil{
			stack.PushFront(cur)
			cur=cur.Left
		}else{
			front := stack.Front()
			node := front.Value.(*TreeNode)
			stack.Remove(front)
			result=append(result,node.Val)
			cur=node.Right
		}

	}

	return result
}
