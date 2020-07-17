package main

import "fmt"

func main() {
	x := make([]int, 50, 50)

	for i := 0; i < len(x); i++ {
		x[i] = i + 10
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
