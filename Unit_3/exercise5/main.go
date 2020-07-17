package main

import "fmt"

func main() {

	x := []int{1, 4, 8, 12, 45, 64, 33, 21}
	y := []int{5, 5, 5, 6}

	x = append(x, y...)

	fmt.Println(x)

}
