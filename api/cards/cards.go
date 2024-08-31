package cards

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Value string
}

// Basert på Case, ser det ut kort identifiseres på første bokstav i sorten
var suit = [...]string{"s", "h", "r", "k"}
var values = [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "t", "j", "q", "k", "a"}

func newDeck() []Card {
	deck := make([]Card, len(suit)*len(values))
	for i := range deck {
		deck[i] = Card{suit[i%len(suit)], values[i/len(suit)]}
	}
	return deck
}

func shuffle(deck []Card) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range deck {
		j := r.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

func dealHand(deck []Card, handSize int) []Card {
	if handSize > len(deck) {
		fmt.Println("Request exceeds deck size.")
		return nil
	}
	return deck[:handSize]
}

func Play() []Card {
	deck := newDeck()
	shuffle(deck)
	hand := dealHand(deck, 5)
	return hand
}
