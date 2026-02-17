package main

import "fmt"

func main() {

	// here closures is the same JS concept i.e inside function always store the scope of outer function variable e.g
	increment:=counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}


func counter() func() int {
	count:=0

	return func() int {
		count+=1
		return count
	}
}
