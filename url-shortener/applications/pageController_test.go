package applications

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginPage(t *testing.T) {
	router := LoadRoutes()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	if err != nil {
		t.Errorf("cannot reach login page: %d", w.Code)
	}

	if w.Code == http.StatusOK {
		t.Log("status OK")
	}
}

func TestSignupPage(t *testing.T) {
	router := LoadRoutes()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/signupUser", nil)
	router.ServeHTTP(w, req)

	if err != nil {
		t.Errorf("cannot reach login page: %d", w.Code)
	}

	if w.Code == http.StatusOK {
		t.Log("status OK")
	}
}
