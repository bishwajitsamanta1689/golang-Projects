package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck{
	cards:= deck{}
	/*
		Adding 52 cards in the slice is the goal, so instead of writing out manually we need to take an approach different
		1. A variable which contains the Suit Name for the Cards Slice.
		2. Another variable which contains the Number for those Cards
		3. We need to make nested loops where first loop is about containing the Suit names of the cards
	       and second loop is going to iterate over the card Values and finally create a new card they will append to the main
	       list of cards which is currently empty now.
		4. Finally return the card Slice
	 */
	cardSuites:= []string{"Spades","Diamonds","Hearts","clubs"}
	cardValues:=[]string{"Ace","One","Two","Three","Four","Five","Six","Seven","Eight","Nine"}
	for _, suit:= range cardSuites {
		for _, values:= range cardValues {
			cards = append(cards, values + "Cards of " + suit)
		}
	}
	return cards
}

func (d deck) print(){
	for _, card:= range d {
		fmt.Println(card)
	}
}