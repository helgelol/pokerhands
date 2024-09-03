package analyze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCard(t *testing.T) {
	card, _ := parseCard("2h")
	assert.Equal(t, "2", card.Value)
	assert.Equal(t, "h", card.Suit)
}

func TestFourOfAKind(t *testing.T) {
	hand := "[\"2k\",\"2s\",\"2h\",\"2r\",\"6h\"]"
	evaluation := EvaluateHand(hand)
	assert.Equal(t, "Four of a Kind", evaluation)
}

func TestPair(t *testing.T) {
	hand := "[\"2k\",\"2s\",\"4h\",\"8r\",\"6h\"]"
	evaluation := EvaluateHand(hand)
	assert.Equal(t, "Pair", evaluation)
}

func TestThreeOfAKind(t *testing.T) {
	hand := "[\"2k\",\"2s\",\"2h\",\"5r\",\"6h\"]"
	evaluation := EvaluateHand(hand)
	assert.Equal(t, "Three of a Kind", evaluation)
}

func TestStraight(t *testing.T) {
	hand := "[\"2k\",\"3s\",\"4h\",\"5r\",\"6h\"]"
	evaluation := EvaluateHand(hand)
	assert.Equal(t, "Straight", evaluation)
}

func TestStraightFlush(t *testing.T) {
	hand := "[\"2s\",\"3s\",\"4s\",\"5s\",\"6s\"]"
	evaluation := EvaluateHand(hand)
	assert.Equal(t, "Straight Flush", evaluation)
}

func TestFlush(t *testing.T) {
	hand := "[\"2r\",\"4r\",\"6r\",\"8r\",\"tr\"]"
	evaluation := EvaluateHand(hand)
	assert.Equal(t, "Flush", evaluation)
}

func TestRoyalStraightFlush(t *testing.T) {
	hand := "[\"ah\",\"kh\",\"qh\",\"jh\",\"th\"]"
	evaluation := EvaluateHand(hand)
	assert.Equal(t, "Royal Straight Flush", evaluation)
}

func TestTwoPairs(t *testing.T) {
	hand := "[\"4h\",\"4r\",\"th\",\"8s\",\"8k\"]"
	evaluation := EvaluateHand(hand)
	assert.Equal(t, "Two Pairs", evaluation)
}

func TestCheatingHand(t *testing.T) {
	hand := "[\"4h\",\"4r\",\"4h\",\"4s\",\"4k\"]"
	evaluation := EvaluateHand(hand)
	assert.Equal(t, "Cheating hand", evaluation)
}

func TestInvalidSuit(t *testing.T) {
	hand := "[\"4k\",\"4s\",\"4r\",\"4k\",\"5f\"]"
	evaluation := EvaluateHand(hand)
	assert.Equal(t, "invalid suit", evaluation)
}

func TestInvalidValue(t *testing.T) {
	hand := "[\"4h\",\"4r\",\"th\",\"8s\",\"fk\"]"
	evaluation := EvaluateHand(hand)
	assert.Equal(t, "invalid card value", evaluation)
}

func TestInvalidLength(t *testing.T) {
	hand := "[\"4h\",\"4r\",\"th\",\"8ks\",\"fk\"]"
	evaluation := EvaluateHand(hand)
	assert.Equal(t, "invalid card format", evaluation)
}
