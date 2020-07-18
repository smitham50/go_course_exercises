package main

import "fmt"

func main() {
	returnee := returnFunc()

	str := returnee()

	fmt.Println(str)
}

func returnFunc() func() string {
	return func() string {
		return "Save me for later"
	}

}
