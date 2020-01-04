package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	result := []string{}

	for i := 1; i <= n; i++ {
		newValue := ""

		if i%3 == 0 {
			newValue += "Fizz"
		}
		if i%5 == 0 {
			newValue += "Buzz"
		}

		if newValue == "" {
			newValue = strconv.Itoa(i)
		}

		result = append(result, newValue)
	}

	return result
}

func main() {
	for _, val := range fizzBuzz(15) {
		fmt.Println(val)
	}
}
