package tree

import (
	"container/list"
	"fmt"
	"testing"
)

//      1
//   2    10
// 5    11  12
func Test_11(t *testing.T){
	root:=&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 10,
			Left: &TreeNode{Val: 11},
			Right: &TreeNode{Val: 12},
		},
	}
	fmt.Println(treePath3(root))

}


func treePath1(root *TreeNode)[]int{
	var result =make([]int,0)
	var stack=list.New()
	for root!=nil||stack.Len()>0{
		if root!=nil{
			//选取根节点的值
			result=append(result,root.Val)
			//存入栈中
			stack.PushFront(root)
			root=root.Left
		}else{
			//在遍历右节点
			node := stack.Front().Value.(*TreeNode)
			stack.Remove(stack.Front())
			root=node.Right
		}
	}
	return result
}

func treePath2(root *TreeNode)[]int{
	var result =make([]int,0)
	var stack=list.New()
	for root!=nil||stack.Len()>0{
		if root!=nil{
			//放入栈中
			stack.PushFront(root)
			//遍历左节点
			root=root.Left
		}else{
			//弹出节点
			//在遍历右节点
			node := stack.Front().Value.(*TreeNode)
			stack.Remove(stack.Front())
			//存入根节点值
			result=append(result,node.Val)
			//遍历右字数
			root=node.Right
		}
	}
	return result
}

func treePath3(root *TreeNode)[]int{
	var result =make([]int,0)
	var stack=list.New()
	for root!=nil||stack.Len()>0{
		if root!=nil{
			//放入栈中
			stack.PushFront(root)
			//遍历左节点
			//存入根节点值
			result=append([]int{root.Val},result...)
			root=root.Right
		}else{
			//弹出节点
			//在遍历右节点
			node := stack.Front().Value.(*TreeNode)
			stack.Remove(stack.Front())
			//遍历右字数
			root=node.Left

		}
	}
	return result
}


