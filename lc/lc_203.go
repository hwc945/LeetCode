package lc

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	for head.Val == val {
		head = head.Next
		if head == nil {
			return head
		}
	}
	next := head.Next
	pre := head
	for next != nil {
		if next.Val == val {
			pre.Next = next.Next
			next = next.Next
		} else {
			pre = next
			next = next.Next
		}
	}
	return head
}

// AI解法
func removeElements1(head *ListNode, val int) *ListNode {
	// 创建哨兵节点简化头节点处理
	dummy := &ListNode{Next: head}
	current := dummy

	for current.Next != nil {
		if current.Next.Val == val {
			// 删除匹配节点
			current.Next = current.Next.Next
		} else {
			// 移动到下一个节点
			current = current.Next
		}
	}

	return dummy.Next
}

// 递归解法
func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
