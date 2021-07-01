package _021

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func Test45(t *testing.T){

	hash := md5.New()
	hash.Write([]byte("666666"))
	fmt.Println(hex.EncodeToString(hash.Sum(nil)))

	hash.Write([]byte("dadadadadaddadadadaa"))
	fmt.Println(len(hex.EncodeToString(hash.Sum(nil))))

}

func swapPairs(head *ListNode) *ListNode {

	//head==nil或者是最后一个不能交换了
	if head==nil||head.Next==nil{
		return head
	}
	var one =head
	var two =head.Next
	var three =two.Next
	//交换第一个节点和第二个节点
	two.Next=one
	one.Next=swapPairs(three)
	return two
}

