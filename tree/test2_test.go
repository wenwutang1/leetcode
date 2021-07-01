package tree

import (
	"fmt"
	"testing"
)

func Test_02(t *testing.T) {
	node:=&TreeNode{Val: 3}
	node.Left=&TreeNode{
		Val: 4,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{Val:2},
	}
	node.Right=&TreeNode{Val: 5}

	node2:=&TreeNode{
		Val: 4,
		Left: &TreeNode{Val: 1},
		Right: nil,
	}
	fmt.Println(isSubStructure(node,node2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


//判断B是不是A的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {

	if A==nil||B==nil{
		return false
	}
	return isSubStructure2(A,B)||isSubStructure(A.Left,B)||isSubStructure(A.Right,B)

}

func isSubStructure2(A *TreeNode, B *TreeNode)  bool{
	if B==nil{
		return true
	}
	if A==nil||A.Val!=B.Val{
		return false
	}
	return isSubStructure2(A.Left,B.Left)&&isSubStructure2(A.Right,B.Right)
}