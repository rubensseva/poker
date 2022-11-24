package main

import (
	"sort"
)

type combination struct {
	name  string
	cards []card
}

func containsName(n string, combs []combination) bool {
	for _, comb := range combs {
		if n == comb.name {
			return true
		}
	}
	return false
}

func analyze(hand []card) []combination {
	combs := []combination{}
	// X of a kind, and single pairs
	for _, c := range hand {
		eqCards := filterCardsByValue(c[0], hand)
		if len(eqCards) > 1 {
			new := combination{
				name:  numEqualToString[len(eqCards)],
				cards: eqCards,
			}
			combs = append(combs, new)
		}
	}

	// House
	foundHouse := false
	for _, c := range hand {
		if foundHouse {
			break
		}
		eqCards := filterCardsByValue(c[0], hand)
		if len(eqCards) == 2 {
			for _, cc := range hand {
				eqqCards := filterCardsByValue(cc[0], hand)
				if len(eqqCards) == 3 {
					new := combination{
						name:  "full house",
						cards: append(eqCards, eqqCards...),
					}
					combs = append(combs, new)
				}
				foundHouse = true
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
		combs = append(
			combs,
			combination{
				name:  "flush",
				cards: hand,
			},
		)
	}

	// Straight
	// Sort hand by values
	sort.Slice(hand, func(i, j int) bool {
		return valueOfValue[hand[i][0]] < valueOfValue[hand[j][0]]
	})

	isStraight := true
	for i := 1; i < len(hand); i++ {
		if valueOfValue[hand[i][0]] != valueOfValue[hand[i-1][0]]+1 {
			isStraight = false
		}
	}
	if isStraight {
		new := combination{
			cards: hand,
		}
		if hand[len(hand)-1][0] == 'a' {
			new.name = "high straight"
		} else {
			new.name = "straight"
		}
		combs = append(combs, new)
	}

	if containsName("high straight", combs) && containsName("flush", combs) {
		combs = append(combs, combination{
			name:  "royal flush",
			cards: hand,
		})
	}

	// Remove duplicates
	for i := 0; i < len(combs); i++ {
		for j := 0; j < len(combs); j++ {
			if i == j {
				continue
			}
			if combs[i].name == combs[j].name &&
				equalHands(
					combs[i].cards,
					combs[j].cards) {
				new := combs[:j]
				new = append(new, combs[j+1:]...)
				combs = new
			}
		}
	}

	// Two pairs
	foundTwoPairs := false
	for i := 0; i < len(combs); i++ {
		if foundTwoPairs {
			break
		}
		if combs[i].name == "pair" {
			for j := 0; j < len(combs); j++ {
				if i == j {
					continue
				}
				if combs[j].name == "pair" {
					combs = append(
						combs,
						combination{
							name: "two pairs",
							cards: append(
								combs[i].cards,
								combs[j].cards...,
							),
						},
					)
					foundTwoPairs = true
					break
				}
			}
		}
	}

	return combs
}
