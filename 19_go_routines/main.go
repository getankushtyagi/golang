package main

import (
	"fmt"
	"time"
)

func main(){

	// if you the output of this all the value printing squentially because it is blocking code 
	for i:=range(10){
		task(i)
	}

	// to do a task simultaneously we use go routine to do task using multithreading 
	fmt.Println("\n\n\nseeing go routine  example\n\n\n")

	// also it run siimultaneouslhy and before completion of task our code ends so we need to hold it for second and for that we use sleep 
	
	// for i:=range(10){
	// 	go task(i)
	// }


	// we can also use closures to see the go routiner examples 
	for i:= range(10){
		go func (i int){
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second*2)



}

func task(i int){
	fmt.Println("printing index of task ", i)
}