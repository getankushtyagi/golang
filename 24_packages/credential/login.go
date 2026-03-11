package auth



// in go we use function name starting withn small then it is close scope i.e we cannot use it outside the package


func login(){
	println("Login successful")
}


// in go we use function name starting with capital then it is open scope i.e we can use it outside the package

func Login(){ // now we can use this function outside the package see the main.go file
	println("Login successful")
}