package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/helgelol/pokerhands/api/cards"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /new", newHand)
	router.HandleFunc("GET /check", checkHand)
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	cards.Play()
	// })

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func newHand(w http.ResponseWriter, r *http.Request) {
	hand := cards.Play()
	// check := analyze.EvaluateHand(hand)
	jsonResponse, err := json.Marshal(hand)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating deck: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func checkHand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Check hand")
}
