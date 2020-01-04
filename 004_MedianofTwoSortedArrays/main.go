package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

}

func main() {
	/*
		nums1 = [1, 3]
		nums2 = [2]

		The median is 2.0
	*/
	n1 := []int{1, 3}
	n2 := []int{2}

	/*
	   nums1 = [1, 2]
	   nums2 = [3, 4]

	   The median is (2 + 3)/2 = 2.5
	*/
	// n1 := []int{1, 2}
	// n2 := []int{3, 4}

	fmt.Println(findMedianSortedArrays(n1, n2))
}
