package lc

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	count := 0
	for current != nil {
		count++
		current = current.Next
	}
	if count == n {
		return head.Next
	}
	current = head
	position := count - n
	for i := 1; i < position; i++ {
		current = current.Next

	}
	current.Next = current.Next.Next
	return head
}

// 双指针一次遍历
func removeNthFromEndPlus(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}

// 栈 官方
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}
