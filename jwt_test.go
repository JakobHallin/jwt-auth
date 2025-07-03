package main
import ( "testing"
	"time"
	"github.com/golang-jwt/jwt/v5"
	//"fmt"
)
func TestCreateToken(t *testing.T){
	token, err := CreateToken("testsuser")
	if err != nil {
		t.Fatal(err)
	}
	if token == "" {
		t.Fatal("expected that token is not empty")
	}
	t.Log("Generated token:", token)
}
func TestValidate(t *testing.T){
	// case 1 call validate on true token
	//creat token to cheack
	tokenstring, err := CreateToken("testsuser")
        if err != nil {
                t.Fatal("error with creatin token", err)
        }
	token, err :=  ValidateToken(tokenstring)
	if err != nil {
		t.Fatal("ValidateToken returned error:", err)
	}
	if !token.Valid {
		t.Fatal("expected token to be valid, but it is invalid")
	}
	t.Log("Validation succeeded for real token")
	//case 2 call validate on fake token
	tokenobject := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "fake",
		 "exp":      time.Now().Add(time.Hour * 24).Unix(), })
	wrongKey := []byte("wrongkey")
	tokenstring, err = tokenobject.SignedString(wrongKey)
	if err != nil {
		t.Fatal("error signing fake token:", err)
	}
	token, err =  ValidateToken(tokenstring)
	if err != nil {
		t.Log("expected error for fake token:", err)
	}
	if token.Valid {
		t.Fatal("expected token to be invalid, but it is valid")
	}
	t.Log("Validation correctly failed on fake token:", err)
}
