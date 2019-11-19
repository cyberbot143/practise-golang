package main

import "fmt"

func MapsDemo() {

	//Maps are reference types. When you assign a map to a new variable,
	//they both refer to the same underlying data structure.
	//Therefore changes done by one variable will be visible to the other

	var m map[string]int
	fmt.Println(m)
	if m == nil {
		fmt.Println("m is nil")
	}
	// Attempting to add keys to a nil map will result in a runtime error
	// m["one hundred"] = 100

	// Initializing a map using the built-in make() function
	var map1 = make(map[string]int)

	if map1 == nil {
		fmt.Println("map1 is nil")
	} else {
		fmt.Println("map1 is not nil")
	}

	// make() function returns an initialized and ready to use map.
	// Since it is initialized, you can add new keys to it.
	map1["one hundred"] = 100
	fmt.Println(m)

	//creating map2 using map Literal

	var map2 = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5, // Comma is necessary
	}

	//adding properties to map

	map2["six"] = 6 //inert or update

	//deleting properties from a map

	delete(map2, "two")

	fmt.Println(map2)

	value, isValuePresent := map2["six"] // 6, true

	fmt.Println(value, isValuePresent)

	value, isValuePresent = map2["seven"] // "", false

	fmt.Println(value, isValuePresent)

	//iterating through the map

	for name, value := range map2 {
		fmt.Println(name, value)
	}

}
