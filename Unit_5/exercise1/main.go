package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	favorite  string
}

func main() {
	p1 := person{
		firstName: "Aaron",
		lastName:  "Smith",
		favorite:  "Mint Chip",
	}

	p2 := person{
		firstName: "Cathi",
		lastName:  "Bartolomeo",
		favorite:  "Cookie Dough",
	}

	fmt.Println(p1.firstName)
	fmt.Println(p2)
}
