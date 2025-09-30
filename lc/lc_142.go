package lc

// 检测链表中的入环节点，使用弗洛伊德快慢指针确保 O(1) 额外空间。
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return head
		}
	}
	return nil
}

// AI 快慢指针
func detectCycleAI(head *ListNode) *ListNode {
	return detectCycle(head)
}

// 官方解答
func detectCycle1(head *ListNode) *ListNode {
	return detectCycle(head)
}
