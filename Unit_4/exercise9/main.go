package main

import "fmt"

func main() {
	x := map[string][]string{}

	x["Bond"] = []string{"bond_james", "shaken, not stirred", "martinis", "women"}
	x["Moneypenny"] = []string{"moneypenny_miss", "james bond", "literature", "computer science"}
	x["No"] = []string{"no_dr", "being evil", "ice cream", "sunsets"}

	delete(x, "No")

	fmt.Println(x)
}
