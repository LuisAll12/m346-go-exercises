package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a float64, b float64) float64 {
	var c float64
	c = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return c
}
func main() {
	// TODO: call the function computeHypotenuse
	var result float64
	result = computeHypotenuse(3.0, 2.0)
	fmt.Println("The computed Hypotenuse is:", result)
}
