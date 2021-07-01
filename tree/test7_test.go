package tree

import (
	"fmt"
	"testing"
)

func Test_07(t *testing.T){
	root:=&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{Val: 6},
		},
	}
	flatten2(root)
	fmt.Println("芜湖")
}

//    1
//   / \
//  2   5
// / \   \
//3   4   6
//将其展开为：
//
//1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6
//
//第一种方法,先前序遍历在组成结构
func flatten(root *TreeNode)  {
	list:=make([]*TreeNode,0)
	front(root,&list)
	for i:=1;i< len(list);i++{
		list[i-1].Right=list[i]
		list[i-1].Left=nil
	}
}

func front(root *TreeNode,list *[]*TreeNode){
	if root==nil{
		return
	}
	*list=append(*list,root)
	front(root.Left,list)
	front(root.Right,list)
}

//第二种方法,第一种中序遍历的方式分成两步是因为
//同时设置节点会导致右节点丢失
//在遍历的同时进行结构的拆分
func flatten2(root *TreeNode)  {
	if root==nil{
		return
	}
	var preNode *TreeNode
	//使用一个栈来保存右节点,因为栈是先进后出
	//所以是先放入右节点在放入左节点
	stack :=make([]*TreeNode,0)
	stack=append(stack,root)
	for len(stack)>0{
		var cur =stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		if preNode!=nil{
			preNode.Right=cur
			preNode.Left=nil
		}
		if cur.Right!=nil{
			stack=append(stack,cur.Right)
		}
		if cur.Left!=nil{
			stack=append(stack,cur.Left)
		}

		preNode=cur
	}
}