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
	testSessionDataTeacher := objects.TestSessionDataTeacher
	testResultsDataStudent := objects.TestResultsDataStudent
	testBasicUserData := objects.TestBasicUserData
	testAllUserDataStudent := objects.TestAllUserDataStudent
	testAllUserDataTeacher := objects.TestAllUserDataTeacher

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
	t.Run("InsertUserInfoTeacher", func(t *testing.T) {
		if err := dbHandler.InsertUserInfo("T", testRegisterPayloadTeacher); err != nil {
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
	t.Run("InsertSessionDataTeacher", func(t *testing.T) {
		if err := dbHandler.InsertSessionData(testSessionDataTeacher); err != nil {
			t.Errorf("Error inserting session data for teacher: %+v\n", err)
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
	t.Run("InsertStudentInfo", func(t *testing.T) {
		if err := dbHandler.InsertStudentInfo(testAllUserDataStudent.ID, testRegisterPayloadStudent); err != nil {
			t.Errorf("Error inserting student info: %+v\n", err)
		}
	})
	t.Run("InsertSessionDataStudent", func(t *testing.T) {
		if err := dbHandler.InsertSessionData(testSessionDataStudent); err != nil {
			t.Errorf("Error inserting session data for student: %+v\n", err)
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
	t.Run("GetBasicUserInfoByID", func(t *testing.T) {
		got, err := dbHandler.GetBasicUserInfoByID(testSessionDataStudent.ID)
		if err != nil {
			t.Errorf("Error when querying for user: %+v\n", err)
		}
		if got != testBasicUserData {
			t.Errorf("got %+v\n, want %+v\n for BasicUserData", got, testBasicUserData)
		}
	})
	t.Run("GetBasicUserInfoByUsername", func(t *testing.T) {
		got, err := dbHandler.GetBasicUserInfoByUsername("a")
		if err != nil {
			t.Errorf("Error when querying for user: %+v\n", err)
		}
		if got != testBasicUserData {
			t.Errorf("got %+v\n, want %+v\n for BasicUserData", got, testBasicUserData)
		}
	})
	t.Run("GetAllUserDataTeacher", func(t *testing.T) {
		got, err := dbHandler.GetAllTeacherDataByUsername(testAllUserDataTeacher.Username)
		if err != nil {
			t.Errorf("Error when querying for user: %+v\n", err)
		}
		if got != testAllUserDataTeacher {
			t.Errorf("got %+v\n, want %+v\n for AllUserDataTeacher", got, testAllUserDataTeacher)
		}
	})
	t.Run("GetAllUserDataStudent", func(t *testing.T) {
		got, err := dbHandler.GetAllStudentDataByUsername(testAllUserDataStudent.Username)
		if err != nil {
			t.Errorf("Error when querying for user: %+v\n", err)
		}
		if got != testAllUserDataStudent {
			t.Errorf("got %+v\n, want %+v\n for AllUserDataStudent", got, testAllUserDataStudent)
		}
	})
	t.Run("GetSessionDataStudent", func(t *testing.T) {
		got, err := dbHandler.GetSessionDataByUserID(int(testAllUserDataStudent.ID))
		if err != nil {
			t.Errorf("Error searching for test result by ID: %+v\n", err)
		}
		if got != testSessionDataStudent {
			t.Errorf("got %+v\n, want %+v\n for SessionData", got, testSessionDataStudent)
		}
	})
	t.Run("GetTestResults", func(t *testing.T) {
		got, err := dbHandler.GetTestResultsByUserID(int(testAllUserDataStudent.ID))
		if err != nil {
			t.Errorf("Error searching for test result by ID: %+v\n", err)
		}
		if got != testResultsDataStudent {
			t.Errorf("got %+v\n, want %+v\n for TestResults", got, testResultsDataStudent)
		}
	})
	t.Run("GetUserIDByUsername", func(t *testing.T) {
		got, err := dbHandler.GetUserIDByUsername("a")
		want := int(testAllUserDataStudent.ID)
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
