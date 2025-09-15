package lc

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

// AI解法,官方
func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := first.Next

		// 交换节点
		prev.Next = second
		first.Next = second.Next
		second.Next = first

		// 移动指针
		prev = first
	}

	return dummy.Next
}
