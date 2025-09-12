package lc

//你可以选择使用单链表或者双链表，设计并实现自己的链表。
//
//单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。
//
//如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。
//
//实现 MyLinkedList 类：
//
//MyLinkedList() 初始化 MyLinkedList 对象。
//int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
//void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
//void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
//void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
//void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。

// ai解法
type MyLinkedList struct {
	dummy *Node // 哨兵节点
	size  int   // 链表大小
}

type Node struct {
	val  int
	next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummy: &Node{}, // 创建哨兵节点
		size:  0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	// 索引无效
	if index < 0 || index >= this.size {
		return -1
	}

	// 从哨兵节点开始遍历
	cur := this.dummy
	for i := 0; i <= index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// 如果索引大于链表长度，不插入节点
	if index > this.size {
		return
	}

	// 索引为负数，视为插入头部
	if index < 0 {
		index = 0
	}

	// 找到插入位置的前一个节点
	pred := this.dummy
	for i := 0; i < index; i++ {
		pred = pred.next
	}

	// 创建新节点并插入
	newNode := &Node{
		val:  val,
		next: pred.next,
	}
	pred.next = newNode
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	// 索引无效
	if index < 0 || index >= this.size {
		return
	}

	// 找到要删除节点的前一个节点
	pred := this.dummy
	for i := 0; i < index; i++ {
		pred = pred.next
	}

	// 删除节点
	pred.next = pred.next.next
	this.size--
}

// 官方解法
type MyLinkedList1 struct {
	head *ListNode
	size int
}

func Constructor1() MyLinkedList1 {
	return MyLinkedList1{&ListNode{}, 0}
}

func (l *MyLinkedList1) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	cur := l.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (l *MyLinkedList1) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList1) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList1) AddAtIndex(index, val int) {
	if index > l.size {
		return
	}
	index = max(index, 0)
	l.size++
	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	toAdd := &ListNode{val, pred.Next}
	pred.Next = toAdd
}

func (l *MyLinkedList1) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	l.size--
	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	pred.Next = pred.Next.Next
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 官方题解 双向链表
type node struct {
	val        int
	next, prev *node
}

type MyLinkedList2 struct {
	head, tail *node
	size       int
}

func Constructor2() MyLinkedList2 {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head
	return MyLinkedList2{head, tail, 0}
}

func (l *MyLinkedList2) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	var curr *node
	if index+1 < l.size-index {
		curr = l.head
		for i := 0; i <= index; i++ {
			curr = curr.next
		}
	} else {
		curr = l.tail
		for i := 0; i < l.size-index; i++ {
			curr = curr.prev
		}
	}
	return curr.val
}

func (l *MyLinkedList2) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList2) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList2) AddAtIndex(index, val int) {
	if index > l.size {
		return
	}
	index = max(0, index)
	var pred, succ *node
	if index < l.size-index {
		pred = l.head
		for i := 0; i < index; i++ {
			pred = pred.next
		}
		succ = pred.next
	} else {
		succ = l.tail
		for i := 0; i < l.size-index; i++ {
			succ = succ.prev
		}
		pred = succ.prev
	}
	l.size++
	toAdd := &node{val, succ, pred}
	pred.next = toAdd
	succ.prev = toAdd
}

func (l *MyLinkedList2) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	var pred, succ *node
	if index < l.size-index {
		pred = l.head
		for i := 0; i < index; i++ {
			pred = pred.next
		}
		succ = pred.next.next
	} else {
		succ = l.tail
		for i := 0; i < l.size-index-1; i++ {
			succ = succ.prev
		}
		pred = succ.prev.prev
	}
	l.size--
	pred.next = succ
	succ.prev = pred
}
