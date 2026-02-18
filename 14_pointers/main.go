package main

import "fmt"

func main() {

	num := 5
	changeNumber(&num) // & is used to pass the address
	fmt.Println(num)
}

// call by reference
func changeNumber(num *int) int { // here if we want to receive the address the * is used

	// num := 10 // it cause error because we need to update the address not the value
	*num = 10
	return *num
}
