package main

import "fmt"


func main() {
	// init variable and assign value to it
	cards := deck{"Test string", newCard()}
	cards = append(cards, "Test string 3")  // add item in SLICE with append
	cards.print()
	fmt.Println("----------------------")
	newCards := newDeck()
	//newCards.print()
	hand, remainingCards := deal(newCards, 3)
	hand.print()  // hand & remainingCards both are type deck so can use print()
	remainingCards.print()

	fmt.Println("----------------------")

	fmt.Println(newCards.toString())

	fmt.Println("----------------------")

	err := newCards.saveToFile("myCards")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	newCardsFromFile :=newDeckFromFile("myCards")
	newCardsFromFile.print()

	fmt.Println("----------------------")

	newCardsFromFile.shuffle()
	newCardsFromFile.print()

}

func newCard() string {
	return "test string 2"
}
