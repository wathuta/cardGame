package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})

	// Output:
	//Ace of Hearts

}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck")
	}
	for i := 0; i < len(cards); i++ {
		if i < 13 {
			if cards[i].Suit != Spade {
				t.Error("wrong order of suits")
			}
		} else if i >= 13 && i < 26 {
			if cards[i].Suit != Diamond {
				t.Error("wrong order of suits")
			}
		} else if i >= 26 && i < 39 {
			if cards[i].Suit != Club {
				t.Error("wrong order of suits")
			}
		} else if i >= 39 && i < 52 {
			if cards[i].Suit != Heart {
				t.Error("wrong order of suits")
			}
		} else if i == 52 {
			if cards[i].Suit != Joker {
				t.Error("wrong possition of Joker")
			}
		}
	}
}
func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	if (cards[0] != Card{Rank: Ace, Suit: Spade}) {
		t.Error("Expected Ace of Spades as first card, Received", cards[0])
	}
}
