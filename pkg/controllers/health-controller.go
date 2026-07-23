package controllers

import (
	"encoding/json"
	"net/http"
)

// Health handles the /health endpoint. It reports service liveness without
// touching the database.
func Health(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(map[string]string{"status": "ok"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
