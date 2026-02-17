package main

import "fmt"

func main() {

	//go can return multiple values from a function
	fmt.Println(getlanguages())

	// OR

	val1, val2, val3 := getlanguages()
	println(val1, val2, val3)

	val4, val5, _ := getlanguages() // if you do not use all value simply pass _ to remove the compilation error
	println(val4, val5)

}

func getlanguages() (string, string, int) {
	return "java", "php", 8
}
