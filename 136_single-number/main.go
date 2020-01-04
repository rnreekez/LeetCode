package main

import "fmt"

func singleNumberV1(nums []int) int {
	seen := map[int]int{}

	for i := 0; i < len(nums); i++ {
		_, ok := seen[nums[i]]
		if ok {
			delete(seen, nums[i])
		} else {
			seen[nums[i]] = 1
		}
	}

	for k := range seen {
		return k
	}

	return -1
}

func singleNumber(nums []int) int {
	var res int
	for _, num := range nums {
		res ^= num
	}

	return res
}

func main() {
	/*
	   Input: [2,2,1]
	   Output: 1

	   Input: [4,1,2,1,2]
	   Output: 4
	*/

	// input := []int{2, 2, 1}
	input := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(input))
}
