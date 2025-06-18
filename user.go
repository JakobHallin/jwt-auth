package main

import "fmt"

type User struct{
	//Id int  //will not use it now will just have it simple with name and password
	name string
	password string //will look how it will be encoded with jws token
}

var users = map[string]User{} //map of the users
var storage = "user.csv" //

func loadUsers(){ //https://www.geeksforgeeks.org/go-language/how-to-read-a-csv-file-in-golang/
	//open file load users
	file, err := os.Open(storage) //https://pkg.go.dev/os
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	reader = csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
		retrurn err
	}
	for _, eachrecord := range records{
		fmt.Println(eachrecord)
	// look at csv how it store to later set in the right
	}
	return nil
	//loop and then set to the map the user should be dynamic cause its go
}

func saveUsers(user User){
	//open file
	csvFile, err := os.Create("temp.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvFile.Close()
	//write user to the file
	writer := csv.NewWriter(csvFile)
	writer.Flush()
	return write.Write([]string{user.name, user,password)
}

//handler in signup should call this function
func createUser(name, password string) {
	//hash the password
	user := User{name, password} //id unsure
	//call saveUser(user) to save it
	users:= user[name]
	return saveUser(user)
}
func auth(name, password string){
	//cheack if name exist
	//hash the password
	// compare it
}

