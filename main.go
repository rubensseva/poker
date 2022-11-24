package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// TODO Simple validation (a hand of cards can't contain duplicates etc)

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/gen", func(w http.ResponseWriter, r *http.Request) {
		hand := drawFive()
		fmt.Fprintf(w, "hand: %v, analysis: %v", hand, analyze(hand))
	})

	// Expect requests in the form .../analyze-hand/3r-5r-6s-3r-tk
	http.HandleFunc("/analyze-hand/", func(w http.ResponseWriter, r *http.Request) {
		handRawStr := strings.TrimPrefix(r.URL.Path, "/analyze-hand/")
		if handRawStr == "" {
			fmt.Printf("no hand on analyze request")
			w.WriteHeader(500)
			return
		}
		handStrs := strings.Split(handRawStr, "-")
		if len(handStrs) != 5 {
			fmt.Printf("need exactly 5 cards, but got %v", handStrs)
			w.WriteHeader(500)
			return
		}

		hand := []card{}
		for _, c := range handStrs {
			hand = append(hand, card(c))
		}
		fmt.Fprintf(w, "hand: %v, analysis: %v", handStrs, analyze(hand))
	})

	fmt.Println("poker server starting...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Printf("http server: %v", err)
		os.Exit(1)
	}
	fmt.Println("bye!")
}
