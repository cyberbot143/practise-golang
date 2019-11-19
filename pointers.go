package main

import "fmt"

func Pointers() {
	var a = 5.67
	var p = &a

	// creating ptr using new keyword
	ptr := new(int) // Pointer to an `int` type
	*ptr = 100

	fmt.Printf("Ptr = %#x, Ptr value = %d\n", ptr, *ptr)

	fmt.Println("Value stored in variable a = ", a)
	fmt.Println("Address of variable a = ", &a)
	fmt.Println("Value stored in variable p = ", p)

	fmt.Println("*p = ", *p) //derefrencing or indirecting (basically accessing the value)

	// Changing the value stored in the pointed variable through the pointer
	*p = 2000

	fmt.Println("a (after) = ", a)

	var pp = &p //pointer to pointer

	// Dereferencing a pointer to pointer
	fmt.Println("*pp = ", *pp)
	fmt.Println("**pp = ", **pp)

	var p1 = &a
	var p2 = &a

	if p1 == p2 {
		fmt.Println("Both pointers p1 and p2 point to the same variable.")
	}

}
