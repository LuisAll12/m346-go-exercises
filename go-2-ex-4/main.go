package main

import "fmt"
// TODO: declare a type for Student (with first and last name)
type Student struct{
	FirstName string
	LastName string
}
// TODO: declare a type for Class (consisting of multiple students)
type Class struct{
	ClassName string
	Students []Student
}
func main() {
	student1 := Student{FirstName: "Alice", LastName: "Schmidt"}
	student2 := Student{FirstName: "Bob", LastName: "Müller"}
	student3 := Student{FirstName: "Charlie", LastName: "Weber"}

	class1 := Class{
		ClassName: "Mathematik 101",
		Students:  []Student{student1, student2},
	}

	class2 := Class{
		ClassName: "Informatik 202",
		Students:  []Student{student2, student3},
	}

	modules := map[string][]Class{
		"Mathematik": {class1},
		"Informatik": {class2},
	}

	fmt.Println("Studenten und Klassen:")
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(student3)

	fmt.Println("\nKlassen und ihre Studenten:")
	fmt.Println(class1)
	fmt.Println(class2)

	fmt.Println("\nModule und die besuchenden Klassen:")
	for module, classes := range modules {
		fmt.Printf("Modul: %s\n", module)
		for _, class := range classes {
			fmt.Printf("  Klasse: %s\n", class.ClassName)
			for _, student := range class.Students {
				fmt.Printf("    Student: %s %s\n", student.FirstName, student.LastName)
			}
		}
	}
}
