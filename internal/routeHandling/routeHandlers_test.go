package routeHandling

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"mathtestr.com/server/internal/dbHandler"
	"mathtestr.com/server/internal/schemas"
)

func TestRouteHandlers(t *testing.T) {
	if os.Getenv("MODE") != "PROD" {
		godotenv.Load("../../config.env")
	}

	db := dbHandler.InitDBHandler(os.Getenv("DB_URL_TEST"))
	routeHandler := InitRouteHandler(db)
	defer routeHandler.dbHandler.DB.Close(context.Background())

	var responseData schemas.RegisterResponse
	var responseUserClientData schemas.UserClientData

	t.Run("Register", func(t *testing.T) {
		if err := db.CreateTables(); err != nil {
			t.Errorf("Error creating tables: %+v\n", err)
		}

		registerPayload := schemas.RegisterPayload{
			Username:  "a",
			Password:  "password",
			FirstName: "Jackson",
			LastName:  "Furr",
			Period:    0,
		}

		marshalled, _ := json.Marshal(registerPayload)

		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		r, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(marshalled))
		router.POST("/api/register", routeHandler.Register)
		router.ServeHTTP(w, r)

		json.Unmarshal(w.Body.Bytes(), &responseData)
		responseUserClientData = responseData.User

		if !checkIfUserInserted(t, routeHandler, registerPayload) {
			t.Errorf("User could not be found, problem inserting")
		}

	})
	t.Run("ValidateSession_valid", func(t *testing.T) {
		var validationReponse schemas.SessionValidationResponse
		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		marshalled, _ := json.Marshal(responseUserClientData)

		r, _ := http.NewRequest(http.MethodPost, "/api/validateSession", bytes.NewReader(marshalled))
		router.POST("/api/validateSession", routeHandler.ValidateSession)
		router.ServeHTTP(w, r)

		json.Unmarshal(w.Body.Bytes(), &validationReponse)
		if !validationReponse.Valid {
			t.Error("Wanted valid response, got invalid")
		}
	})
	t.Run("ValidateSession_invalid", func(t *testing.T) {
		var validationReponse schemas.SessionValidationResponse
		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		responseUserClientData.SessionKey = "invalid"

		marshalled, _ := json.Marshal(responseUserClientData)

		r, _ := http.NewRequest(http.MethodPost, "/api/validateSession", bytes.NewReader(marshalled))
		router.POST("/api/validateSession", routeHandler.ValidateSession)
		router.ServeHTTP(w, r)

		json.Unmarshal(w.Body.Bytes(), &validationReponse)
		if validationReponse.Valid {
			t.Error("Wanted invalid response, got valid")
		}
	})
	t.Run("Login_valid", func(t *testing.T) {
		var loginResponse schemas.LoginResponse
		loginPayload := schemas.LoginPayload{
			Username: "a",
			Password: "password",
		}
		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		marshalled, _ := json.Marshal(loginPayload)

		r, _ := http.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(marshalled))
		router.POST("/api/login", routeHandler.Login)
		router.ServeHTTP(w, r)

		json.Unmarshal(w.Body.Bytes(), &loginResponse)
		if !loginResponse.Valid {
			t.Error("Wanted valid response, got invalid")
		}
	})
	t.Run("Login_invalid_username", func(t *testing.T) {
		var loginResponse schemas.LoginResponse
		loginPayload := schemas.LoginPayload{
			Username: "b",
			Password: "password",
		}
		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		marshalled, _ := json.Marshal(loginPayload)

		r, _ := http.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(marshalled))
		router.POST("/api/login", routeHandler.Login)
		router.ServeHTTP(w, r)

		json.Unmarshal(w.Body.Bytes(), &loginResponse)
		if loginResponse.Valid {
			t.Error("Wanted invalid response, got valid")
		}
	})
	t.Run("Login_invalid_password", func(t *testing.T) {
		var loginResponse schemas.LoginResponse
		loginPayload := schemas.LoginPayload{
			Username: "a",
			Password: "passwor",
		}
		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		marshalled, _ := json.Marshal(loginPayload)

		r, _ := http.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(marshalled))
		router.POST("/api/login", routeHandler.Login)
		router.ServeHTTP(w, r)

		json.Unmarshal(w.Body.Bytes(), &loginResponse)
		if loginResponse.Valid {
			t.Error("Wanted invalid response, got valid")
		}
	})

	if err := db.DropTables(); err != nil {
		t.Errorf("Error dropping tables: %+v\n", err)
	}
}

func checkIfUserInserted(t *testing.T, r *RouteHandler, p schemas.RegisterPayload) bool {
	t.Helper()
	result, _ := r.dbHandler.CheckIfUsernameExists(p.Username)
	return result
}
