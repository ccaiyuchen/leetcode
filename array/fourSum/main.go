package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var best [][]int
	// 枚举第 1 个数
	for i := 0; i < n; i++ {
		// 与上一个数去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, n-1
			sum := target - nums[i] - nums[j]
			for l < r {
				if nums[l]+nums[r] == sum {
					best = append(best, []int{nums[i], nums[j], nums[l], nums[r]})
					// 去重
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if nums[l]+nums[r] < sum {
					l++
				} else {
					r--
				}
			}
		}
	}
	return best
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
}
