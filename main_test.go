package main_test
 
import (
     "os"
     "testing"
	 "net/http"
     "net/http/httptest"
     "github.com/succoDiPompelmo/testAPI"
 )

 var a main.App

func TestMain(m *testing.M) {
	a := main.App{}
	a.Init()
    code := m.Run()
    os.Exit(code)
}

func TestSomething(t *testing.T) {
	req, _ := http.NewRequest("GET", "/product?id=1", nil)
    response := executeRequest(req)

    checkResponseCode(t, http.StatusOK, response.Code)

    if body := response.Body.String(); body != "{}" {
        t.Errorf("Expected an empty array. Got %s", body)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(main.GetProduct)
    handler.ServeHTTP(rr, req)
    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}