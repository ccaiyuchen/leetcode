package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) (ans int) {
	// 定义一个数组存储每个字符上一次出现的位置
	pos := make(map[byte]int)
	//滑动窗口、双指针
	// i 是右边界。j 是左边界
	for i, j := 0, 0; i < len(s); i++ {
		if p, ok := pos[s[i]]; ok {
			j = max(j, p+1)
		}
		// 当前窗口的长度
		ans = max(ans, i-j+1)
		pos[s[i]] = i
	}
	fmt.Println()
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func main() {
	s := " "
	fmt.Println(lengthOfLongestSubstring(s))
	//fmt.Println(lengthOfLongestSubstring2(s))
}
