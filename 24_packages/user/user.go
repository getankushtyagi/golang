package user

// here alse we use capital letter for the struct name and its fields because we want to use this struct in other packages as well so we need to make it open scope
type User struct {
	Name string
	Email string
}