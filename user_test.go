package main
import (
	"testing"
	"os"
	"log"
	//"golang.org/x/crypto/bcrypt"
	"database/sql"
)
func setUpTestStore(t *testing.T) *UserStore{
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		log.Fatal("TEST_DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("FAILED to open database", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}
	store := &UserStore{db: db}
	if err := store.InitSchema(); err != nil {
		log.Fatalf("Failed to initialize schema: %v", err)
	}
	// Clean users before each test
	_, err = db.Exec(`DELETE FROM users`)
	if err != nil {
		t.Fatalf("Failed to clear users table: %v", err)
	}
	return store
}

func TestCreateUser(t *testing.T){

	store:=setUpTestStore(t)
	err := store.CreateUser("testuser", "word")
	if err != nil {
		t.Errorf("Failed to create user: %v", err)
	}
}
/* this function dont exist no more
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
}*/
func TestAuth(t *testing.T){
	store:=setUpTestStore(t)
	err := store.CreateUser("test", "password")
	if err != nil {
		t.Errorf("Failed to create user: %v", err)
	}
	//case 1 correct
	if !store.Auth("test","password"){
		t.Error("expected to succsed")
	}
	//case 2 wrong
	if store.Auth("test", "falsepassword"){
		t.Error("expected to fail wrong password")
	}
	//case 3 non existing
	if store.Auth("nouser","wrongpassword"){
		t.Error("expected to fail non existing user")
	}


}
