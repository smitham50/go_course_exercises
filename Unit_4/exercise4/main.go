package main

import "fmt"

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	s = append(s, 52)

	fmt.Println(s)

	s = append(s, 53, 54, 55)

	fmt.Println(s)

	y := []int{56, 57, 58, 59, 60}

	s = append(s, y...)

	fmt.Println(s)
}
