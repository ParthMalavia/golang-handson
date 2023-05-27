package main

import (
	"fmt"
	"strings"

	"example/blackjack/pkg/deck"
)

const NO_DECK = 1

type Hand []deck.Card

func (h Hand) String() string {
	s := []string{}
	for _, h := range h {
		s = append(s, h.String())
	}
	return strings.Join(s, ", ")
}

// function to count score
func (h Hand) Score() int {
	score := 0
	aces := false

	// Count score with taking 1 point for Ace
	for _, c := range h {
		if c.Rank == deck.Ace {
			aces = true
		}
		score += min(int(c.Rank), 10)
	}

	// use 11 points for ace if score is < 11
	if aces && score < 11 {
		score += 10
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Function for dealer's hand

func main() {
	// Create new deck
	cards := deck.New(deck.MultiplyDecks(NO_DECK), deck.Shuffle)

	play := "Y"
	for play == "y" || play == "Y" {

		// Shuffle deck
		if len(cards) < 20 {
			fmt.Println("********** SHUFFLING **********")
			cards = deck.New(deck.MultiplyDecks(NO_DECK), deck.Shuffle)
		}

		var card deck.Card
		var Player, Dealer Hand

		// Dead each player 2 cards
		for i := 0; i < 2; i++ {
			for _, hand := range []*Hand{&Player, &Dealer} {
				card = draw(&cards)
				*hand = append(*hand, card)
			}
		}
		// Dealer's visible cards score
		pScore, dScore := Player.Score(), min(int(Dealer[0].Rank), 10)
		if Dealer[0].Rank == deck.Ace {
			dScore = 11
		}

		// Player's Move
		var input string
		for input != "s" && pScore < 21 {
			fmt.Println("Player:", pScore, "\nCards:", Player)
			fmt.Println("\nDealer:", dScore, "\nCards:", Dealer[0], "****HIDDEN CARD****")
			fmt.Println("\nPlayer's Move: hit(h), stand(s)")
			fmt.Scanf("%s\n", &input)

			// Add moves: split and double
			switch input {
			case "h":
				card = draw(&cards)
				Player = append(Player, card)
			case "s":
				fmt.Println("Player stands, Dealer's move")
			default:
				fmt.Println("Enter valid input")
			}

			pScore = Player.Score()
		}

		// Dealer's Move
		// Hit till score > player-score
		// NOTE: No need to Hit if Player is busted
		dScore = Dealer.Score()
		if pScore <= 21 {
			for dScore < pScore {
				card = draw(&cards)
				Dealer = append(Dealer, card)
				dScore = Dealer.Score()
			}
		}

		// Print final hands
		showFinalResult(Player, Dealer)

		fmt.Println("\nWould you like to continue?(Y/N)")
		fmt.Scanln(&play)
	}

}

func draw(deck *[]deck.Card) deck.Card {
	card := (*deck)[0]
	*deck = (*deck)[1:]
	return card
}

func showFinalResult(Player, Dealer Hand) {
	pScore, dScore := Player.Score(), Dealer.Score()

	fmt.Println("######## FINAL Hands ##########")
	fmt.Println("Player:", pScore)
	fmt.Println("Cards:", Player)
	fmt.Println("\nDealer:", dScore)
	fmt.Println("Cards:", Dealer)
	fmt.Println()

	switch {
	case pScore > 21:
		fmt.Println("YOU BUSTED.")
	case dScore > 21:
		fmt.Println("DEALER BUSTED.")
	case pScore > dScore:
		fmt.Println("YOU WIN!")
	case dScore > pScore:
		fmt.Println("YOU LOSE.")
	case dScore == pScore:
		fmt.Println("DRAW!!!")
	}
}
