package router_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"

	"github.com/hfantin/goblank/router"
)

func init() {
	fmt.Println("init test")
}

func TestHealthHandler(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	t.Logf("Current test filename: %s", filename)
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
