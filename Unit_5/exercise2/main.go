package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	favorite  []string
}

func main() {

	p1 := person{
		firstName: "Aaron",
		lastName:  "Smith",
		favorite: []string{
			"Mint Chip",
			"Chocolate",
			"Raspberry",
		},
	}

	p2 := person{
		firstName: "Cathi",
		lastName:  "Bartolomeo",
		favorite: []string{
			"Cookie Dough",
			"Peanut Butter",
			"Vanilla Bean",
		},
	}

	people := map[string]person{
		"smith_aaron":      p1,
		"bartolomeo_cathi": p2,
	}

	fmt.Println(people)

	for k, v := range people {
		fmt.Println(k, v)
	}
}
