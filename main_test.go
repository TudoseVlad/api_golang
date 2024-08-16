package main

import (
	"encoding/json"
	"golang_api/src/api"
	"golang_api/src/data"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestPOSTHandler(t *testing.T) {
	// Create a request body
	body := "Ana are mere"
	data.InitData()
	req, err := http.NewRequest("POST", "/cuvinte", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("root", "1234")

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Wrap the handler with the auth middleware
	handler := api.AuthMiddleware(http.HandlerFunc(api.POSTHandler))

	// Perform the request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	expectedResp := map[string]string{"message": "Words saved successfully"}
	var actualResp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&actualResp); err != nil {
		t.Fatalf("Failed to decode actual response: %v", err)
	}
	if !reflect.DeepEqual(actualResp, expectedResp) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedResp)
	}

}
func TestPOSTIncerment(t *testing.T) {
	// Create a request body
	body := "Ana are mere"
	data.InitData()
	v1 := data.GetValue("Ana") + 1
	v2 := data.GetValue("are") + 1
	v3 := data.GetValue("mere") + 1
	req, err := http.NewRequest("POST", "/cuvinte", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("root", "1234")
	data.InitData()
	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Wrap the handler with the auth middleware
	handler := api.AuthMiddleware(http.HandlerFunc(api.POSTHandler))

	// Perform the request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body
	if v1 != data.GetValue("Ana") || v2 != data.GetValue("are") || v3 != data.GetValue("mere") {
		t.Errorf("Data has not been changed")
	}
}

func TestGETHandler(t *testing.T) {
	// Populate the map with some data for testing
	data.InitData()
	// Create a GET request
	req, err := http.NewRequest("GET", "/cuvinte?words=Ana_are_mere", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to capture the response
	req.SetBasicAuth("root", "1234")
	rr := httptest.NewRecorder()

	// Wrap the handler with the auth middleware
	handler := api.AuthMiddleware(http.HandlerFunc(api.GETHandler))

	// Perform the request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expectedResp := make(map[string]int)
	expectedResp["Ana"] = data.GetValue("Ana")
	expectedResp["are"] = data.GetValue("are")
	expectedResp["mere"] = data.GetValue("mere")
	var actualResp map[string]int
	if err := json.NewDecoder(rr.Body).Decode(&actualResp); err != nil {
		t.Fatalf("Failed to decode actual response: %v", err)
	}
	if !reflect.DeepEqual(actualResp, expectedResp) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedResp)
	}
}

// TestAuthMiddleware tests the AuthMiddleware for a valid request
func TestAuthMiddleware(t *testing.T) {
	req, err := http.NewRequest("GET", "/cuvinte", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetBasicAuth("root", "1234")

	rr := httptest.NewRecorder()
	handler := api.AuthMiddleware(http.HandlerFunc(api.GETHandler))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestAuthMiddleware_Unauthorized tests the AuthMiddleware for an unauthorized request
func TestAuthMiddleware_Unauthorized(t *testing.T) {
	req, err := http.NewRequest("GET", "/cuvinte", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := api.AuthMiddleware(http.HandlerFunc(api.GETHandler))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
}
