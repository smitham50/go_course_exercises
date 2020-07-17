package main

import "fmt"

func main() {
	var x int
	var y string
	var z bool

	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("%v %v %v", x, y, z)

	fmt.Println(s)
}
