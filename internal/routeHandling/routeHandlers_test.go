package routeHandling

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"mathtestr.com/server/internal/dbHandling"
	"mathtestr.com/server/internal/schemas"
)

func TestRouteHandlers(t *testing.T) {
	if os.Getenv("MODE") != "PROD" {
		godotenv.Load("../../config.env")
	}

	dbHandler := dbHandling.InitDBHandler(os.Getenv("DB_URL_TEST"))
	routeHandler := InitRouteHandler(dbHandler)

	t.Run("Register", func(t *testing.T) {
		registerPayload := schemas.RegisterPayload{
			Username:  "a",
			Password:  "password",
			FirstName: "Jackson",
			LastName:  "Furr",
		}
		marshalled, _ := json.Marshal(registerPayload)

		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		r, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(marshalled))
		router.POST("/register", routeHandler.Register)
		router.ServeHTTP(w, r)

		fmt.Println(w.Body.String())
	})
}
