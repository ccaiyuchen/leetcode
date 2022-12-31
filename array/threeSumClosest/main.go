package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// 1.给一个默认值
	best := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			}
			if abs(sum-target) < abs(best-target) {
				best = sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	// 输出2
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}
