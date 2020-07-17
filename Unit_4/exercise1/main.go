package main

import "fmt"

func main() {
	arr := [5]int{}

	arr[0] = 2
	arr[1] = 5
	arr[2] = 8
	arr[3] = 6
	arr[4] = 11

	fmt.Printf("%T\n", arr)

	for _, v := range arr {
		fmt.Printf("the value is: %v of type %T\n", v, v)
	}
}
