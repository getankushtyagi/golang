package main

import (
	login "auth/credential" // here auth is the module name and credential is the package name and we are giving it an alias login so that we can use it in our code
	user "auth/user"        // here auth is the module name and user is the package name and we are giving it an alias user so that we can use it in our code
	"fmt"
)

func main() {
	login.Login() // we can use this function because it is open scope



	// so this how we can use the struct from another package and we can also use the fields of that struct because they are also open scope
	user:= user.User{
		Name: "ankush Tyagi",
		Email: "ankush.tyagi@example.com",
	}

	fmt.Println(user.Email,"\n", user.Name)
}
