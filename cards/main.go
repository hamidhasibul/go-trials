package main

func main() {

	// var card string = "Ace of spades"

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
