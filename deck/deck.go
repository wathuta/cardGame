//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"sort"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)
const (
	maxRank = King
	minRank = Ace
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

//to do add options
//eg shuffling etc
//we use functional options
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		opt(cards)
	}

	return cards
}

//DefaultSort takes a slice of cards as an argument and returns a slice that is sorted
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

//Sort is a generic function that allows the user to sort the cards using a generic function
func Sort(less func(cards []Card) func(i, j int) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

//Less takes in slice of cards and returns a function that compares the cards at the indicies of its arguments.
func Less(cards []Card) func(i, j int) bool {

	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}
func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
