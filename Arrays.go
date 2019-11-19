package main

import "fmt"

func ArraysDemo() {
	fmt.Println("**** Arrays & slices Demo ****")
	var x [5]int
	var xLen = len(x)

	//iterating using for loop
	for i := 0; i < xLen; i++ {
		fmt.Println(i)
		x[i] = i
	}

	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	//iterating using for loop
	for index, value := range daysOfWeek {
		fmt.Printf("Day %d of week = %s\n", index, value)
	}

	a := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for _, value := range a {
		sum = sum + value
	}

	fmt.Printf("Sum of all the elements in array %v = %f", a, sum)

	fmt.Println()
	fmt.Printf(" x %d", x)
	fmt.Println()

	a1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
	a2 := a1 // A copy of the array `a1` is assigned to `a2`

	a2[1] = "German"

	fmt.Println("a1 = ", a1) // The array `a1` remains unchanged
	fmt.Println("a2 = ", a2)

	// slices

	var s = make([]int, 100) // var s =[]int
	for i := 0; i < 100; i++ {
		s[i] = i * 10
	}

	fmt.Println(s)

	var nilSlice []string
	// Appending to a nil slice
	nilSlice = append(nilSlice, "Cat", "Dog", "Lion", "Tiger")

	fmt.Printf("s = %v, len = %d, cap = %d\n", nilSlice, len(nilSlice), cap(nilSlice))

	// Creating a slice from the array and we can also create from another slice
	// but has an capcity of parent Array
	var weekend []string = daysOfWeek[4:]

	fmt.Println(weekend)

	src := []string{"Sublime", "VSCode", "IntelliJ", "Eclipse"}
	dest := make([]string, 2)

	numElementsCopied := copy(dest, src)

	fmt.Println("src = ", src)
	fmt.Println("dest = ", dest)
	fmt.Println("No. of elems copied 4m src 2 dest = ", numElementsCopied)

	slice3 := append(src, dest...)

	fmt.Println(slice3)

	fmt.Println("All ABout Slice & Arrays end......")

}
