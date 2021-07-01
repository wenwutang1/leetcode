package tree

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_05(t *testing.T){
	root:=&TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{Val: 15},

		},
	}
	bottom := levelOrderBottom(root)
	fmt.Println(bottom)
}

//    3
//   / \
//  9  20
//    /  \
//   15   7
//return
//[
//  [15,7],
//  [9,20],
//  [3]
//]
func levelOrderBottom(root *TreeNode) [][]int {
	if root==nil{
		return [][]int{}
	}
	var result =make([][]int,0)
	//队列
	var queue =list.New()
	//存入根节点
	queue.PushBack(root)
	for queue.Len()>0{
		number:=queue.Len()
		var temp=make([]int,0,number)
		for i:=0;i<number;i++{
			//拿出当前层级的所有的节点
			node:=queue.Front()
			queue.Remove(node)
			treeNode:=node.Value.(*TreeNode)
			if treeNode.Left!=nil{
				queue.PushBack(treeNode.Left)
			}
			if treeNode.Right!=nil{
				queue.PushBack(treeNode.Right)
			}
			temp=append(temp,treeNode.Val)
		}
		result=append(result,temp)
	}
	//倒序result
	i,j:=0, len(result)-1
	for i<j{
		result[i],result[j]=result[j],result[i]
		i++
		j--
	}
	return result
}

