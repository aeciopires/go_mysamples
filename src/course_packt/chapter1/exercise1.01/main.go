package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// Select randon number between 0 and 5 and add 1
	randon := rand.Intn(5) + 1
	// Repeat start according randon number
	stars := strings.Repeat("*", randon)
	// Print stars
	fmt.Println(stars)
}
