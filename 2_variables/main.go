package main

import "fmt"

func main() {
	println("Hello, World!")


	// syntax to define variables 

	var name string = "Ankush" // it explicitly specifies the type of the variable as string
	// or

	var name1 = "Ankush" // it dynamically infers the type of the variable based on the assigned value	
	// or

	name2 := "Ankush" // it is a shorthand syntax for declaring and initializing a variable in one line. It is only used within functions and cannot be used at the package level. The type of the variable is inferred from the assigned value.

	fmt.Println(name , name1 , name2)
}
