package controllers

import (
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	server := Server{}
	t.Run("Login", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "https://example.com/foo", nil)
		server.Login(w, r)
	})
}
