package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
	s := []int{}
	for i := 0; i < 100; i++ {
		s = append(s, rand.IntN(100))
	}
	fmt.Println(s)
	// Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules:
	for _, item := range s {
		// If the value is divisible by 2, print “Two!”
		// If the value is divisible by 3, print “Three!”

		// IIf the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.

		// Otherwise, print “Never mind”.
		if item%2 == 0 && item%3 == 0 {
			fmt.Println("Six!")
		} else if item%2 == 0 {
			fmt.Println("Two!")
		} else if item%3 == 0 {
			fmt.Println("Three!")
		} else {
			fmt.Println("Never mind")
		}

	}

	// Start a new program. In main, declare an int variable called total. Write a for loop that uses a variable named i to iterate from 0 (inclusive) to 10 (exclusive). The body of the for loop should be as follows:
	total := 0
	// total := total + i
	// fmt.Println(total)
	for i := 0; i <= 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	// After the for loop, print out the value of total.  What is printed out? What is the likely bug in this code?
	fmt.Println(total)

}
