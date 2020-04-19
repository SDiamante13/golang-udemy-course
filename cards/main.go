package main

import "fmt"

func main() {

	cards := newDeckFromFile("deck.txt")
	cards.shuffle()
	cards.print()
	fmt.Println(len(cards))




}
