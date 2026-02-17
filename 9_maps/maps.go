package main

import "fmt"

// maps is basically a key value pair data structure.
// It is also known as hash table or dictionary in other programming languages.
// It is a collection of unordered key value pairs. The keys are unique and the values can be of any type.
// The keys are used to access the values in the map.
func main(){

	// m:=make(map[string] string)
	// m:=make(map[string] bool)
	m:=make(map[string] string)

	// m["name"]="ankush"

	// fmt.Println(m["name"])

	m["language"]="go lang"
	m["number"]="098783247834"

	fmt.Println(m["number"])
	fmt.Println(m["mobile"]) // if key does not exist it return 0 if it is int , and false for boolean 
	fmt.Println(len(m))

	// if you want to delete any key value 
	delete(m, "number")

	fmt.Println(m)

	// clear functiion use to clear the map 

	clear(m)

	fmt.Println(m)

	inbuilt()

}

func inbuilt(){

	// this is how we create a mao 
	data:=map[string] int {"ankush":1,"tyagi":2,"prateek":3}

	fmt.Println(data)

	// go return 2 value at a time in map e.g

	value , ok:=data["ankush"] // value here return the key's value 

	fmt.Println(value)
	fmt.Println(ok) // here it return value is available  or not return True/false

	if ok {
		fmt.Println("it is okay")
	}else{
		fmt.Println("it is not okay")
	}
}