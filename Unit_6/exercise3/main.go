package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Laterrr...")
	}()

	fmt.Println("Me first")
}
