package main
import (
	//"fmt"
	//"encoding/csv"
	//"errors"
	//"os"
	//"log"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
)
type User struct{
	//Id int  //will not use it now will just have it simple with name and password
	Name string
	Password string //will be hashed
}
type UserStore struct{
	db *sql.DB
}

//handler in signup should call this function
func (s *UserStore) CreateUser(name, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err = s.db.Exec(`INSERT INTO users (username, password) VALUES ($1, $2)`, name, string(hashed))
	return err
}


func (s *UserStore) Auth(name, password string) bool {
	var hashed string
	err := s.db.QueryRow(`SELECT password FROM users WHERE username = $1`, name).Scan(&hashed)
	if err != nil {
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
}
func (s *UserStore) InitSchema() error {
	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		username TEXT PRIMARY KEY,
		password TEXT NOT NULL
	);
	`)
	return err
}
