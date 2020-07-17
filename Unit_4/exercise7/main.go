package main

import "fmt"

func main() {
	x := [][]string{}
	jb := []string{"James", "Bond", "Shaken, not stirred."}
	mm := []string{"Miss", "Moneypenny", "Hellooooo, James."}

	x = append(x, jb, mm)

	for _, row := range x {
		for _, data := range row {
			fmt.Println(data)
		}
	}
}
