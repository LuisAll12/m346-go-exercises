package main

import "fmt"

// FullName enthält den Vor- und Nachnamen einer Person
type FullName struct {
	FirstName string
	LastName  string
}

// BirthDate enthält das Geburtsdatum einer Person
type BirthDate struct {
	Day   int
	Month int
	Year  int
}

// Profile enthält Informationen über eine Person, einschließlich Name, Geburtsdatum, Anzahl der Geschwister und Sternzeichen
type Profile struct {
	FullName       // eingebettete Struktur für den vollständigen Namen
	BirthDate      // eingebettete Struktur für das Geburtsdatum
	NumberOfSiblings byte
	ZodiacSign      rune
}

func main() {
	// Erstellen eines Profile-Objekts mit der Initialisierung der eingebetteten Strukturen
	var me = Profile{
		FullName: FullName{
			FirstName: "Luis",
			LastName:  "Allamand",
		},
		BirthDate: BirthDate{
			Day:   19,
			Month: 6,
			Year:  2008,
		},
		NumberOfSiblings: 0,   
		ZodiacSign:       'Z', 
	}

	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)

	me.NumberOfSiblings++

	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
