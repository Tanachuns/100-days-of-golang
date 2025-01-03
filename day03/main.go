package main

import (
	"fmt"
)

func main() {
	/*Write a program that defines a variable named greetings of type slice of strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
	//Create a subslice containing the first two values; a second subslice with the second, third, and fourth values;
	 and a third subslice with the fourth and fifth values. Print out all four slices.*/
	list := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	sub1 := list[0:2]
	sub2 := list[2:]
	fmt.Println(sub1)
	fmt.Println(sub2)
	//   Write a program that defines a string variable called message with the value "Hi  and " and prints the fourth rune in it as a character, not a number.
	s := "Hi  and "
	fmt.Println(string(s[4]))
	//   Write a program that defines a struct called Employee with three fields:
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	//   firstName, lastName, and id. The first two fields are of type string, and the last field (id) is of type int.
	//Create three instances of this struct using whatever values you’d like.
	var e = Employee{
		id: 1,
	}

	var e1 = Employee{
		firstName: "asdaaa",
		lastName:  "asdd",
	}

	// Initialize the first one using the struct literal style without names,
	// the second using the struct literal style with names, and the third with a var declaration.
	var e2 = Employee{}
	e2.firstName = "TON"
	e2.lastName = "TANA"

	fmt.Println(e)
	fmt.Println(e1)
	fmt.Println(e2)
	// Use dot notation to populate the fields in the third struct. Print out all three structs.*/
}
