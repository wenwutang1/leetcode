package _021

import (
	"fmt"
	"testing"
)

func Test17(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 10,
			Right: &TreeNode{
				Val: 50,
			},
		},
		Right: &TreeNode{
			Val: 80,
			Left: &TreeNode{
				Val: 88,
			},
			Right: &TreeNode{
				Val: 99,
				Left: &TreeNode{
					Val: 111,
				},
			},
		},
	}
	fmt.Println(rightSideView2(root))
}

//给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//输入:[1,2,3,null,5,null,4]
//输出:[1, 3, 4]
//解释:
//
//   1            <---
// /   \
//2     3         <---
// \     \
//  5     4       <---
//最先想到的解法,广度优先遍历,每一层最后的节点，便是二叉树的右视图
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var queue = make([]*TreeNode, 0, 1)
	queue = append(queue, root)
	var result = make([]int, 0)
	for len(queue) > 0 {
		tm := make([]*TreeNode, 0)
		result = append(result, queue[len(queue)-1].Val)
		for i := range queue {
			if queue[i].Left != nil {
				tm = append(tm, queue[i].Left)
			}
			if queue[i].Right != nil {
				tm = append(tm, queue[i].Right)
			}
		}

		queue = tm
	}
	return result
}


//二叉树的右视图
func rightSideView2(root *TreeNode) []int {
	type  DFS func(node *TreeNode,level int)
	var   result=make([]int,0)
	var f DFS
	f= func(node *TreeNode, level int) {
		if node==nil{
			return
		}
		if level== len(result){
			result=append(result,node.Val)
		}
		level++
		f(node.Left,level)
		f(node.Right,level)
	}
	f(root,0)
	return result
}
