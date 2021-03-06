package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			break
		}
		digits[i] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
		}
	}

	return digits
}

func main() {
	var d = []int{7, 8, 9, 8}
	fmt.Println(plusOne(d))
}
