package main

import "fmt"

// in go only for loop is present
func main() {

	//implementing while loop using for
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println("======================")

	// classic for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("======================")

	for i := range 4 {
		fmt.Println(i)
	}
}
