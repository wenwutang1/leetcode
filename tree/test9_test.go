package tree

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_09(t *testing.T) {

	var root =&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(binaryTreePaths(root))
}

//二叉树的所有的路径
func binaryTreePaths(root *TreeNode) []string {
	if root==nil{
		return []string{}
	}
	var result =make([]string,0)
	RoadTree(root,"",&result)
	return result
}


func RoadTree(root *TreeNode,road string,res *[]string){
	if root==nil{
		return
	}
	if root.Left!=nil||root.Right!=nil{
		road+=strconv.Itoa(root.Val)+"->"
	}else if root.Left==nil&&root.Right==nil{
		road+=strconv.Itoa(root.Val)
		*res=append(*res,road)
	}
	RoadTree(root.Left,road,res)
	RoadTree(root.Right,road,res)
}