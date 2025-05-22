package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerStatus(t *testing.T) {
	srv := httptest.NewServer(NewServer())
	defer srv.Close()
	resp, err := http.Get(srv.URL + "/status")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status")
	}
}
