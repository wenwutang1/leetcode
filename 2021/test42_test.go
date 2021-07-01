package _021

import (
	"fmt"
	"testing"
)

func Test43(t *testing.T){

	root:=&TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val: 2,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{Val: 888},
			Right: &TreeNode{Val: 999},
		},
	}
	fmt.Println(gfs(root))
}

func isSymmetric(root *TreeNode) bool {
	type fn func(node1 *TreeNode,node2 *TreeNode)bool
	var f1 fn
	f1= func(node1 *TreeNode, node2 *TreeNode) bool{
		if node1==nil&&node2==nil{
			return true
		}
		if node1==nil||node2==nil{
			return false
		}
		return node1.Val==node2.Val&&f1(node1.Left,node2.Right)&&f1(node1.Right,node2.Left)
	}
	return f1(root,root)
}

func isSymmetric2(root *TreeNode) bool {

	if root==nil{
		return true
	}
	queue:=make([]*TreeNode,0)
	queue=append(queue, root)
	queue= append(queue, root)
	for len(queue)>0{
		node1 := queue[0]
		node2 := queue[1]
		queue=queue[2:]
		if node2==nil&&node1==nil{
			continue
		}
		if node1==nil||node2==nil{
			return false
		}
		if node1.Val!=node2.Val{
			return false
		}
		queue=append(queue,node1.Left)
		queue=append(queue,node2.Right)

		queue=append(queue,node1.Right)
		queue=append(queue,node2.Left)
	}
	return true
}











func gfs(root *TreeNode)[]int{
	var result =make([]int,0)
	if root==nil{
		return result
	}
	queue:=make([]*TreeNode,0)
	queue=append(queue,root)
	for len(queue)>0{
		temp:=make([]*TreeNode,0, len(queue))
		for i:=0;i< len(queue);i++{
			result=append(result,queue[i].Val)
			if queue[i].Left!=nil{
				temp=append(temp,queue[i].Left)
			}
			if queue[i].Right!=nil{
				temp=append(temp,queue[i].Right)
			}
		}
		queue=temp
	}
	return result
}