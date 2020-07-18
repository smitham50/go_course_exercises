package main

import "fmt"

func main() {
	receiver(hello)
}

func receiver(f func() string) {
	str := f()

	fmt.Println(str)
}

func hello() string {
	return "Hi there!"
}
