package analyze

import (
	"encoding/json"
	"fmt"
	"log"
)

// Should've reused existing Cards struct
type Card struct {
	Value string
	Suit  string
}

func parseCard(cardStr string) Card {
	return Card{
		Value: string(cardStr[0]),
		Suit:  string(cardStr[1]),
	}
}

func EvaluateHand(hand string) string {
	fmt.Printf("\nHand in: %s", hand)
	var cardStrs []string
	err := json.Unmarshal([]byte(hand), &cardStrs)
	if err != nil {
		log.Fatalf("Error unmarshalling hand: %v", err)
	}

	cardCount := make(map[string]int)
	suitCount := make(map[string]int)

	for _, cardStr := range cardStrs {
		card := parseCard(cardStr)
		cardCount[card.Value]++
		suitCount[card.Suit]++
	}

	// fmt.Printf("\nCard counts: %v", cardCount)
	// fmt.Printf("\nSuit counts: %v", suitCount)

	// Convert card values to numbers
	for card, count := range cardCount {
		switch card {
		case "t":
			cardCount["10"] = count
			delete(cardCount, card)
		case "j":
			cardCount["11"] = count
			delete(cardCount, card)
		case "q":
			cardCount["12"] = count
			delete(cardCount, card)
		case "k":
			cardCount["13"] = count
			delete(cardCount, card)
		case "a":
			cardCount["14"] = count
			delete(cardCount, card)
		}
	}
	fmt.Printf("\nCard count after convert: %v", cardCount)
	fmt.Printf("\n---\n")

	var result string
	for _, count := range cardCount {
		switch count {
		case 4:
			result = "Four of a Kind"
		case 3:
			result = "Three of a Kind"
		case 2:
			result = "Pair"
		}
	}
	return hand + result

	// return hand
}
