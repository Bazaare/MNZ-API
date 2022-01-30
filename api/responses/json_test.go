package responses

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestERROR(t *testing.T) {
	t.Run("TestError", func(t *testing.T) {
		w := httptest.NewRecorder()
		ERROR(w, http.StatusBadGateway, errors.New("error"))
	})
}
