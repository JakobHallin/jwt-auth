package main
import (
	//"sync"
	//"errors"
	"net/http"
	"encoding/json"
	"io"
	"strings"
)
type Server struct {
    store *UserStore
}
//https://go.dev/play/p/zxbo1VH0Foy
func (s *Server) SignupHandler(w http.ResponseWriter, r *http.Request){ //need to be post
	if r.Method != http.MethodPost{
		http.Error(w, "method not Post it is", http.StatusMethodNotAllowed)
		return
	}
	var creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := s.store.CreateUser(creds.Username, creds.Password); err != nil {
		http.Error(w, "User creation failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "User created")
}
func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if !s.store.Auth(creds.Username, creds.Password) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tokenString, err := CreateToken(creds.Username)
	if err != nil {
		http.Error(w, "Token creation failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
func (s *Server) ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := ValidateToken(tokenString)
	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "This is protected data")
}
