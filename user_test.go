package main
import (
	"testing"
	"os"
)

func TestCreateUser(t *testing.T){
	err := createUser("user","word")
	if err != nil {
		t.Errorf("Failed to create user: %v", err)
	}
}
func TestLoadUser(t *testing.T){
	//Create a temporary CSV file
	tmpFile, err := os.CreateTemp("", "test_users.csv")
	if err != nil {
        	t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	// write to the file the user
	_, err = tmpFile.WriteString("testuser,password123\n")
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()
	 //Set the global storage so it dont use the orher file
	storage = tmpFile.Name()
	// try to call the loadusers
	users = make(map[string]User)
	if err := loadUsers(); err != nil {
		t.Fatalf("loadUsers failed: %v", err)
	}
	// validate
	user, ok := users["testuser"]
	if !ok {
	t.Fatal("Expected user 'testuser' not found")
	}
	if user.Password != "password123" {
        	t.Errorf("Expected password 'password123', got '%s'", user.Password)
	}
}
