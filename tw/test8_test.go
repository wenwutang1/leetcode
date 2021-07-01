package tw

type ListNode struct {
	Val  int
	Next *ListNode
}

//给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
//请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	h1 := head
	t2 := h1.Next
	h2 := t2
	for h2 != nil && h2.Next != nil {
		h1.Next = h1.Next.Next
		h1=h1.Next
		h2.Next = h2.Next.Next
		h2 = h2.Next
	}

	h1.Next = t2
	return head
}
