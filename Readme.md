# 代码随想录

每日锻炼一些算法，现在主要来源于 LeetCode。

## 目录

- [数组](#数组)
- [链表](#链表)

## 数组

| 题目 | 完成日期 | 要点 | 题解 |
| :--- | :--- | :--- | :--- |
| [704. 二分查找](https://leetcode.cn/problems/binary-search/) | 2025-09-04 | <ul><li>注意「左闭右开」`[left, right)` 和「左闭右闭」`[left, right]` 区间的区别</li><li>`mid` 的取值方式 `left + (right - left) / 2` 可以有效防止整数溢出</li><li>需要根据区间的定义，正确地更新 `left` 和 `right` 的位置</li><li>对于已排序数组，二分查找已是理论最优解 O(log n)</li></ul> | [Go](./lc/lc_704.go) |
| [27. 移除元素](https://leetcode.cn/problems/remove-element/) | 2025-09-05 | <ul><li>注意算法思想和语言能力的区别。</li><li>首位指针法，left代表下一个有效位置，right代表当前可替换的位置，直接覆盖而非交换，减少赋值次数，效率最佳</li><li>快慢指针法，快指针遍历数组，慢指针记录新数组的位置，保持原数组顺序</li></ul> | [Go](./lc/lc_27.go) |
| [977. 有序数组的平方](https://leetcode.cn/problems/squares-of-a-sorted-array/description/) | 2025-09-08 | <ul><li>双指针，寻找平方最大值，记录在答案数组最后</li></ul> | [Go](./lc/lc_977.go) |
| [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/description/) | 2025-09-09 | <ul><li>滑动窗口，维护一个窗口的和，当和大于等于目标值时，尝试收缩左边界，更新最小长度</li></ul> | [Go](./lc/lc_209.go) |
| [59. 螺旋矩阵 II](https://leetcode.cn/problems/spiral-matrix-ii/description/) | 2025-09-09 | <ul><li>模拟螺旋过程：按照顺时针方向（右→下→左→上）填充矩阵。</li><li>维护边界：使用 left, right, top, bottom四个变量控制填充范围。</li><li>边界收缩：每次填充完一圈后，收缩边界以避免重复填充。</li><li>终止条件：当填充的数字达到 n²时结束。</li></ul> | [Go](./lc/lc_59.go) |

## 链表

| 题目                                                                                             | 完成日期           | 要点                                                                                                                     | 题解                     |
|:-----------------------------------------------------------------------------------------------|:---------------|:-----------------------------------------------------------------------------------------------------------------------|:-----------------------|
| [203. 移除链表元素](https://leetcode.cn/problems/remove-linked-list-elements/description/)           | 2025-09-10     | <ul><li>使用 pre和 next指针进行链表遍历和删除操作</li><li>头结点特殊情况可以增加哨兵节点处理</li><li>可以使用递归方法简化代码</li></ul>                             | [Go](./lc/lc_203.go)   |
| [707. 设计链表](https://leetcode.cn/problems/design-linked-list/description/)                      | 2025-09-12（失败） | <ul><li>使用哨兵节点：简化头节点操作</li><li>维护链表长度：便于索引验证</li><li>正确修改链表头：通过返回新链表头或使用二级指针</li><li>统一空链表处理：确保所有方法都能正确处理空链表</li></ul> | [Go](./lc/lc_707.go)   |
| [24. 两两交换链表中的节点](https://leetcode.cn/problems/swap-nodes-in-pairs/description/)                | 2025-09-12     | <ul><li>使用哨兵节点：简化头节点操作</li><li>改用迭代方法：避免递归的栈溢出风险，降低空间复杂度</li><li>更清晰的指针操作：明确展示节点交换过程</li></ul>                         | [Go](./lc/lc_24.go)    |
| [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/description/)                     | 2025-09-12     | <ul><li>迭代方法：避免递归的栈溢出风险，降低空间复杂度</li><li>递归算法：只负责深入到链表末尾，在递归返回时执行指针反转操作，始终保持返回的是原链表的尾节点（新头节点）</li></ul>                 | [Go](./lc/lc_206.go)   |
| [19. 删除链表的倒数第N个节点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/) | 2025-09-16     | <ul><li>迭代方法：求出链表长度，找到倒数第n个节点</li><li>双指针算法：快指针指向结尾，慢指针为快指针前第n个节点</li><li>栈算法：所有节点存入栈，直接找到倒数第n个节点删除</li></ul>          | [Go](./lc/lc_19.go)    |
| [02_07.链表相交](https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/)              | 2025-09-17     | <ul><li>哈希方法：通过哈希值找到第一个相交节点</li><li>双指针算法：分头两次遍历两个链表，找到相交节点</li><li>常规算法：先判断是否相交，若相交，通过长度差确定不想交段落长度差找到相交节点</li></ul>   | [Go](./lc/lc_02_07.go) |
| [142. 环形链表II](https://leetcode.cn/problems/linked-list-cycle-ii/)                              | 2025-09-17     | <ul><li>哈希方法：通过哈希值找到第一个重复节点</li><li>双指针算法：快指针追上慢指针个，数学计算剩余到环节点距离</li></ul>                                             | [Go](./lc/lc_142.go)   |
