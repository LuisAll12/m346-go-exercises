package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Mathematik",
		117: "Informatik",
		346: "Physik",
	}

	fmt.Println("Modul 104:", modules[104]) 
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346]) 

	// TODO: Löschen eines Moduls (z.B. Modul 117)
	delete(modules, 117)

	// TODO: Hinzufügen eines neuen Moduls (z.B. Modul 200: Chemie)
	modules[200] = "Chemie"

	// TODO: Ersetzen eines Moduls (z.B. Modul 104 durch Biologie)
	modules[104] = "Biologie"

	// Ausgabe der aktualisierten Map
	fmt.Println(modules)
}
