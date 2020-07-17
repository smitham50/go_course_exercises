package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(sum(nums...))

	fmt.Println(sum2(nums))
}

func sum(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}

func sum2(nums []int) int {
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}
