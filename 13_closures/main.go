package main

import "fmt"

func main() {

	// here closures is the same JS concept i.e inside function always store the scope of outer function variable e.g
	increment:=counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(sum())
}


func counter() func() int {
	count:=0

	return func() int {
		count+=1
		return count
	}
}


// this is also the closures example 
func sum() int {
	a:=1

	add := func() int {
		b:=10
		c:=a+b
		return c
	}
	
	return add()
}