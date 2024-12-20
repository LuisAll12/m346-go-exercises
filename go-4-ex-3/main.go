package main

import (
	"fmt"
	"math"
)
// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a , b , c float64) (float64, float64) {
	var result1, result2 float64

	result1 = ((b*-1)+math.Sqrt(math.Pow(b, 2)-4*(a*c))) / (2 *a)
	result2 = ((b*-1)-math.Sqrt(math.Pow(b, 2)-4*(a*c))) / (2 *a)
	return result1, result2
}

func main() {
	// TODO: call the function computeQuadraticFormula
	var result1, result2 float64
	result1, result2 = computeQuadraticFormula(1.0, 2.0, -3.0)
	fmt.Println("The results are:", result1, result2)
}
