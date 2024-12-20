package main

import (
		"fmt"
		"math"
)

// TODO: implement the function computeGrade
func computeGrade(maxPoints float32, Points float32) float32 {
	var score float32
	score = (Points / maxPoints) * 5 + 1
	score = roundGrade(score)
	return score
}
func roundGrade(score float32) float32 {
	return float32(math.Round(float64(score)*100) / 100)
}

func main() {
	// TODO: call the function computeGrade
	var result float32 
	result = computeGrade(23.0, 17.5)

	fmt.Println("The computedGrade is:", result)
}
