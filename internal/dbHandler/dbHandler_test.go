package dbHandler

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"mathtestr.com/server/internal/schemas"
)

func TestDBHandler(t *testing.T) {
	if os.Getenv("MODE") != "PROD" {
		godotenv.Load("../../config.env")
	}

	dbHandler := InitDBHandler(os.Getenv("DB_URL_TEST"))
	defer dbHandler.DB.Close(context.Background())

	// SETUP
	testRegisterPayload := schemas.RegisterPayload{
		Username:  "a",
		Password:  "a",
		FirstName: "Jackson",
		LastName:  "Furr",
		Period:    0,
		Teacher:   "Mrs. Furr",
	}

	testSessionData := schemas.SessionData{
		ID:         1,
		SessionKey: "test_uuid",
		Expires:    1234,
	}

	testResultsData := schemas.TestResults{
		ID:            1,
		Score:         100,
		Min:           0,
		Max:           12,
		QuestionCount: 10,
		Operations:    "multiplication,addition",
	}

	testAllUserData := schemas.AllUserData{
		Username:   "a",
		Password:   "a",
		FirstName:  "Jackson",
		LastName:   "Furr",
		Period:     0,
		Teacher:    "Mrs. Furr",
		Role:       "S",
		ID:         1,
		SessionKey: "test_uuid",
		Expires:    1234,
	}

	t.Run("Ping connection", func(t *testing.T) {
		if err := dbHandler.DB.Ping(context.Background()); err != nil {
			t.Errorf("Error initializing database: %+v\n", err)
		}
	})
	t.Run("InitTables", func(t *testing.T) {
		if err := dbHandler.CreateTables(); err != nil {
			t.Errorf("Error initializing tables: %+v\n", err)
		}
	})
	t.Run("InsertUserInfo", func(t *testing.T) {
		if err := dbHandler.InsertUserInfo(testRegisterPayload); err != nil {
			t.Errorf("Error inserting user: %+v\n", err)
		}

		exists, err := dbHandler.CheckIfUsernameExists(testRegisterPayload.Username)
		if err != nil {
			t.Errorf("Error checking to see if user was inserted: %+v\n", err)
		}
		if !exists {
			t.Errorf("User could not be found after inserting")
		}
	})
	t.Run("InsertSessionData", func(t *testing.T) {
		if err := dbHandler.InsertSessionData(testSessionData); err != nil {
			t.Errorf("Error inserting session data: %+v\n", err)
		}
	})
	t.Run("InsertTestResults", func(t *testing.T) {
		if err := dbHandler.InsertTestResults(testResultsData); err != nil {
			t.Errorf("Error inserting test results: %+v\n", err)
		}
	})
	t.Run("Existing CheckIfUsernameExists", func(t *testing.T) {
		got, err := dbHandler.CheckIfUsernameExists("a")

		if err != nil {
			t.Errorf("Error querying for username: %+v\n", err)
		}
		if !got {
			t.Errorf("Username exists but returned false")
		}
	})
	t.Run("Non-existent CheckIfUsernameExists", func(t *testing.T) {
		got, err := dbHandler.CheckIfUsernameExists("")

		if err != nil {
			t.Errorf("Error querying for username: %+v\n", err)
		}
		if got {
			t.Errorf("Username doesn't exist but returned true")
		}
	})
	t.Run("GetUserByUsername", func(t *testing.T) {
		got, err := dbHandler.GetUserByUsername("a")
		const wantSessionID = "test_uuid"
		if err != nil {
			t.Errorf("Error when querying for user: %+v\n", err)
		}
		if got != testAllUserData {
			t.Errorf("got %+v\n, want %+v\n for AllUserData", got, testAllUserData)
		}
		if got.SessionKey != wantSessionID {
			t.Errorf("got %s, want %s for session_key", got.SessionKey, wantSessionID)
		}
	})
	t.Run("GetTestResults", func(t *testing.T) {
		got, err := dbHandler.GetTestResultsByUserID(1)
		if err != nil {
			t.Errorf("Error searching for test result by ID: %+v\n", err)
		}
		if got != testResultsData {
			t.Errorf("got %+v\n, want %+v\n for TestResults", got, testResultsData)
		}
	})
	t.Run("GetUserIDByUsername", func(t *testing.T) {
		got, err := dbHandler.GetUserIDByUsername("a")
		want := 1
		if err != nil {
			t.Errorf("Error searching for ID by username: %+v\n", err)
		}
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("DropTables", func(t *testing.T) {
		if err := dbHandler.DropTables(); err != nil {
			t.Errorf("Error dropping tables: %+v\n", err)
		}
	})
}
