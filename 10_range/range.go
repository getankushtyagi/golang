package main

import "fmt"

func main(){
	nums:=[]int {6,7,8,9,10}


	// this one way to do it 
	// for i:=0 ; i<len(nums)-1 ; i++ {
	// 	fmt.Println(nums[i])
	// }


	// using range 

	for i, num:= range nums{ // here i return the index and num return the value 

		fmt.Println(i)
		fmt.Println(num)

		fmt.Println("===")
	}

	sum(nums) // iterate for the sum of the array
	 rangeInString()
}


func sum(nums []int){
	sum:=0
	for i,val:= range(nums){
		sum+=val
		fmt.Println(i)
		fmt.Println(val)
	}
	fmt.Println("Sum:", sum)
}

func rangeInString(){
	str:="ankush tyagi"

	for i,code:= range str{
		fmt.Printf("index: %d, code: %d, char: %c\n", i, code, code)
	}
}