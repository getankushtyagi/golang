package main

import "fmt"

func main() {

	// result:=make(chan int)

	// result<-10 // this is deadlock if try to print this

	// fmt.Println(result)

	result := make(chan int) // this is the integer channel

	go add(result, 4, 5) // here no dead lock is present because it is running on go routines 

	// see here we do not need wait group or time sleep beacuse sum := <-result this is the blocking code 
	sum := <-result // this is how we receive the value from the channels 
	fmt.Println(sum)


	// lets do a similar task of wiat group 
	done :=make(chan bool) // create a channel 
	go process(done) // it will call the function but there may be a chance it run before the process function run to ensure this we need to block this until the output is coming 
	// so here we sue the blocking 
	<-done // this ensure the output from the process func and then runn




	// as we see above we have face deadlock because we are trying to send the value in channel but no one is receiving it 
	// and also we are trying to receive the value from channel but no one is sending it so this is the deadlock

	// to avoid this we can use go routine and also we can use buffered channel

	bufferedChannel := make(chan int, 2) // this is the buffered channel with capacity of 2

	bufferedChannel <- 10 // this is how we send the value in buffered channel 
	bufferedChannel <- 20 // this is how we send the value in buffered channel

	fmt.Println(<-bufferedChannel) // this is how we receive the value from buffered channel 
	fmt.Println(<-bufferedChannel) // this is how we receive the value from buffered channel
}

func add(result chan int, num1 int, num2 int) {
	sum := num1 + num2
	result <- sum // this is how we add the value in channels 
}


func process(done chan bool ){

	// here we write this function because defer call at the last of fucntion and ensure sucessfully running the code 
	defer func(){
		fmt.Println("processing ....")
	}()

	done <- true

}