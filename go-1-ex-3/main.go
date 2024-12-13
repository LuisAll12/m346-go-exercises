package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// Writing to stdout using fmt.Fprintln
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")
	fmt.Fprintln(os.Stdout, "the dice was rolled at", when)

	// TODO: write the output into eyes.txt and dice.log
	// Open or create eyes.txt
	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating eyes.txt:", err)
		return
	}
	defer eyesFile.Close()

	fmt.Fprintln(eyesFile, "the dice shows", eyes, "eyes")
	fmt.Fprintln(eyesFile, "the dice was rolled at", when)

	diceLogFile, err := os.Create("dice.log")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating dice.log:", err)
		return
	}
	defer diceLogFile.Close()

	fmt.Fprintln(diceLogFile, "the dice shows", eyes, "eyes")
	fmt.Fprintln(diceLogFile, "the dice was rolled at", when)

	fmt.Fprintln(os.Stdout, "Output written to eyes.txt and dice.log")
}

