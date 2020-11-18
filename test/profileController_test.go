package test

import (
	"burndown-ticket/app"
	"burndown-ticket/controllers"
	"fmt"

	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

var a app.App

func TestMain(m *testing.M) {
	a = app.App{}
	a.Initialize("root", "", "rest_api_example")

	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, req)
	return recorder
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/profile/1", controllers.GetProfile).Methods("GET")
	return router
}

func TestGetProfile404(t *testing.T) {
	request, _ := http.NewRequest("GET", "/profile/1234", nil)
	response := executeRequest(request)
	checkResponseCode(t, http.StatusNotFound, response.Code)
	fmt.Println(response.Body)
	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}
func TestGetProfileFail(t *testing.T) {

	t.Fatalf("this test failed because i said so")
	t.Fail()
}

// func TestGetProfile(t *testing.T) {
// 	request, _ := http.NewRequest("GET", "/profile/1", nil)
// 	response := httptest.NewRecorder()
// 	Router().ServeHTTP(response, request)
// 	assert.Equal(t, 200, response.Code, "OK response is expected")
// }
