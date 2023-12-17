package main

// Go Standard Library:

import (
	"fmt"
	"sort"
	"strings"
)

func function1(n string) {
	fmt.Printf("Printing n in function 1: %v", n)
}

// Entry Point Function when running code
// Only 1 main function in every application
// Execute with `go run fileName.go`
func main() {

	fmt.Println("Exportable methods must begin with a capital letter.")

	var variableName string = "Strings are double quotes"

	var secondVariable = "Go can infer types"

	var thirdVariable string
	fmt.Println(variableName, secondVariable, thirdVariable)
	variableName = "new name 1"
	secondVariable = "new name 2"
	thirdVariable = "new name 3"
	fmt.Println(variableName, secondVariable, thirdVariable)

	weirdVar := "initialized"
	fmt.Println(weirdVar)

	var ageOne int = 20
	var ageTwo = 25
	ageThree := 30

	fmt.Println(ageOne, ageTwo, ageThree)

	var num1 int8 = 126
	var num2 uint = 25 // u int is Not negative

	var scoreOne float32 = 18.19
	var scoreTwo float64 = 12345.652
	fmt.Println(num1, num2, scoreOne, scoreTwo)
	fmt.Print("Hello, ")   // Print does not add a new line
	fmt.Print("World, \n") // \n adds a new line
	fmt.Println("World")   // adds a new line by default

	// PrintF format specifiers
	// %v for value, %q for quote (must be string); %T for type;
	// Format floats with %0.1f to round to tenths

	fmt.Printf("my age is %v and my name is %v \n", num1, num2)

	// Sprint saves the formated string
	var str = fmt.Sprintf("my age is %v and my name is %v \n", num1, num2)
	fmt.Println(str)

	// Arrays have fixed length in go
	var ages [3]int = [3]int{20, 25, 30}
	var agesShorthand = [3]int{20, 25, 30}
	fmt.Println(ages)
	fmt.Println(len(agesShorthand))

	// Slices are non-fixed length arrays
	scores := []int{10, 20, 30, 40, 50}
	fmt.Println(scores)
	scores[2] = 23
	x := append(scores, 89) // Doesn't mutate original, returns a new one
	fmt.Println(x)
	fmt.Println(scores[2:4]) // inclues:excludes

	greeting := "hello"
	fmt.Println(strings.Contains(greeting, "hello"))

	ages2 := []int{5, 6, 7, 8, 9}
	sort.Ints(ages2)
	fmt.Println(ages2)

	// For Loops
	xx := 0
	for xx < 5 {
		fmt.Println("xx is: ", xx)
		xx++
	}

	namez := []string{"Mario", "Luigi", "Yoshi"}
	for ii := 0; ii < len(namez); ii++ {
		fmt.Println("ii is: ", namez[ii])
	}

	for _, value := range namez {
		fmt.Printf("value is %v", value)
	}

	num := 5
	if num < 4 {
		fmt.Println("num is less than 4")
	} else if num == 5 {
		fmt.Println("num is 5")
	} else {
		fmt.Println("num is not 5")
	}

	function1("apple")
}
