package main

import "fmt"

const value int = 1

func main() {
	var i int = 20
	var f float32 = float32(i)

	fmt.Println(i)
	fmt.Println(f)

	i = value
	f = float32(value)

	fmt.Println(i)
	fmt.Println(f)

	var b byte = 11
	var smallI int32 = 12
	var bigI uint64 = 13

	fmt.Println(b + 1)
	fmt.Println(smallI + 1)
	fmt.Println(bigI + 1)
}
