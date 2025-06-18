package main

import "fmt"

type User struct{
	Id int
	name string
	password string //will look how it will be encoded with jws token
}

var users = map[int]User{} //map of the users
var storage = "user.csv"

func loadUsers(){
	//open file load users
	//loop and then set to the map the user should be dynamic cause its go
}

func saveUsers(user User){
	//open file
	//write user to the file
}

//handler in signup should call this function
func createUser(name, password string) { //auto increment for id so depending on how i build i should have the db handle that if use sql database
	//hash the password
	//user := User{name, password} //id unsure
	//call saveUser(user) to save it
}
func auth(name, password string){
	//cheack if name exist
	//hash the password
	// compare it
}

