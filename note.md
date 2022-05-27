## 题目思路

### 2. 两数相加 ✅
模拟加法，进位操作  

### 4. 寻找两个正序数组的中位数 ✅
找第 k 小的数字，在两个数组中逐步找出第 k/2 小第数字，然后排除掉，再剩余数字里按照新的 k 继续找  

### 11. 盛最多水的容器 ✅
> 如果我们移动数字较大的那个指针，那么前者「两个指针指向的数字中较小值」不会增加，后者「指针之间的距离」会减小，那么这个乘积会减小。因此，我们移动数字较大的那个指针是不合理的。因此，我们移动 数字较小的那个指针。

### 56. 合并区间 ✅
将排序排序，检测是否有重叠，有就合并

### 10. 正则表达式匹配
TODO

### 55. 跳跃游戏 ✅
如果某一个作为 起跳点 的格子可以跳跃的距离是 3，那么表示后面 3 个格子都可以作为 起跳点
可以对每一个能作为 起跳点 的格子都尝试跳一次，把 能跳到最远的距离 不断更新
如果可以一直跳到最后，就成功了

链接：https://leetcode-cn.com/problems/jump-game/solution/55-by-ikaruga/

### 300. 最长递增子序列 ✅

### 23. 合并K个升序链表 ✅ 
使用堆排序

### 72. 编辑距离
TODO

### 124. 二叉树中的最大路径和
TODO

### 32. 最长有效括号
TODO

### 31. 下一个排列 ✅
将后面的「大数」与前面的「小数」交换  
在尽可能靠右的低位进行交换，需要从后向前查找  
将一个 尽可能小的「大数」 与前面的「小数」交换。比如 123465，下一个排列应该把 5 和 4 交换而不是把 6 和 4 交换  
将「大数」换到前面后，需要将「大数」后面的所有数重置为升序，升序排列就是最小的排列。  

### 543. 二叉树的直径 ✅
递归判断直径最长值，注意直径不一定通过根节点

### 128. 最长连续序列
先用哈希记录列表里的数字，随后对每个 nums[i]（假设为 x），逐步判断 x+1, x+2 是否存在于哈希中，并记录最大连续长度。
循环开始前，若发现 x-1 存在于哈希中，则跳过数字 x，因为在以 x-1 为起点的遍历中，已经记录过 x 了。

### 300. 最长递增子序列
dp[i] 的值代表 nums 以 nums[i]nums[i] 结尾的最长子序列长度。  
dp[i] = max(dp[i], dp[j] + 1) for j in [0, i)  

### 152. 乘积最大子数组
令imax为当前最大值，则当前最大值为 imax = max(imax * nums[i], nums[i])
由于存在负数，那么会导致最大的变最小的，最小的变最大的。因此还需要维护当前最小值imin，imin = min(imin * nums[i], nums[i])
当负数出现时则imax与imin进行交换再进行下一步计算

作者：guanpengchn
链接：https://leetcode-cn.com/problems/maximum-product-subarray/solution/hua-jie-suan-fa-152-cheng-ji-zui-da-zi-xu-lie-by-g/

### 560. 和为 K 的子数组
TODO

### 48. 旋转图像
先水平翻转，再对角线翻转
```go
package main
func rotate(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
```

### 169. 多数元素
从第一个数开始count=1，遇到相同的就加1，遇到不同的就减1，减到0就重新换个数开始计数，总能找到最多的那个

### 78. 子集
可以直接遍历数字，遇到一个数就把所有子集加上该数组成新的子集，遍历完毕即是所有子集。  
或者采用回溯法

### 394. 字符串解码
```
class Solution:
    def decodeString(self, s: str) -> str:
        stack, res, multi = [], "", 0
        for c in s:
            if c == '[':
                stack.append([multi, res])
                res, multi = "", 0
            elif c == ']':
                cur_multi, last_res = stack.pop()
                res = last_res + cur_multi * res
            elif '0' <= c <= '9':
                multi = multi * 10 + int(c)            
            else:
                res += c
        return res

作者：jyd
链接：https://leetcode-cn.com/problems/decode-string/solution/decode-string-fu-zhu-zhan-fa-di-gui-fa-by-jyd/
```
### 739. 每日温度
使用递减栈记录温度，如果下一个温度比当前栈顶的温度高，则弹出，并将这两个温度对应的天数相减，获得 answer[i]  

题解参考：https://leetcode-cn.com/problems/daily-temperatures/solution/leetcode-tu-jie-739mei-ri-wen-du-by-misterbooo/

### 198. 打家劫舍 ✅
动态规划 dp[i] = max(dp[i-2]+nums[i], dp[i-1])

### 64. 最小路径和 ✅
动态规划，使用二维数组存储到达每个格子的最小路径和即可

### 49. 字母异位词分组
使用 [26]int 统计各个字符串的字母频率，存到 map[[26]int]int 中

### 347. 前 K 个高频元素
哈希+堆

### 17. 电话号码的字母组合
回溯法

### 84. 柱状图中最大的矩形
TODO  
https://leetcode-cn.com/problems/largest-rectangle-in-histogram/solution/bao-li-jie-fa-zhan-by-liweiwei1419/

## 二分法的一些基本思路和特性  
使用闭区间的情况下，如果数组中查找不到目标元素，最终结束循环时，总是 right == left - 1