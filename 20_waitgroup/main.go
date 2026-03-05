package main

import (
	"fmt"
	"sync"
)

//wait group helps to hold the code until all the go routines end we do not need to sleep the code i.e wrong practise

func main(){
	var wg sync.WaitGroup // to use wait group we use the sync package inside sync we have a wait group function 

	for i:=range(10){
		wg.Add(1) // we need to add one delta inside the wait group which helps to track the task is running on some thread 
		go task(i, &wg)
	}

	wg.Wait() // it will automatically realize that now all the task is completed in all thread 
}


func task(i int , wg *sync.WaitGroup){ // we need to accept one paraamenter so that once the task is complete it will remove the delta we have added earlier
	defer wg.Done() // here we defer because it will only run once all the task is completed inside th function 
	fmt.Println("printing the index", i)
}
