# LeetCode in Go

个人刷题仓库，主要使用 Go 语言实现 LeetCode 题目并记录思路。代码按题号归档在 `./lc` 目录中，`main.go` 可用于快速运行或调试单道题目。

## 优化亮点

- **27. 移除元素**（`lc/lc_27.go`）：统一为双指针原地覆盖，避免额外切片分配，降低内存压力。
- **28. 找出字符串中第一个匹配项**（`lc/lc_28.go`）：使用 KMP 的前缀函数表，主串扫描保持 O(n) 时间复杂度。
- **142. 环形链表 II**（`lc/lc_142.go`）：采用弗洛伊德快慢指针寻找入环节点，将空间复杂度优化为 O(1)。
- **209. 长度最小的子数组**（`lc/lc_209.go`）：精简滑动窗口逻辑，实时压缩窗口更新答案，保持 O(n) 时间。

## 题目索引

### 数组

| 题目 | 核心技巧 | 源码 |
| :--- | :--- | :--- |
| [27. 移除元素](https://leetcode.cn/problems/remove-element/) | 原地双指针删除 | [Go](./lc/lc_27.go) |
| [59. 螺旋矩阵 II](https://leetcode.cn/problems/spiral-matrix-ii/) | 按层模拟顺时针填充 | [Go](./lc/lc_59.go) |
| [70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/) | 动态规划 + 状态压缩 | [Go](./lc/lc_70.go) |
| [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/) | 滑动窗口 + 动态收缩 | [Go](./lc/lc_209.go) |
| [704. 二分查找](https://leetcode.cn/problems/binary-search/) | 左闭右闭二分 | [Go](./lc/lc_704.go) |
| [977. 有序数组的平方](https://leetcode.cn/problems/squares-of-a-sorted-array/) | 双指针从两端收集平方 | [Go](./lc/lc_977.go) |

### 链表

| 题目 | 核心技巧 | 源码 |
| :--- | :--- | :--- |
| [19. 删除链表的倒数第 N 个节点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/) | 快慢指针定位目标 | [Go](./lc/lc_19.go) |
| [24. 两两交换链表中的节点](https://leetcode.cn/problems/swap-nodes-in-pairs/) | 哨兵 + 原地指针交换 | [Go](./lc/lc_24.go) |
| [142. 环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/) | 弗洛伊德判环定位入环点 | [Go](./lc/lc_142.go) |
| [203. 移除链表元素](https://leetcode.cn/problems/remove-linked-list-elements/) | 哨兵节点细分删除场景 | [Go](./lc/lc_203.go) |
| [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/) | 迭代反转指针 | [Go](./lc/lc_206.go) |
| [707. 设计链表](https://leetcode.cn/problems/design-linked-list/) | 前哨节点维护长度 | [Go](./lc/lc_707.go) |

### 其他

| 题目 | 核心技巧 | 源码 |
| :--- | :--- | :--- |
| [15. 三数之和](https://leetcode.cn/problems/3sum/) | 排序 + 双指针去重 | [Go](./lc/lc_15.go) |
| [18. 四数之和](https://leetcode.cn/problems/4sum/) | 剪枝 + 双指针搜索 | [Go](./lc/lc_18.go) |
| [202. 快乐数](https://leetcode.cn/problems/happy-number/) | 哈希判环 / 快慢指针 | [Go](./lc/lc_202.go) |
| [242. 有效的字母异位词](https://leetcode.cn/problems/valid-anagram/) | 频次数组对比 | [Go](./lc/lc_242.go) |
| [349. 两个数组的交集](https://leetcode.cn/problems/intersection-of-two-arrays/) | 哈希集合去重 | [Go](./lc/lc_349.go) |

## 使用方式

1. 修改 `main.go` 中的调用，指向想要验证的题目函数。
2. 运行 `go run main.go` 快速查看输出或调试。
3. 需要批量验证时，可在本地初始化 `go mod` 后执行 `go test ./lc`。

欢迎继续补充题目或提交新的优化思路。