package analyze

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
)

// Should've reused existing Cards struct
type Card struct {
	Value string
	Suit  string
}

func parseCard(cardStr string) (Card, error) {

	value := string(cardStr[0])
	suit := string(cardStr[1])

	if len(cardStr) != 2 {
		return Card{}, fmt.Errorf("invalid card format")
	}

	validCards := []string{"2", "3", "4", "5", "6", "7", "8", "9", "t", "j", "q", "k", "a"}
	if !contains(validCards, value) {
		return Card{}, fmt.Errorf("invalid card value")
	}

	validSuits := []string{"k", "r", "h", "s"}
	if !contains(validSuits, suit) {
		return Card{}, fmt.Errorf("invalid suit")
	}

	return Card{
		Value: value,
		Suit:  suit,
	}, nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func EvaluateHand(hand string) string {
	var cardStrs []string
	err := json.Unmarshal([]byte(hand), &cardStrs)
	if err != nil {
		log.Fatalf("Error unmarshalling hand: %v", err)
	}

	if len(cardStrs) != 5 {
		return "Invalid Hand"
	}

	cardCount := make(map[string]int)
	suitCount := make(map[string]int)

	for _, cardStr := range cardStrs {
		card, err := parseCard(cardStr)
		if err != nil {
			return err.Error()
		}
		cardCount[card.Value]++
		suitCount[card.Suit]++
	}

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

	var sortedCards []string
	for card := range cardCount {
		sortedCards = append(sortedCards, card)
	}
	sort.Strings(sortedCards)

	// Check for straight
	if len(cardCount) == 5 {
		// Check if the difference between the maximum and minimum card values is 4
		maxCardValue := sortedCards[len(sortedCards)-1]
		minCardValue := sortedCards[0]
		maxValue, _ := strconv.Atoi(maxCardValue)
		minValue, _ := strconv.Atoi(minCardValue)
		if cardCount[maxCardValue] == 1 && cardCount[minCardValue] == 1 && (maxValue-minValue) == 4 {
			if len(suitCount) == 1 {
				if maxCardValue == "14" && minCardValue == "10" {
					return "Royal Straight Flush"
				}
				return "Straight Flush"
			}
			return "Straight"
		}
	}

	// Check for two pairs
	if len(cardCount) == 3 {
		pairCount := 0
		for _, count := range cardCount {
			if count == 2 {
				pairCount++
			}
		}
		if pairCount == 2 {
			return "Two Pairs"
		}
	}

	// Check for flush
	if len(suitCount) == 1 {
		return "Flush"
	}

	// Check for full house
	if len(cardCount) == 2 {
		for _, count := range cardCount {
			if count == 3 || count == 2 {
				return "Full House"
			}
		}
	}

	for _, count := range cardCount {
		switch count {
		case 5:
			return "Cheating hand"
		case 4:
			return "Four of a Kind"
		case 3:
			return "Three of a Kind"
		case 2:
			return "Pair"
		}
	}

	return "High Card"
}
