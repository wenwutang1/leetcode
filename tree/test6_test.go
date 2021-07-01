package tree

import (
	"fmt"
	"testing"
)

//     			5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1
//
func Test_06(t *testing.T){
	root:=&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1},
			},
		},

	}
	sum := pathSum(root,22)
	fmt.Println(sum)
}


func pathSum(root *TreeNode, sum int) [][]int {
	result:=make([][]int,0)
	dfs(root,sum,[]int{},0,&result)
	return result

}
func dfs(root *TreeNode, sum int,road []int,roadSum int,result *[][]int) {
	if root==nil{
		return
	}
	road=append(road,root.Val)
	roadSum+=root.Val
	if root.Left==nil&&root.Right==nil{
		//判断路径是否=目标值
		//拷贝数组,不然数组会被后面的路径污染
		temp:=make([]int, len(road))
		copy(temp,road)
		if roadSum==sum{
			*result=append(*result,temp)
		}
		return
	}
	dfs(root.Left,sum,road,roadSum,result)
	dfs(root.Right,sum,road,roadSum,result)
	//返回上层节点的时候去掉这个节点的值
	road=road[:len(road)-1]
	roadSum-=root.Val
}

