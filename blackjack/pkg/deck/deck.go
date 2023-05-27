//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8
type Rank uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

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
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(options ...func([]Card) []Card) []Card {
	var cards []Card
	for _, s := range suits {
		for rank := Ace; rank <= King; rank++ {
			cards = append(cards, Card{Rank: rank, Suit: s})
		}
	}

	for _, opt := range options {
		cards = opt(cards)
	}
	return cards
}

// Return a function to use as options in `New`
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	// Returns a sort function
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func Shuffle(cards []Card) []Card {
	shuffled := make([]Card, len(cards))

	// r := rand.New(rand.NewSource(99))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	prem := r.Perm(len(cards))

	for i, j := range prem {
		shuffled[i] = cards[j]
	}

	return shuffled
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

// Return a function to use as options in `New`
func AddJokers(n int) func([]Card) []Card {
	return func(c []Card) []Card {
		for i := 0; i < n; i++ {
			c = append(c, Card{Suit: Joker, Rank: Rank(i)})
		}
		return c
	}
}

// Return a function to use as options in `New`
func FilterCards(filterFunc func(card Card) bool) func([]Card) []Card {
	// Takes a function that return true if card is to keep
	return func(cards []Card) []Card {
		filtered := []Card{}
		for _, c := range cards {
			if filterFunc(c) {
				filtered = append(filtered, c)
			}
		}
		return filtered
	}
}

func MultiplyDecks(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		lenCards := len(cards)
		newDeck := make([]Card, lenCards*n)
		for i := 0; i < n; i++ {
			copy(newDeck[i*lenCards:(i+1)*lenCards], cards)
		}
		return newDeck
	}
}
