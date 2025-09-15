package lc

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{}
	current := head
	for current != nil {
		next := current.Next
		current.Next = newHead.Next
		newHead.Next = current
		current = next
	}
	return newHead.Next
}

// 官方递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
