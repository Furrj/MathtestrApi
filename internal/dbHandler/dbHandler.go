// Package dbHandler handles communication with the Postgres database
package dbHandler

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"mathtestr.com/server/internal/testHelpers/objects"
	"os"
	"strconv"

	"mathtestr.com/server/internal/schemas"
)

// DBHandler object contains a pointer to pgx connection, one per server
type DBHandler struct {
	DB *pgxpool.Pool
}

// InitDBHandler is constructor for DBHandler, takes a database connection
// string and returns a pointer to instantiated DBHandler.
func InitDBHandler(connectionString string) *DBHandler {
	var newDBHandler DBHandler
	db, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	newDBHandler.DB = db
	return &newDBHandler
}

// CheckIfUsernameExists takes in username string and searches database
// for it. Returns true/false and error.
func (dbHandler *DBHandler) CheckIfUsernameExists(username string) (bool, error) {
	var returnedUsername string
	err := dbHandler.DB.QueryRow(context.Background(), QCheckIfUsernameExists, username).Scan(&returnedUsername)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetUserIDByUsername takes in username string and searches database
// for it. Returns userID (-1 if error) and error
func (dbHandler *DBHandler) GetUserIDByUsername(username string) (int, error) {
	var id int
	err := dbHandler.DB.QueryRow(context.Background(), QGetUserIDByUsername, username).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetBasicUserInfoByID takes in user_id uint32 and searches database,
// binds query to BasicUserData schema, then returns result and error
func (dbHandler *DBHandler) GetBasicUserInfoByID(UserID uint32) (schemas.BasicUserData, error) {
	var user schemas.BasicUserData
	err := dbHandler.DB.QueryRow(context.Background(), QGetBasicUserDataByID, UserID).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Role)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetBasicUserInfoByUsername takes in username string and searches database,
// binds query to BasicUserData schema, then returns result and error
func (dbHandler *DBHandler) GetBasicUserInfoByUsername(username string) (schemas.BasicUserData, error) {
	var user schemas.BasicUserData
	err := dbHandler.DB.QueryRow(context.Background(), QGetBasicUserDataByUsername, username).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Role)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetAllStudentDataByUsername takes in username string and searches database,
// binds query to AllUserDataStudent schema, then returns result and error
func (dbHandler *DBHandler) GetAllStudentDataByUsername(username string) (schemas.AllUserDataStudent, error) {
	var user schemas.AllUserDataStudent
	if err := dbHandler.DB.QueryRow(context.Background(), QGetStudentByUsername, username).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Role, &user.Period, &user.TeacherID, &user.SessionKey, &user.Expires); err != nil {
		return user, err
	}
	return user, nil
}

// GetAllTeacherDataByUsername takes in username string and searches database,
// binds query to AllUserDataTeacher schema, then returns result and error
func (dbHandler *DBHandler) GetAllTeacherDataByUsername(username string) (schemas.AllUserDataTeacher, error) {
	var user schemas.AllUserDataTeacher
	if err := dbHandler.DB.QueryRow(context.Background(), QGetTeacherByUsername, username).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Role, &user.Periods, &user.SessionKey, &user.Expires); err != nil {
		return user, err
	}
	return user, nil
}

// GetSessionDataByUserID takes in username string, searches then returns
// session data and error
func (dbHandler *DBHandler) GetSessionDataByUserID(id int) (schemas.SessionData, error) {
	var sessionData schemas.SessionData

	if err := dbHandler.DB.QueryRow(context.Background(), QGetSessionDataByUserID, id).Scan(&sessionData.ID, &sessionData.SessionKey, &sessionData.Expires); err != nil {
		return sessionData, err
	} else if id != int(sessionData.ID) {
		return sessionData, errors.New("id mismatch when searching for session data by id")
	}

	return sessionData, nil
}

// GetTestResultsByUserID takes in username string and searches test_results
// for all rows under username and error
// FIXME: Iterate through multiple rows
func (dbHandler *DBHandler) GetTestResultsByUserID(id int) (schemas.TestResults, error) {
	var results schemas.TestResults

	if err := dbHandler.DB.QueryRow(context.Background(), QGetTestResultsByUserID, id).Scan(&results.ID, &results.Score, &results.Min, &results.Max, &results.QuestionCount, &results.Operations, &results.Timestamp); err != nil {
		return results, err
	}

	return results, nil
}

// InsertUserInfo takes Role and RegisterPayload and inserts user data into user_info
// table, returns error
func (dbHandler *DBHandler) InsertUserInfo(role string, r schemas.RegisterPayload) error {
	_, err := dbHandler.DB.Exec(context.Background(), EInsertUserInfo, r.Username, r.Password, r.FirstName, r.LastName, role)
	if err != nil {
		return err
	}
	return nil
}

