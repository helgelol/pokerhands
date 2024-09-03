package analyze

import (
	"encoding/json"
	"log"
	"sort"
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
	// log.Printf("\nCard count: %v", cardCount)
	// log.Printf(string(len(card.Value)))
	// if len(cardCount) != 5 {
	// 	return "Invalid Hand"
	// }

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
	// Sort cardCount in increasing order
	var sortedCards []string
	for card := range cardCount {
		sortedCards = append(sortedCards, card)
	}
	sort.Strings(sortedCards)

	// Print sorted cardCount
	// log.Printf("\nCard count in increasing order: %v", sortedCards)
	// log.Printf("\nCard count after convert: %v", cardCount)
	// log.Printf("\n---\n")

	var result string
	for _, count := range cardCount {
		switch count {
		// case isRoyalFlush(cardCount):

		case 4:
			result = "Four of a Kind"
		case 3:
			result = "Three of a Kind"
		case 2:
			result = "Pair"
			// case 1:
			// 	result = "High Card"
			// default:
			// 	result = "High Card"
		}
	}

	// Check for other common types of poker hands
	// if len(cardCount) == 5 {
	// 	// Check for straight
	// 	if isStraight(cardCount) {
	// 		result = "Straight"
	// 	}
	// 	// Check for flush
	// 	if isFlush(suitCount) {
	// 		result = "Flush"
	// 	}
	// 	// Check for straight flush
	// 	if isStraight(cardCount) && isFlush(suitCount) {
	// 		result = "Straight Flush"
	// 	}
	// }

	// jsonResult, err := json.Marshal("hand: " + result)
	// if err != nil {
	// 	log.Fatalf("Error marshalling: %v", err)
	// }
	// return string(jsonResult)
	return result
}

// func isRoyalFlush(cardCount map[string]int)

// func isFullHouse(cardCount map[string]int) {
// 	if cardCount[0] == 3 && cardCount[1] == 2 {
// 		return true
// 	}
// }
