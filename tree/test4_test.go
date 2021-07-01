package tree

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_04(t *testing.T){
	var root *TreeNode
	fmt.Println(levelOrder(root))
}

//    3
//   / \
//  9  20
//    /  \
//   15   7
//return [
//  [3],
//  [9,20],
//  [15,7]
//]
func levelOrder(root *TreeNode) [][]int {
	if root==nil{
		return [][]int{}
	}
	var result=make([][]int,0)
	//队列
	var queue =list.New()
	//根入队列
	queue.PushBack(root)
	for queue.Len()>0{
		nodeNum:=queue.Len()
		//拿出队列中的所有的节点存入结果集中
		var temp=make([]int,0,nodeNum)
		for i:=0;i<nodeNum;i++{
			node := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())
			temp=append(temp,node.Val)
			if node.Left!=nil{
				queue.PushBack(node.Left)
			}
			if node.Right!=nil{
				queue.PushBack(node.Right)
			}
		}
		result=append(result,temp)
	}
	return result
}


