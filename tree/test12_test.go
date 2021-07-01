package tree

import (
	"container/list"
	"fmt"
	"testing"
)

//Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func Test_13(t *testing.T) {
	root:=&Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	n := levelOrderN(root)
	fmt.Println(n)
}

func levelOrderN(root *Node) [][]int {
	if root==nil{
		return [][]int{}
	}
	var result = make([][]int, 0)
	//队列
	var queue=list.New()
	//存入根节点
	queue.PushFront(root)
	for queue.Len()>0{
		var curNum =queue.Len()
		var temp =make([]int,0,curNum)
		//取出节点
		for i:=0;i<curNum;i++{
			//拿出节点存入结果集中
			node := queue.Back()
			treeNode := node.Value.(*Node)
			temp=append(temp,treeNode.Val)
			//移除这个节点
			queue.Remove(node)
			//遍历子节点存入队列中
			for i:=0;i< len(treeNode.Children);i++{
				queue.PushFront(treeNode.Children[i])
			}
		}
		result=append(result,temp)
	}
	return result
}


