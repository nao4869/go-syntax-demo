package main

import "fmt"

// type of 'deck'
type deck []string

// reciever
// d is instanc of deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
