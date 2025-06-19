package main
import (
	"fmt"
	"encoding/csv"
	//"errors"
	"os"
	"log"
)
type User struct{
	//Id int  //will not use it now will just have it simple with name and password
	Name string
	Password string //will look how it will be encoded with jws token
}
var users = map[string]User{} //map of the users
var storage = "temp.csv" //

func loadUsers() error{ //https://www.geeksforgeeks.org/go-language/how-to-read-a-csv-file-in-golang/
	//open file load users
	file, err := os.Open(storage) //https://pkg.go.dev/os
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return err
	}
	for _, eachrecord := range records{
		fmt.Println(eachrecord)
	// look at csv how it store to later set in the right
		if len(eachrecord) < 2 {
			log.Printf("Skipping invalid record: %v\n", eachrecord)
			continue
		}
		users[eachrecord[0]] = User{
			Name:     eachrecord[0],
			Password: eachrecord[1],
		}
	}
	return nil
}
func saveUser(user User) error{
	//open file
	csvFile, err := os.OpenFile(storage, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//absPath, _ := filepath.Abs(csvFile) 
	fmt.Println("CSV file created at:", csvFile.Name())
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		return err
	}
	defer csvFile.Close()
	//write user to the file
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()
	return writer.Write([]string{user.Name, user.Password})
	//return writer.Write("test")
}

//handler in signup should call this function
func createUser(name, password string) error{
	//hash the password
	user := User{Name: name, Password: password}
	//call saveUser(user) to save it
	users[name]= user
	return saveUser(user)
}
func auth(name, password string){
	//cheack if name exist
	//hash the password
	// compare it
}

