package main

import (
	"errors"
	"fmt"
	"os"
)

// These exercises test your knowledge of functions in Go and how to use them. Solutions are available in the exercise_solutions directory in the Chapter 5 repository.

// 1. The simple calculator program doesn’t handle one error case: division by zero. Change the function signature for the math operations to return both an int and an error. In the div function, if
// the divisor is 0, return errors.New("division by zero") for the error. In all other cases, return nil. Adjust the main function to check for this error.
func add(i int, j int) (int, error) { return i + j, nil }

func sub(i int, j int) (int, error) { return i - j, nil }

func mul(i int, j int) (int, error) { return i * j, nil }

func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

// 2. Write a function called fileLen that has an input parameter of type string and returns an int and an error.
// The function takes in a filename and returns the number of bytes in the file.
// If there is an error reading the file, return the error. Use defer to make sure the file is closed properly.
func fileLen(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	stats, err := file.Stat()
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return stats.Size(), nil
}

// 3. Write a function called prefixer that has an input parameter of type string and returns a function that has an input parameter of type string and returns a string.
// The returned function should prefix its input with the string passed into prefixer.
// Use the following main function to test prefixer:
func prefixer(inp string) func(string) string {

	ret := func(p string) string {
		return inp + " " + p
	}
	return ret
}

// func main() {
//
// }
func main() {
	//ex1
	fmt.Println(add(1, 1))
	fmt.Println(sub(1, 1))
	fmt.Println(mul(2, 5))
	fmt.Println(div(3, 0))
	fmt.Println(div(21, 7))
	//ex2
	y, err := fileLen("./filename.txt")
	fmt.Println(y, err)
	//ex3
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria

}
