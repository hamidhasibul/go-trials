package main

func main() {

	// var card string = "Ace of spades"

	cards := newDeckFromFile("my-cards.txt")

	cards.print()

}
