package main

import (
	"math/rand"
	"sort"
)

func shuffle[T any](s []T) {
	for i := range s {
		j := rand.Intn(len(s))
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
	}
}

func equalHands(_h1 []card, _h2 []card) bool {
	// Need to convert to []string type in order to use sort.Strings.
	// Additionally, it is polite to not modify the arguments here.
	h1 := []string{}
	h2 := []string{}

	for _, el := range _h1 {
		h1 = append(h1, string(el))
	}
	for _, el := range _h2 {
		h2 = append(h2, string(el))
	}

	sort.Strings(h1)
	sort.Strings(h2)

	if len(h1) != len(h2) {
		return false
	}

	for i := range h1 {
		if h1[i] != h2[i] {
			return false
		}
	}
	return true
}
