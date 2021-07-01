package _021

import (
	"testing"
)

func Test36(t *testing.T) {

	//head:=&ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 1,
	//		Next: &ListNode{
	//			Val: 2,
	//			Next: &ListNode{
	//				Val: 2,
	//			},
	//		},
	//	},
	//}
	//
	//node := deleteDuplicates3(head)
	//
	//for node!=nil{
	//	fmt.Println(node)
	//	node=node.Next
	//}
}

/**
 * Definition for singly-linked list
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//{ 1,1,2,2,3,5,6 }
func deleteDuplicates(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	var slow =head
	var quick =head.Next
	for quick!=nil{
		if quick.Val!=slow.Val{
			slow.Next=quick
			slow=quick
		}
		quick=quick.Next
	}
	slow.Next=quick
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	var cur =head
	for cur!=nil&&cur.Next!=nil{
		if cur.Val==cur.Next.Val{
			cur.Next=cur.Next.Next
		}else{
			cur=cur.Next
		}
	}
	return head
}

