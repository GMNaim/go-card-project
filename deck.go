package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string // SLICE : array with unlimited range

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "two", "Three", "Four"}

	for _, card := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+card)
		}
	}
	return cards
}

// RECEIVER: like class object, d is an object of class deck
// every instance/variable of type deck can access the method print()
// run the loop upto range of d
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // will return 2 value of type deck
	// divide the decks based on hand size/person and keep rest of the deck separate

	// SLICE can be sliced like python list,
	// like, a = [1, 2, 3, 4]; a[0:2] == 1, 2
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// make the deck to string

	return strings.Join(d, ",") // Join method of package strings
}

func (d deck) saveToFile(fileName string) error {
	// save the decks to a file

	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	// Get the decks from file

	byteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		// if there is error...
		fmt.Println("Error", err)
		os.Exit(1) // if value of code parameter of Exit() is other than 0
		// it thinks there is something went wrong and exit the program.
	}
	stringSlice := strings.Split(string(byteSlice), ",")

	return deck(stringSlice)
}

func (d deck) shuffle() {
	// Shuffle the decks

	// Creating random seed value (using time package) for package rand for getting random number
	customSource := rand.NewSource(time.Now().UnixNano())
	customRand := rand.New(customSource)

	for i := range d {
		newPosition := customRand.Intn(len(d) - 1)
		// shuffling the positions
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
