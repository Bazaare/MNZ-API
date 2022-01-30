package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestTokenCreateNoError(t *testing.T) {
	t.Run("Token Create no errors", func(t *testing.T) {
		_, err := CreateToken()
		if err != nil {
			t.Errorf("expected no error but got: %s", err)
		}
	})

	t.Run("Token Create check expiry", func(t *testing.T) {
		tokenString, _ := CreateToken()
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("API_SECRET")), nil
		})
		claims, _ := token.Claims.(jwt.MapClaims)
		exp, _ := strconv.ParseUint(fmt.Sprintf("%.0f", claims["exp"]), 10, 32)
		if int64(exp) < time.Now().Unix() {
			t.Errorf("generated jwt has already expired")
		}
	})
}

func TestTokenValid(t *testing.T) {
	t.Run("Check Token is valid", func(t *testing.T) {
		tokenString, _ := CreateToken()
		r := httptest.NewRequest("POST", "https://example.com/foo", nil)
		bearer := "Bearer " + tokenString
		r.Header.Add("Authorization", bearer)
		err := TokenValid(r)
		if err != nil {
			t.Errorf("expected no error but got: %s", err)
		}
	})

	t.Run("Check error when invalid", func(t *testing.T) {
		r := httptest.NewRequest("POST", "https://example.com/foo", nil)
		r.Header.Add("Authorization: Bearer ", "invalid token")
		err := TokenValid(r)
		if err == nil {
			t.Error("expected error but did not get one")
		}
	})
}
