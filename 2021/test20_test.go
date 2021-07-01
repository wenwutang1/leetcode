package _021

import (
	"fmt"
	"testing"
)

func Test20(t *testing.T){

	root:=&TreeNode{

		Val: 1,
		Left: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 25,
			},
			Right: &TreeNode{
				Val: 35,
			},
		},
		Right: &TreeNode{
			Val: 80,
			Left: &TreeNode{Val: 100},
			Right: &TreeNode{Val:200},
		},
	}

	order := zigzagLevelOrder(root)
	fmt.Println(order)
}


//二叉树的锯齿形便利
//即每一层先从左到右便利，下一层从右到左便利
func zigzagLevelOrder(root *TreeNode) [][]int {

	if root==nil{
		return [][]int{}
	}
	queue:=make([]*TreeNode,0)
	queue=append(queue,root)
	var result =make([][]int,0)
	var flag =true
	for len(queue)>0{
		tm:=make([]int,0, len(queue))
		if flag{
			for i:=range queue{
				tm=append(tm,queue[i].Val)
			}
		}else{
			for j:= len(queue)-1;j>-1;j--{
				tm=append(tm,queue[j].Val)
			}
		}
		result=append(result,tm)
		flag=!flag
		link:=make([]*TreeNode,0, len(queue)*2)
		//广度优先
		for i:=0;i< len(queue);i++{
			if queue[i].Left!=nil{
				link=append(link,queue[i].Left)
			}
			if queue[i].Right!=nil{
				link=append(link,queue[i].Right)
			}
		}
		queue=link
	}
	return result
}
