package main

import (
	"math/rand"
)

func shuffle[T any](s []T) {
	for i := range s {
		j := rand.Intn(len(s))
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
	}
}
