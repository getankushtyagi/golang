package main


const age = 25 // we can also define variables outside the function as well
func main() {
	const pi = 3.14 // it is a constant variable that cannot be changed after it is defined. It is used to represent the value of pi in mathematics.

	// this will cause a compile-time error because pi is a constant and cannot be reassigned a new value.
	// pi = 3.14159 

	println(pi)
	println(age)


	// we can also group multipl constants together using parentheses, like this:

	const(
		country = "India"
		city = "Delhi"
	)

	println(city)
	println(country)
}