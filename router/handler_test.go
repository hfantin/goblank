package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hfantin/goblank/router"
)

func TestHealthHandler(t *testing.T) {
	t.Run("returns health information", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/actuator/info", nil)
		response := httptest.NewRecorder()
		router.HealthHandler(response, request)
		got := response.Body.String()
		want := `{"health":"OK"}`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
