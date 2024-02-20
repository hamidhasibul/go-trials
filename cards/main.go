package main

func main() {

	// var card string = "Ace of spades"

	cards := newDeck()

	cards.saveToFile("my-cards.txt")

}
