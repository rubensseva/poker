package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Hello world!")

	fmt.Println(newStack())
	for i := 0; i < 5; i++ {
		fmt.Println(drawFive())
	}
}
