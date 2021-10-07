package deck

import (
	"fmt"
	"testing"
)

func TestSuit(t *testing.T) {
	cards := NewDeck()
	fmt.Println(cards)
}
