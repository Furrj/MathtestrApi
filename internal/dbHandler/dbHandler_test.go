package dbHandler

import (
	"context"
	"mathtestr.com/server/internal/testHelpers/objects"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestDBHandler(t *testing.T) {
	if os.Getenv("MODE") != "PROD" {
		godotenv.Load("../../config.env")
	}

	// Vars
	testRegisterPayloadStudent := objects.TestRegisterPayloadStudent
	testSessionDataStudent := objects.TestSessionDataStudent
	testResultsDataStudent := objects.TestResultsDataStudent
	testAllUserDataStudent := objects.TestAllUserDataStudent

	testRegisterPayloadTeacher := objects.TestRegisterPayloadTeacher
	testTeacherInfo := objects.TestTeacherInfo

	dbHandler := InitDBHandler(os.Getenv("DB_URL_TEST"))
	defer dbHandler.DB.Close(context.Background())

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
	t.Run("InsertUserInfoStudent", func(t *testing.T) {
		if err := dbHandler.InsertUserInfo("S", testRegisterPayloadStudent); err != nil {
			t.Errorf("Error inserting user: %+v\n", err)
		}

		exists, err := dbHandler.CheckIfUsernameExists(testRegisterPayloadStudent.Username)
		if err != nil {
			t.Errorf("Error checking to see if user was inserted: %+v\n", err)
		}
		if !exists {
			t.Errorf("User could not be found after inserting")
		}
	})
	t.Run("InsertUserInfoTeacher", func(t *testing.T) {
		if err := dbHandler.InsertUserInfo("S", testRegisterPayloadTeacher); err != nil {
			t.Errorf("Error inserting user: %+v\n", err)
		}

		exists, err := dbHandler.CheckIfUsernameExists(testRegisterPayloadTeacher.Username)
		if err != nil {
			t.Errorf("Error checking to see if user was inserted: %+v\n", err)
		}
		if !exists {
			t.Errorf("User could not be found after inserting")
		}
	})
	t.Run("InsertTeacherInfo", func(t *testing.T) {
		if err := dbHandler.InsertTeacherInfo(testTeacherInfo); err != nil {
			t.Errorf("Error inserting teacher info: %+v\n", err)
		}
	})
	t.Run("InsertStudentInfo", func(t *testing.T) {
		if err := dbHandler.InsertStudentInfo(1, testRegisterPayloadStudent); err != nil {
			t.Errorf("Error inserting student info: %+v\n", err)
		}
	})
	t.Run("InsertSessionData", func(t *testing.T) {
		if err := dbHandler.InsertSessionData(testSessionDataStudent); err != nil {
			t.Errorf("Error inserting session data: %+v\n", err)
		}
	})
	t.Run("InsertTestResults", func(t *testing.T) {
		if err := dbHandler.InsertTestResults(testResultsDataStudent); err != nil {
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
		if got != testAllUserDataStudent {
			t.Errorf("got %+v\n, want %+v\n for AllUserData", got, testAllUserDataStudent)
		}
		if got.SessionKey != wantSessionID {
			t.Errorf("got %s, want %s for session_key", got.SessionKey, wantSessionID)
		}
	})
	t.Run("GetSessionData", func(t *testing.T) {
		got, err := dbHandler.GetSessionDataByUserID(1)
		if err != nil {
			t.Errorf("Error searching for test result by ID: %+v\n", err)
		}
		if got != testSessionDataStudent {
			t.Errorf("got %+v\n, want %+v\n for SessionData", got, testSessionDataStudent)
		}
	})
	t.Run("GetTestResults", func(t *testing.T) {
		got, err := dbHandler.GetTestResultsByUserID(1)
		if err != nil {
			t.Errorf("Error searching for test result by ID: %+v\n", err)
		}
		if got != testResultsDataStudent {
			t.Errorf("got %+v\n, want %+v\n for TestResults", got, testResultsDataStudent)
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
