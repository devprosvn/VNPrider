// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package api

import (
	"encoding/json"
	"net/http"
)

// NewServer creates JSON-RPC server
// NewServer creates a simple HTTP server with JSON handlers.
func NewServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"status": "ok"})
	})
	return mux
}
