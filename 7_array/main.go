package main

import "fmt"

func main(){

	var nums[4] int

	// array length 
	fmt.Println(len(nums))

	// how to add value in a particular index 
	nums[0]=1
	fmt.Println(nums[0])

	// by defualt all value is 0 if type int 
	fmt.Println(nums)

	// same for bool , budefault all value is false
	var data[4] bool
	fmt.Println(data)
}
