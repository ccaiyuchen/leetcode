package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var best [][]int
	// 枚举第3个数
	for i := 0; i < n; i++ {
		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		target := 0 - nums[i]
		for l < r {
			if nums[l]+nums[r] == target {
				best = append(best, []int{nums[i], nums[l], nums[r]})
				// 移动的过程中注意去重
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if nums[l]+nums[r] < target {
				l++
			} else {
				r--
			}
		}
	}
	return best
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
