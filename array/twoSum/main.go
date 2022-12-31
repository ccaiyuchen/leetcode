package main

import (
	"fmt"
)

// twoSum2 使用哈希表
func twoSum2(nums []int, target int) []int {
	var best []int
	record := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := target - nums[i]
		if _, ok := record[num]; ok {
			return []int{record[num], i}
		}
		record[nums[i]] = i
	}
	return best
}

// twoSum1 使用两重循环
func twoSum1(nums []int, target int) []int {
	var best []int
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				best = append(best, i, j)
			}
		}
	}
	return best
}

func main() {
	nums := []int{7, 11, 15, 2}
	target := 9
	fmt.Println(twoSum2(nums, target))
}
