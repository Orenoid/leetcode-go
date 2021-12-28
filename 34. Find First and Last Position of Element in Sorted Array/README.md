[34. Find First and Last Position of Element in Sorted Array]()

`二分法` `数组`

思路：  
分别查找左右边界，二分法查找过程中，若 mid 等于 target，判断其左/右侧的数字是否还是 target，若不是，说明已经找到边界了。

犯错点：
1. 找到边界后，忘记 break