package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// type of 'deck'
type deck []string

// return value of type deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clover"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// reciever
// d is instanc of deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//["a", "b", "c"] -> "a", "b", "c"
	return strings.Join(
		[]string(d), 
		", ",
	) 
}

// function to save deck as file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(
		filename, 
		[]byte(d.toString()), 
		0666,
	)
}
