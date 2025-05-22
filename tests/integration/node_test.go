// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/devprosvn/VNPrider/pkg/api"
)

func TestNodeStartup(t *testing.T) {
	srv := httptest.NewServer(api.NewServer())
	defer srv.Close()
	resp, err := http.Get(srv.URL + "/status")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status %d", resp.StatusCode)
	}
}
