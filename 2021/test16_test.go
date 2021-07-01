package _021

import (
	"fmt"
	"math"
	"testing"
)

//

func Test16(t *testing.T){
	root:=&TreeNode{
		  Val: 4,
		  Left: &TreeNode{
		  	Val:2,
		  	Left: &TreeNode{Val: 1},
		  	Right: &TreeNode{Val: 3},
		  },
		  Right: &TreeNode{
		  	Val: 6,
		  },
	}
	fmt.Println(minDiffInBST(root))
}

//给定一个二叉搜索树的根节点 root,
//返回树中任意两节点的差的最小值
//          4
//        /   \
//      2      6
//     / \
//    1   3
//思路任意节点的值的最小值一定是相邻两个节点的最小值
//因为二叉搜素树转换过来就是一个升序的数组
//使用中序遍历，每次跟新最小值
var pre,minn int=-1,math.MaxInt32
func minDiffInBST(root *TreeNode) int {
	minDiffInBST2(root)
	return minn
}

func minDiffInBST2(root *TreeNode){
	if root==nil{
		return
	}
	minDiffInBST2(root.Left)
	if pre!=-1&&minn>(root.Val-pre){
		minn=root.Val-pre
	}
	pre=root.Val
	minDiffInBST2(root.Right)
}

