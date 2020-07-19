package main

import "fmt"

type person struct {
	name string
}

func (p *person) changeMe() {
	p.name = "My name is gone!"
}

func main() {
	aaron := person{
		name: "Aaron Smith",
	}

	aaron.changeMe()

	fmt.Println(aaron.name)
}
