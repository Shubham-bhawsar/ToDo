package views

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetAllTask(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/task", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllTask)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestCreateTask(t *testing.T) {

	var mockedTask = []byte(`{"Id":1,"Task":"Running","Status":true}`)
	req, err := http.NewRequest("POST", "/api/task", bytes.NewBuffer(mockedTask))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTask)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	// //fmt.Println(rr.Body.String())
	// expected := `{"Status":"Creates Successfully","StatusCode":201,"Proto":"","ProtoMajor":0,"ProtoMinor":0,"Header":null,"Body":null,"ContentLength":0,"TransferEncoding":null,"Close":false,"Uncompressed":false,"Trailer":null,"Request":null,"TLS":null}`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body, expected)
	// }

}

func TestGetTaskById(t *testing.T) {
	Data = Data[:0]
	var mockedTask = []byte(`{"Id":2,"Task":"Swimming","Status":true}`)
	req, err := http.NewRequest("POST", "/api/task", bytes.NewBuffer(mockedTask))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTask)

	handler.ServeHTTP(rr, req)

	req, err = http.NewRequest("GET", `/api/task/2`, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/task/{id}", GetTaskById)

	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	fmt.Println(rr.Body.String())
	output := rr.Body.String()
	expected := "{\"Id\":2,\"Task\":\"Swimming\",\"Status\":true}"
	fmt.Println(strings.Compare(output, expected))
	if strings.Compare(rr.Body.String(), expected) != 0 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body, expected)
	}

}
