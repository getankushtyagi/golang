package main

import "fmt"

// in go enums is not a type we will use this using const

type UserStatus string

type OrderStatus int

const (
	Recieved OrderStatus = iota // so now reveived value become 0 also iota is only integer
	Confirmed
	Prepared
	Deleiverd
)

// we can also create enums for string also but iota is not used then

const (
	Pending UserStatus = "pending"
	Onboard            = "onboard "
)

func changeStatus(status OrderStatus) {
	fmt.Println("status change to ", status)
}
func changeUserStatus(status UserStatus){
	fmt.Println("user status is now ", status)
}

func main() {
	changeStatus(Confirmed) // 1
	changeStatus(Recieved)  // 0
	changeStatus(Deleiverd) // 3
	changeUserStatus(Pending)
}
