package main

import "fmt"

func main() {
	age := 16

	if age < 16 {
		fmt.Println("not adult")
	} else if age == 16 {
		fmt.Println("person is 16")
	} else {
		fmt.Println("adult")
	}
}
