package main

import "fmt"

//genereic helps us to. create a function that can work with any type, without specifying the exact type in advance.
//This is achieved using type parameters, which are placeholders for the types that will be used when the function is called. By using generics,
//we can write more flexible and reusable code, as the same function can operate on different types without needing to be rewritten for each specific type.

// but this is not genric here it only accept int slice 
func printSlice(items[] int ){
	fmt.Println(items)

	for _ , item:=range(items){
		fmt.Println(item)
	}
}


//now lets create the genereic 

func print[T any](items[]T){ // but it all all type of data types which is wrong practise 
	for _ , i:=range(items){
		fmt.Println(i)
	}
}

// here it only accept integer and  string 
func printIntString[T int|string](items[] T){
	for _ , i := range(items){
		fmt.Print(i)
	}
}

//how we can implement generic in struct
type stack[T any] struct{
	element []T
}

func main(){

	printSlice([]int{1,2,3,4,5,5})

	print([]int{2,3,4,})
	print([]string{"aank","sbdsbd"})

	// printIntString([]bool{true}) // getting error 
	printIntString([]string{"true"}) 


	mystack:= stack[string]{
		element: []string{"golang"},
	}
	fmt.Println(mystack)
}
