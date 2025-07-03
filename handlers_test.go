package main
import (
	//"bytes"
	"os"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"database/sql"

)
func setupTestServer(t *testing.T) *Server {
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		t.Fatal("TEST_DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal("Failed to open DB:", err)
	}

	store := &UserStore{db: db}

	if err := store.InitSchema(); err != nil {
		t.Fatal("InitSchema failed:", err)
	}

	// Clean users before each test
	_, err = db.Exec(`DELETE FROM users`)
	if err != nil {
		t.Fatalf("Failed to clear users table: %v", err)
	}

	return &Server{store: store}
}
func TestSignupHandler(t *testing.T) {
	server := setupTestServer(t)
	//json request
	body := `{"username":"testuser","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	server.SignupHandler(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "User created") {
		t.Errorf("Expected body to contain 'User created', got %q", rr.Body.String())
	}
	t.Log("Response body:", rr.Body.String())
}

func TestLoginHandler(t *testing.T) {
	server := setupTestServer(t)
	// Create test user
	server.store.CreateUser("testuser", "secret")

	body := `{"username":"testuser","password":"secret"}`

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	server.LoginHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "token") {
		t.Errorf("Expected response to contain token, got %s", rr.Body.String())
	}
	t.Log("Response body:", rr.Body.String())
}
func TestProtectedHandler(t *testing.T) {
	server := setupTestServer(t)
	// Create a token
	tokenString, _ := CreateToken("testuser")

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rr := httptest.NewRecorder()

	server.ProtectedHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "protected") {
		t.Errorf("Expected response body to contain 'protected', got %q", rr.Body.String())
	}
	t.Log("Response body:", rr.Body.String())
}
func TestProtectedHandler_NoToken(t *testing.T) {
	server := setupTestServer(t)
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer foobar")

	rr := httptest.NewRecorder()

	server.ProtectedHandler(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401 Unauthorized for bad token, got %d", rr.Code)
	}
}
func TestProtectedHandler_BadToken(t *testing.T) {
	server := setupTestServer(t)

	tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "fake",
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	wrongKey := []byte("wrongkey")
	tokenString, err := tokenObject.SignedString(wrongKey)
	if err != nil {
		t.Fatal("error signing fake token:", err)
	}
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rr := httptest.NewRecorder()
	server.ProtectedHandler(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401 Unauthorized for bad token, got %d", rr.Code)
	}
	t.Log("Response body:", rr.Body.String())

}
