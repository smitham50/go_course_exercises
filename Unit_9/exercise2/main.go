package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hi, I'm", p.name)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	aaron := person{
		name: "Aaron",
	}

	saySomething(&aaron)
}
