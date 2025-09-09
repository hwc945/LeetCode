# 代码随想录

每日锻炼一些算法，现在主要来源于 LeetCode。

## 解题记录

### [704. 二分查找](https://leetcode.cn/problems/binary-search/)

- **完成日期**: 2025-09-04
- **要点**:
    - 注意「左闭右开」`[left, right)` 和「左闭右闭」`[left, right]` 区间的区别
    - `mid` 的取值方式 `left + (right - left) / 2` 可以有效防止整数溢出
    - 需要根据区间的定义，正确地更新 `left` 和 `right` 的位置
    - 对于已排序数组，二分查找已是理论最优解 O(log n)

### [27. 移除元素](https://leetcode.cn/problems/remove-element/)

- **完成日期**: 2025-09-05
- **要点**:
    - 注意算法思想和语言能力的区别。
    - 首位指针法，left代表下一个有效位置，right代表当前可替换的位置，直接覆盖而非交换，减少赋值次数，效率最佳
    - 快慢指针法，快指针遍历数组，慢指针记录新数组的位置，保持原数组顺序

### [977. 有序数组的平方](https://leetcode.cn/problems/squares-of-a-sorted-array/description/)

- **完成日期**: 2025-09-08
- **要点**:
    - 双指针，寻找平方最大值，记录在答案数组最后
  
### [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/description/)

- **完成日期**: 2025-09-09
- **要点**:
    - 滑动窗口，维护一个窗口的和，当和大于等于目标值时，尝试收缩左边界，更新最小长度
  
### [59. 螺旋矩阵 II](https://leetcode.cn/problems/spiral-matrix-ii/description/)

- **完成日期**: 2025-09-09
- **要点**:
    - 模拟螺旋过程：按照顺时针方向（右→下→左→上）填充矩阵。
    - 维护边界：使用 left, right, top, bottom四个变量控制填充范围。
    - 边界收缩：每次填充完一圈后，收缩边界以避免重复填充。
    - 终止条件：当填充的数字达到 n²时结束。