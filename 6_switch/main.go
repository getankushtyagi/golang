package main

import (
	"fmt"
	"time"
)


func main (){

	i:=1

	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}


	switch time.Now().Weekday(){
	case time.Friday:
		fmt.Println("friday")
	case time.Monday:
		fmt.Println("monday")
	}
}
