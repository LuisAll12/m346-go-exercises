package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	var firstName string = "Luis"
	var lastName string = "Allamand"
	var dayOfBirth int = 19
	var monthOfBirth int = 6
	var yearOfBirth int = 2008
	var numberOfSiblings int = 2
	var heightInMeters float32 = 1.68
	var zodiacSign string = "twin"
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %s\n", zodiacSign)
}
