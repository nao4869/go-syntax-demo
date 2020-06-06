package main

//import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	
	//fmt.Println(cards.toString())

	//cards.print()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// print byte slice, type conversion
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	
}
