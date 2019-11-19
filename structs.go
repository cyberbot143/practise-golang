package main

import (
	"fmt"

	"github.com/cyberbot143/go_crash_course/customStruct"
)

type Point struct {
	X float64
	Y float64
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Student struct {
	RollNumber int
	Name       string
}

func StructsDemo() {
	fmt.Println("****** Structs Demo *******")

	// structs are value Types

	var p Person // // All the struct fields are initialized with their zero value
	fmt.Println(p)

	// Declaring and initializing a struct using a struct literal
	p1 := Person{"SAI", "KRISHNA", 26}
	fmt.Println("Person1: ", p1)

	// Naming fields while initializing a struct
	p2 := Person{
		FirstName: "John",
		LastName:  "Snow",
		Age:       22,
	}
	fmt.Println("Person2: ", p2)

	// Uninitialized fields are set to their corresponding zero-value
	p3 := Person{FirstName: "Robert"}
	fmt.Println("Person3: ", p3)

	c := customStruct.Car{
		Name:       "Ferrari",
		Model:      "GTC4",
		Color:      "Red",
		WeightInKg: 1920,
	}

	// Accessing struct fields using the dot operator
	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)

	// Assigning a new value to a struct field
	c.Color = "Black"
	fmt.Println("Car: ", c)

	// instance of student struct type
	s := Student{11, "Jack"}

	// Pointer to the student struct
	ps := &s
	fmt.Println(ps)

	// Accessing struct fields via pointer
	fmt.Println((*ps).Name)
	fmt.Println(ps.Name) // Same as above: No need to explicitly dereference the pointer

	ps.RollNumber = 31
	fmt.Println(ps)

	// You can also get a pointer to a struct using the built-in new() function
	// It allocates enough memory to fit a value of the given struct type, and returns a pointer to it
	pEmp := new(customStruct.Employee)

	pEmp.Id = 1000
	pEmp.Name = "Sachin"

	fmt.Println(pEmp)

	// Two structs are equal if all their corresponding fields are equal.
	point1 := Point{3.4, 5.2}
	point2 := Point{3.4, 5.2}

	if point1 == point2 {
		fmt.Println("Point p1 and p2 are equal.")
	} else {
		fmt.Println("Point p1 and p2 are not equal.")
	}

}
