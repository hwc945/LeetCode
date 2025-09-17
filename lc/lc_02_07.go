package lc

// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, la := headA, headA
	b, lb := headB, headB
	lengthA, lengthB := 1, 1
	for la.Next != nil {
		la = la.Next
		lengthA++
	}
	for lb.Next != nil {
		lb = lb.Next
		lengthB++
	}
	if la != lb {
		return nil
	}
	if lengthA > lengthB {
		for i := 0; i < lengthA-lengthB; i++ {
			a = a.Next
		}
	} else {
		for i := 0; i < lengthB-lengthA; i++ {
			b = b.Next
		}
	}
	for a != b {
		a = a.Next
		b = b.Next
	}
	return a
}

// AI 优化
func getIntersectionNodeAI(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB
	lenA, lenB := 1, 1

	// 计算链表A的长度和尾节点
	tailA := a
	for tailA.Next != nil {
		tailA = tailA.Next
		lenA++
	}

	// 计算链表B的长度和尾节点
	tailB := b
	for tailB.Next != nil {
		tailB = tailB.Next
		lenB++
	}

	// 如果尾节点不同，说明不相交
	if tailA != tailB {
		return nil
	}

	// 对齐起点
	for lenA > lenB {
		a = a.Next
		lenA--
	}
	for lenB > lenA {
		b = b.Next
		lenB--
	}

	// 同步前进直到找到交点
	for a != b {
		a = a.Next
		b = b.Next
	}

	return a
}

// 官方解答
// 哈希集合
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

// 双指针
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pointerA, pointerB := headA, headB

	// 遍历两个链表，直到找到交点或同时到达末尾
	for pointerA != pointerB {
		// 如果pointerA到达链表末尾，切换到链表B
		if pointerA == nil {
			pointerA = headB
		} else {
			pointerA = pointerA.Next
		}

		// 如果pointerB到达链表末尾，切换到链表A
		if pointerB == nil {
			pointerB = headA
		} else {
			pointerB = pointerB.Next
		}
	}

	// 返回交点或nil（如果不相交）
	return pointerA
}
