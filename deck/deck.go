package deck

import (
	"fmt"
)

type Card struct {
	Shape  string
	Number string
}

func NewCard(shape string, number string) Card {
	return Card{
		Shape:  shape,
		Number: number,
	}
}

func Suit(shape string) []Card {
	var ret []Card
	a := []string{"A", "J", "K", "Q"}

	for i := 2; i <= 10; i++ {
		ret = append(ret, NewCard(shape, fmt.Sprintf("%v", i)))

	}
	for _, letter := range a {
		ret = append(ret, NewCard(shape, letter))
	}

	return ret
}
func NewDeck() [][]Card {
	var deck [][]Card
	var shapes = []string{"Diamond", "Heart", "Spade", "Flower"}
	counter := 0
	for _, shape := range shapes {
		if counter <=3 {
			deck = append(deck, Suit(shape))
			counter++
		}
	}
	return deck
}
