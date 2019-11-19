package main

import (
	"fmt"
	"math"
	"math/rand"
	str "strings" // Package Alias
	"time"

	"github.com/cyberbot143/go_crash_course/numbers"
	"github.com/cyberbot143/go_crash_course/strings"
	"github.com/cyberbot143/go_crash_course/strings/greeting" // Importing a nested package
)

//calc is the function name which accepts two integers num1 and num2
//(int, int) says that the function returns two values, both of integer type.
func calc(num1 int, num2 int) (int, int) {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("cyberbot143"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
	sum := num1 + num2
	diff := num1 - num2
	return sum, diff
}

func main() {
	fmt.Println("Hello Main.go")

	a := 20
	fmt.Println(a)

	const b = 10
	fmt.Println(b)

	//declaring a integer variable x
	var x int
	x = 3                //assigning x the value 3
	fmt.Println("x:", x) //prints 3

	//declaring a integer variable y with value 20 in a single statement and prints it
	var y int = 20
	fmt.Println("y:", y)

	//declaring a variable z with value 50 and prints it
	//Here type int is not explicitly mentioned
	var z = 50
	fmt.Println("z:", z)

	//Multiple variables are assigned in single line- i with an integer and j with a string
	var i, j = 100, "hello"
	fmt.Println("i and j:", i, j)

	var k int
	for k = 1; k <= 5; k++ {
		fmt.Println(k)
	}

	//calls the function calc with x and y an d gets sum, diff as output
	sum, diff := calc(10, 20)
	fmt.Println("Sum", sum)
	fmt.Println("Diff", diff)

	// Finding the Max of two numbers
	fmt.Println(math.Max(73.15, 92.46))

	// Calculate the square root of a number
	fmt.Println(math.Sqrt(225))

	// Printing the value of `𝜋`
	fmt.Println(math.Pi)

	// Epoch time in milliseconds
	epoch := time.Now().Unix()
	fmt.Println(epoch)

	// Generating a random integer between 0 to 100
	rand.Seed(epoch)
	fmt.Println(rand.Intn(100))

	var numbers [3]string //Declaring a string array of size 3 and adding elements
	numbers[0] = "One"
	numbers[1] = "Two"
	numbers[2] = "Three"
	fmt.Println(numbers[1])   //prints Two
	fmt.Println(len(numbers)) //prints 3
	fmt.Println(numbers)      // prints [One Two Three]

	directions := [...]int{1, 2, 3, 4, 5} // creating an integer array and the size of the array is defined by the number of elements
	fmt.Println(directions)               //prints [1 2 3 4 5]
	fmt.Println(len(directions))          //prints 5

	//Executing the below commented statement prints invalid array index 5 (out of bounds for 5-element array)
	//fmt.Println(directions[5])

}
