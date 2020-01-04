package main

import "fmt"

// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

func twoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for currentIdx, val := range nums {
		remainder := target - val

		if _, ok := seen[remainder]; ok {
			return []int{idx, currentIdx}
		}

		seen[val] = currentIdx
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}
