package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
}

// see the  example of  mutex 

type postWithMutex struct {
	views int
	mu sync.Mutex // we add a mutex field to the struct, this will be used to protect the views field from concurrent access
}


// this is how attch the function to the struct, we can call this function on the struct instance, and it will modify the views field of the struct, but as mentioned before,
// this is not thread safe, if we have multiple go routine accessing the same post struct, we might end
func (p *post) inc(){
	p.views+=1
}

//mutex function to increment the views count, we will lock the mutex before modifying the views field, and then unlock it after we are done, this will ensure that only one go routine can access the views field at a time,
//  and it will prevent race condition from happening, because if one go routine is modifying the views field, other go routine will be blocked until the first go routine is done, and then they can access the views field one by one.

func (p *postWithMutex) inc(){
	
	p.mu.Lock() // we lock the mutex before modifying the views field
	p.views+=1
	p.mu.Unlock() // we unlock the mutex after we are done
}

func main(){

	// this is how we normally do a thing and use struct and their attach method, but this is not thread safe, if we have multiple go routine accessing the same post struct,
	// we might end up witha race condition,where the views count might not be accurate, because multiple go routine might read the same value of views and then increment it,
	// and then write it back, which will result in lost updates.

	myPost:=post{views:0}
	myPost.inc()
	fmt.Println(myPost.views)

	// how race condition can happen, we will create multiple go routine that will increment the views count of the same post struct, and we will see that the final views count is not accurate.

	// for i:=0;i<1000;i++{
	// 	go myPost.inc()
	// }

	// we need to wait for all the go routine to finish before we can print the final views count,
	// otherwise we might print the views count before all the go routine have finished incrementing it, which will result in an inaccurate views count.

	// we can use time.Sleep to wait for all the go routine to finish, but this is not a good way to do it,
	// because we don't know how long it will take for all the go routine to finish, and we might end up waiting for too long or not long enough.

	// time.Sleep(time.Second)

	// we can use sync.WaitGroup to wait for all the go routine to finish, this is a better way to do it, because we can specify how many go routine we are waiting for,
	// and we can wait for them to finish without worrying about how long it will take.

	var wg sync.WaitGroup
	wg.Add(1000) // we are waiting for 1000 go routine to finish

	for i:=0;i<1000;i++{
		go func(){
			myPost.inc()
			wg.Done() // we call Done to indicate that this go routine has finished
		}()
	}

	wg.Wait() // we wait for all the go routine to finish

	fmt.Println(myPost.views) // we print the final views count, but this is still not accurate
	 
	// to avoid this we use mutex, a mutex is a lock that can be used to protect shared resources from concurrent access,
	// we can use a mutex to protect the views field of the post struct, so that only one go routine can access it at a time.

	myPostWithMutex:=postWithMutex{views:0}

	var wg2 sync.WaitGroup
	wg2.Add(1000) // we are waiting for 1000 go routine to finish

	for i:=0;i<1000;i++{
		go func(){
			myPostWithMutex.inc() // we call the inc function that uses mutex to increment the views count
			wg2.Done() // we call Done to indicate that this go routine has finished
		}()
	}

	wg2.Wait() // we wait for all the go routine to finish

	fmt.Println(myPostWithMutex.views) // we print the final views count, this should be accurate because we are using mutex to protect the views field from concurrent access
	
}
