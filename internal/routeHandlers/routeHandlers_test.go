package routeHandlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/types"
)

func TestRouteHandlers(t *testing.T) {
	t.Run("Register", func(t *testing.T) {
		registerPayload := types.RegisterPayload{
			Username:  "Poemmys",
			Password:  "password",
			FirstName: "Jackson",
			LastName:  "Furr",
		}
		marshalled, _ := json.Marshal(registerPayload)

		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(marshalled))

		_, router := gin.CreateTestContext(w)
		router.POST("/register", Register)
		router.ServeHTTP(w, r)

		fmt.Println(w.Body.String())
	})
}
