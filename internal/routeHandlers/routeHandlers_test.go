package routeHandlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouteHandlers(t *testing.T) {
	t.Run("Register", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/register", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(Register())
	})
}
