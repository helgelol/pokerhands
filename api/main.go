package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/helgelol/pokerhands/api/analyze"
	"github.com/helgelol/pokerhands/api/cards"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /new", newHand)
	router.HandleFunc("POST /check", checkHand)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Struct for Generated Hand
type GeneratedHandResponse struct {
	Hand       []string `json:"hand"`
	Evaluation string   `json:"evaluation"`
}

type EvaluatedHandResponse struct {
	Evaluation string `json:"evaluation"`
}

func newHand(w http.ResponseWriter, r *http.Request) {
	// Generate deck and deal cards
	hand := cards.Play()
	log.Printf("Hand Generated: %s", hand)

	// Decode generated hand
	var handStr []string
	err := json.Unmarshal([]byte(hand), &handStr)
	if err != nil {
		log.Fatalf("Error unmarshalling hand: %v", err)
	}
	evaluation := analyze.EvaluateHand(hand)
	response := GeneratedHandResponse{
		Hand:       handStr,
		Evaluation: evaluation,
	}

	result, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error marshalling response: %v", err)
	}
	log.Printf("Evaluation: %v", response.Evaluation)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func checkHand(w http.ResponseWriter, r *http.Request) {
	// Check Request Method
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode Body
	var handData GeneratedHandResponse
	err := json.NewDecoder(r.Body).Decode(&handData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Check if Hand is empty
	if len(handData.Hand) == 0 {
		http.Error(w, "Hand data cannot be empty", http.StatusBadRequest)
		return
	}

	// First data prep before analysis
	formattedAnalysis, err := json.Marshal(handData.Hand)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
	}

	// Second data prep and Evaluate received hand
	log.Printf("Hand Received for evaluation: %v", string(formattedAnalysis))
	response := analyze.EvaluateHand(string(formattedAnalysis))
	formatResponse := EvaluatedHandResponse{
		Evaluation: response,
	}
	log.Printf("Evaluation: %v", formatResponse)

	result, err := json.Marshal(formatResponse)
	if err != nil {
		log.Fatalf("Error marshalling response: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
