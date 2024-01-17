package applications

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestLoadroutes checks whether public pages are up and running
func TestLoadroutes(t *testing.T) {
	router := LoadRoutes()

	testRoutes := []string{"/"}

	for _, route := range testRoutes {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", route, nil)
		router.ServeHTTP(w, req)

		if err != nil {
			t.Errorf("error: %s", err)
		}

		if w.Code == http.StatusOK {
			t.Log("status OK")
		} else {
			t.Errorf("error code: %d", w.Code)
		}
	}
}

//func TestLogin(t *testing.T) {
//	router := LoadRoutes()
//
//	w := httptest.NewRecorder()
//	testData := "username=arnavrneo&email=arnavrainays@gmail.com&password=123"
//	req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(testData))
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	router.ServeHTTP(w, req)
//
//	if err != nil {
//		t.Errorf("error %s", err)
//	}
//
//	if w.Code == http.StatusFound {
//		t.Log("status OK")
//	} else {
//		t.Errorf("error code: %d", w.Code)
//	}
//}
