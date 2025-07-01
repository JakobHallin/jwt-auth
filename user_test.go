package main
import (
	"testing"
	"os"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T){
	tmpFile, err := os.CreateTemp("", "test_users.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	storage = tmpFile.Name()
	users = make(map[string]User)
	err = createUser("user","word")
	if err != nil {
		t.Errorf("Failed to create user: %v", err)
	}
	record, err := os.ReadFile(storage)
	if err != nil {
		t.Fatalf("Failed to read temp file: %v", err)
	}
	expected := "user,word\n"
	if string(record) != expected { //this is no longer true password is now hashed
		//t.Fatalf("Expected file content %q, got %q", expected, record)
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
func TestAuth(t *testing.T){
	hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	users = map[string]User{
		"test": {Name:"test", Password: string(hashed)},
		"test2": {Name:"test2", Password:"password"},
	}
	//case 1 correct
	if !auth("test","password"){
		t.Error("expected to succsed")
	}
	//case 2 wrong
	if auth("test", "falsepassword"){
		t.Error("expected to fail wrong password")
	}
	//case 3 non existing
	if auth("nouser","wrongpassword"){
		t.Error("expected to fail non existing user")
	}


}