func (dbHandler *DBHandler) InsertStudentInfo(userId uint32, s schemas.RegisterPayload) error {
	_, err := dbHandler.DB.Exec(context.Background(), EInsertStudentInfo, strconv.Itoa(int(userId)), s.TeacherID, s.Period)
	if err != nil {
		return err
	}
	return nil
}

func (dbHandler *DBHandler) InsertTeacherInfo(t schemas.TeacherData) error {
	if _, err := dbHandler.DB.Exec(context.Background(), EInsertTeacherInfo, t.ID, t.Periods); err != nil {
		return err
	}
	return nil
}

// InsertSessionData takes SessionData object and inserts it into database
// session_data table, returns error
func (dbHandler *DBHandler) InsertSessionData(s schemas.SessionData) error {
	_, err := dbHandler.DB.Exec(context.Background(), EInsertSessionData, s.ID, s.SessionKey, s.Expires)
	if err != nil {
		log.Printf("Error inserting session data: %+v\n", err)
		return err
	}
	return nil
}

// InsertTestResults takes SessionData object and inserts it into database
// session_data table, returns errors
func (dbHandler *DBHandler) InsertTestResults(t schemas.TestResults) error {
	_, err := dbHandler.DB.Exec(context.Background(), EInsertTestResults, t.ID, t.Score, t.Min, t.Max, t.QuestionCount, t.Operations, t.Timestamp)
	if err != nil {
		log.Printf("Error inserting test results: %+v\n", err)
		return err
	}
	return nil
}

// TESTING

func (dbHandler *DBHandler) TestInsertTeacher() error {
	if err := dbHandler.InsertUserInfo("T", objects.TestRegisterPayloadTeacher); err != nil {
		fmt.Printf("Error insert test teacher user_info: %+v\n", err)
		return err
	}
	if err := dbHandler.InsertTeacherInfo(objects.TestTeacherInfo); err != nil {
		fmt.Printf("Error inserting test teacher_info: %+v\n", err)
		return err
	}
	return nil
}

// CreateTables is used for testing, creates all tables from database, returns error
func (dbHandler *DBHandler) CreateTables() error {
	if _, err := dbHandler.DB.Exec(context.Background(), ECreateRole); err != nil {
		fmt.Printf("Error initializing user_info: %+v\n", err)
	}
	if _, err := dbHandler.DB.Exec(context.Background(), EInitUserInfo); err != nil {
		fmt.Printf("Error initializing user_info: %+v\n", err)
	}
	if _, err := dbHandler.DB.Exec(context.Background(), EInitSessionData); err != nil {
		fmt.Printf("Error initializing session_data: %+v\n", err)
	}
	if _, err := dbHandler.DB.Exec(context.Background(), EInitTeacherInfo); err != nil {
		fmt.Printf("Error initializing teacher_info: %+v\n", err)
	}
	if _, err := dbHandler.DB.Exec(context.Background(), EInitStudentInfo); err != nil {
		fmt.Printf("Error initializing student_info: %+v\n", err)
	}
	if _, err := dbHandler.DB.Exec(context.Background(), EInitTestResults); err != nil {
		fmt.Printf("Error initializing session_data: %+v\n", err)
	}
	return nil
}

// DropTables is used for testing, drops all tables from database, returns error
func (dbHandler *DBHandler) DropTables() error {
	if _, err := dbHandler.DB.Exec(context.Background(), EDeleteAllStudentInfo); err != nil {
		fmt.Printf("Error dropping student_info table: %+v\n", err)
		return err
	}
	if _, err := dbHandler.DB.Exec(context.Background(), EDeleteAllTeacherInfo); err != nil {
		fmt.Printf("Error dropping teacher_info table: %+v\n", err)
		return err
	}
	if _, err := dbHandler.DB.Exec(context.Background(), EDeleteAllSessionData); err != nil {
		fmt.Printf("Error dropping session_data table: %+v\n", err)
		return err
	}
	if _, err := dbHandler.DB.Exec(context.Background(), EDeleteAllTestResults); err != nil {
		fmt.Printf("Error deleting test_results table: %+v\n", err)
		return err
	}
	if _, err := dbHandler.DB.Exec(context.Background(), EDeleteAllUserInfo); err != nil {
		fmt.Printf("Error dropping user_info table: %+v\n", err)
		return err
	}
	if _, err := dbHandler.DB.Exec(context.Background(), EDeleteRole); err != nil {
		fmt.Printf("Error dropping user_info table: %+v\n", err)
		return err
	}
	return nil
}
