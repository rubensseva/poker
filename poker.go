package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type card string
type analysis struct {
	goal string
	cards []card
}

var (
	colors = []byte{'r', 's', 'h', 'k'}
	values = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 't', 'k', 'd', 'k', 'a'}
)

func newStack() []card {
	cards := make([]card, 0, 52)

	for i := 0; i < 4; i++ {
		color := colors[i]
		for j := 0; j < 12; j++ {
			value := values[j]
			cards = append(cards, card([]byte{color, value}))
		}
	}

	return cards
}

func shuffle(foo []card) {
	for i := range foo {
		j := rand.Intn(len(foo))
		tmp := foo[i]
		foo[i] = foo[j]
		foo[j] = tmp
	}
}

func drawFive() []card {
	cards := newStack()
	shuffle(cards)
	return cards[:5]
}

func equalCards(val byte, hand []card) []card {
	equal := []card{}
	for _, c := range hand {
		if c[0] == val {
			equal = append(equal, c)
		}
	}
	return equal
}



// TODO: Fix var names...
func containsGoal(g string, a []analysis) bool {
	for _, aa := range a {
		if g == aa.goal {
			return true
		}
	}
	return false
}

func analyze(hand []card) []analysis {
	analysises := []analysis{}
	// Pairs
	for _, c := range hand {
		eqCards := equalCards(c[0], hand)
		if len(eqCards) > 1 {
			new := analysis{
				goal: fmt.Sprintf("%d-PAIR", len(eqCards)),
				cards: eqCards,
			}
			analysises = append(analysises, new)
		}
	}

	// House
	for _, c := range hand {
		eqCards := equalCards(c[1], hand)
		if len(eqCards) == 2 {
			for _, cc := range hand {
				eqqCards := equalCards(cc[1], hand)
				if len(eqqCards) == 3 {
					new := analysis{
						goal: "FULL HOUSE",
						cards: append(eqCards, eqqCards...),
					}
					analysises = append(analysises, new)
				}
			}
		}
	}

	// Flush
	col := hand[0][1]
	isFlush := true
	for i := 1; i < len(hand); i++ {
		if hand[i][1] != col {
			isFlush = false
		}
	}
	if isFlush {
		analysises = append(
			analysises,
			analysis{
				goal:  "FLUSH",
				cards: hand,
			},
		)
	}

	// Small Straight
	// Sort hand by values
	sort.Slice(hand, func(i, j int) bool {
		return hand[i][1] < hand[j][1]
	})

	isStraight := true
	for i := 1; i < len(hand); i++ {
		if hand[i][0] != hand[i - 1][0] + 1 {
			isStraight = false
		}
	}

	if isStraight {
		new := analysis{
			cards: hand,
		}
		if hand[len(hand) - 1][1] == 'a' {
			new.goal = "HIGH STRAIGHT"
		} else {
			new.goal = "STRAIGHT"
		}
		analysises = append(analysises, new)
	}

	if containsGoal("HIGH STRAIGHT", analysises) && containsGoal("FLUSH", analysises) {
		analysises = []analysis{{
			goal:  "ROYAL FLUSH",
			cards: hand,
		}}
	}

	return analysises
}
