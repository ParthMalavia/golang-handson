package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Spade}.String())
	fmt.Println(Card{Rank: Two, Suit: Heart}.String())
	fmt.Println(Card{Rank: Jack, Suit: Diamond}.String())
	fmt.Println(Card{Suit: Joker}.String())

	// Output:
	// Ace of Spades
	// Two of Hearts
	// Jack of Diamonds
	// Joker

}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in the new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("First expected card is Ace of spades and got,", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(AddJokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}

	if count != 3 {
		t.Error("Expected jokers are 3, Got", count)
	}
}

func TestFilterFunction(t *testing.T) {
	spadeFilter := func(c Card) bool {
		return c.Suit == Spade
	}

	cards := New(FilterCards(spadeFilter))
	for _, c := range cards {
		if c.Suit != Spade {
			t.Error("After applying spade filter we have card with suit", c.Suit.String())
		}
	}
}

func TestMultipleDecks(t *testing.T) {
	cards := New(MultiplyDecks(3))

	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards got %d cards", 13*4*3, len(cards))
	}
}
