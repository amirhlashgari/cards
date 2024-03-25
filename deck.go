package main

import "fmt"

type deck []string

// below is a "Receiver Function"
// you can think "d" as this, self in other languages
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
