package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDashboard(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.URL.Path {
		case "/users/1":
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{
				"id": 1,
				"firstName": "Pablo",
				"lastName": "Bustos",
				"age": 33
			}`))

		case "/todos/user/1":
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{
				"todos": [
					{"id":1,"todo":"task 1","completed":false},
					{"id":2,"todo":"task 2","completed":true},
					{"id":3,"todo":"task 3","completed":false}
				]
			}`))

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))

	defer mockServer.Close()

	client := mockServer.Client()
	svc := NewDashboardService(client, mockServer.URL)

	result, err := svc.GetDashboard("1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.FullName != "Pablo Bustos" {
		t.Errorf("expected result full name 'Pablo Bustos', got %s", result.FullName)
	}

	if result.Status != "Rookie" {
		t.Errorf("expected result status Rookie, got %s", result.Status)
	}

	if result.PendingTaskCount != 2 {
		t.Errorf("expected result pending task count 2, got %d", result.PendingTaskCount)
	}

	if result.NextUrgentTask != "task 1" {
		t.Errorf("expected result next urgent task 'task 1', got %s", result.NextUrgentTask)
	}
}
