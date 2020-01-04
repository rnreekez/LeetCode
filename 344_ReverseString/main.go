package main

import "fmt"

func reverseString(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}

func reverseStringV2(s []byte) {
	l, r := 0, len(s)-1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l, r = l+1, r-1
	}
}

func main() {
	/*
	   Example 1:

	   Input: ["h","e","l","l","o"]
	   Output: ["o","l","l","e","h"]
	   Example 2:

	   Input: ["H","a","n","n","a","h"]
	   Output: ["h","a","n","n","a","H"]
	*/

	// input := []byte{'h', 'e', 'l', 'l', 'o'}
	// input := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	input := []byte{'A', ' ', 'm', 'a', 'n', ',', ' ', 'a', ' ', 'p', 'l', 'a', 'n', ',', ' ', 'a', ' ', 'c', 'a', 'n', 'a', 'l', ':', ' ', 'P', 'a', 'n', 'a', 'm', 'a'}
	fmt.Printf("%+v\n", input)
	reverseString(input)
	fmt.Printf("%+v\n", input)
}
