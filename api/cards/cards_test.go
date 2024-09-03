// FILEPATH: /Users/hfalch/kubes/pokerhands/api/cards/cards_test.go

package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	// Check if the deck has the correct number of cards
	expectedDeckSize := len(suit) * len(values)
	if len(deck) != expectedDeckSize {
		t.Errorf("Expected deck size of %d, but got %d", expectedDeckSize, len(deck))
	}

	// Check if each card in the deck has a valid suit and value
	for _, card := range deck {
		if !isValidSuit(card.Suit) {
			t.Errorf("Invalid suit: %s", card.Suit)
		}
		if !isValidValue(card.Value) {
			t.Errorf("Invalid value: %s", card.Value)
		}
	}
}

func TestShuffle(t *testing.T) {
	sampleDeck := newDeck()
	deck1 := newDeck()
	deck2 := shuffle(newDeck())

	// Check if the shuffled deck has the same number of cards as the original deck
	if len(deck1) != len(deck2) {
		t.Errorf("Expected shuffled deck size of %d, but got %d", len(deck1), len(deck2))
	}

	// Testing that all decks are created equal
	assert.Equal(t, sampleDeck, deck1)
	// And testing that shuffle works
	assert.NotEqual(t, deck1, deck2)
}

func TestDealHand(t *testing.T) {
	deck := newDeck()
	handSize := 5
	hand := dealHand(deck, handSize)

	// Check if the hand has the correct number of cards
	if len(hand) != handSize {
		t.Errorf("Expected hand size of %d, but got %d", handSize, len(hand))
	}

	// Check if the cards in the hand are from the original deck
	for _, card := range hand {
		if !isCardInDeck(card, deck) {
			t.Errorf("Card %s not found in the deck", cardToString(card))
		}
	}
}

func TestCardToString(t *testing.T) {
	card := Card{Suit: "s", Value: "a"}
	expectedString := "as"
	result := cardToString(card)

	if result != expectedString {
		t.Errorf("Expected card string '%s', but got '%s'", expectedString, result)
	}
}

func isValidSuit(suit string) bool {
	for _, s := range suit {
		if s != 's' && s != 'h' && s != 'r' && s != 'k' {
			return false
		}
	}
	return true
}

func isValidValue(value string) bool {
	for _, v := range value {
		if v != '2' && v != '3' && v != '4' && v != '5' && v != '6' && v != '7' && v != '8' && v != '9' && v != 't' && v != 'j' && v != 'q' && v != 'k' && v != 'a' {
			return false
		}
	}
	return true
}

func isCardInDeck(card Card, deck []Card) bool {
	for _, c := range deck {
		if c == card {
			return true
		}
	}
	return false
}
