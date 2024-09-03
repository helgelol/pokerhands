package cards

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Value string
}

var suit = [...]string{"s", "h", "r", "k"}
var values = [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "t", "j", "q", "k", "a"}

func newDeck() []Card {
	deck := make([]Card, len(suit)*len(values))
	for i := range deck {
		deck[i] = Card{suit[i%len(suit)], values[i/len(suit)]}
	}
	return deck
}

func shuffle(deck []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(deck) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

func dealHand(deck []Card, handSize int) []Card {
	if handSize > len(deck) {
		log.Fatal("Request exceeds deck size.")
		return nil
	}
	return deck[:handSize]
}

// Convert cards to specified string format
func cardToString(card Card) string {
	return card.Value + card.Suit
}

func Play() string {
	deck := newDeck()
	shuffledDeck := shuffle(deck)
	hand := dealHand(shuffledDeck, 5)

	var handStr []string
	for _, card := range hand {
		handStr = append(handStr, cardToString(card))
	}

	handJson, err := json.Marshal(handStr)
	if err != nil {
		log.Fatal("Error marshalling hand to JSON:", err)
	}
	return string(handJson)
}
