package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	Health(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected content-type %q, got %q", "application/json", contentType)
	}

	expectedBody := `{"status":"ok"}`
	body := rec.Body.String()
	if body != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, body)
	}
}
