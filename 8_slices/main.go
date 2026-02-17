package main

import (
	"fmt"
	"slices"
)

func main() {

	// slice only declare also we need to define the type just like array but size

	var nums []int
	fmt.Println(len(nums))

	// fmt.Println("========")

	// other way to define
	data := make([]int, 3, 5)

	fmt.Println(len(data)) // length
	fmt.Println(cap(data)) // capacity
	fmt.Println(data)

	// fmt.Println(len(data))
	// fmt.Println(cap(nums)) // it shows the capacity of array how much element can be filled

	//capacity is dynamic it will automaticaaly increase

	data = append(data, 1)
	data = append(data, 2)
	data = append(data, 3)

	// also put value using index

	data[4]=5
	fmt.Println("========")
	fmt.Println(data)
	fmt.Println(cap(data)) // here capacity increase to double i.e 10 

	data1 := make([]int , 0, 5)
	fmt.Println("for empty slice ", data1)

	fmt.Println("slicing =========")

	slicing()

}



func slicing() {

	//slicing works same as python 
	nums := []int{1, 2, 3}
	fmt.Println(nums[0:2])

	//reverse 
	slices.Reverse(nums)
	fmt.Println(nums)

	//equals

	fmt.Println(slices.Equal(nums, []int{3, 2, 1}))

	
}


