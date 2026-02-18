package main

import (
	"fmt"
	"time"
)

type order struct {
	id        int
	name      string
	status    string
	createdAT time.Time // inbuilt function and their precision is very high specially in nanosecond
}

// how to create a constructor in gollang using struct
func NewOrder(Id int, Name string, Status string) *order {
	return &order{
		id:     Id,
		name:   Name,
		status: Status,
	}
}

// create a function which runs on struct
func (o *order) changeStatus(status string) {
	o.status = status
}

//create a func which return the struct value

func (o order) getName() string { // here we do not need to pass the * because it only return the value not changing
	return o.name
}

func main() {

	myorder := order{
		id:     1,
		name:   "laptop",
		status: "placed",
	}

	// we can also add another valu in struct
	myorder.createdAT = time.Now()

	fmt.Println("struct order===", myorder)

	// we can also print any value in struct
	fmt.Println(myorder.name)

	myorder.changeStatus("confirmed")

	fmt.Println(myorder)

	fmt.Println(myorder.getName())

	// how we use struct as constructor

	o := NewOrder(1, "mobile", "deleivered")
	fmt.Println(o)
}
