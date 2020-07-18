package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v, and I'm %v\n", p.first, p.last, p.age)
}

func main() {
	aaron := person{
		first: "Aaron",
		last:  "Smith",
		age:   36,
	}

	aaron.speak()
}
