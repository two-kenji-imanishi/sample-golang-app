package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("unexpected status: got %v want %v", status, http.StatusOK)
	}

	expected := "Welcome to World!\n"
	if rr.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("unexpected status: got %v want %v", status, http.StatusOK)
	}

	expected := fmt.Sprintf("Hello! My name is %s.", os.Getenv("NAME"))
	if rr.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
