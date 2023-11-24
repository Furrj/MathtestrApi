// dbHandler package handles communication with the Postgres database
package dbHandler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"mathtestr.com/server/internal/schemas"
)

// DBHandler object contains a pointer to pgx connection, one per server
type DBHandler struct {
	DB *pgx.Conn
}

// InitDBHandler is constructor for DBHandler, takes a database connection
// string and returns a pointer to instantiated DBHandler.
func InitDBHandler(connectionString string) *DBHandler {
	var newDBHandler DBHandler
	db, err := pgx.Connect(context.Background(), connectionString)
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

// GetUserByUsername takes in username string and searches database
// for it. Binds query to AllUserData schema, then returns it and error
func (dbHandler *DBHandler) GetUserByUsername(username string) (schemas.AllUserData, error) {
	var user schemas.AllUserData
	err := dbHandler.DB.QueryRow(context.Background(), QGetUserByUsername, username).Scan(&user.ID, &user.Username, &user.Password, &user.Firstname, &user.Lastname, &user.Role, &user.Period, &user.Teacher, &user.SessionKey, &user.Expires)
	if err != nil {
		return user, err
	}
	return user, nil
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

// InsertUserInfo takes RegisterPayload and inserts user data into user_info
// table, returns error
// FIXME: Handle Role, as of now hardcoded
func (dbHandler *DBHandler) InsertUserInfo(r schemas.RegisterPayload) error {
	// Hardcoded user Role
	const DefaultRole = "S"
	_, err := dbHandler.DB.Exec(context.Background(), EInsertUserInfo, r.Username, r.Password, r.FirstName, r.LastName, DefaultRole, r.Period, r.Teacher)
	if err != nil {
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
	if _, err := dbHandler.DB.Exec(context.Background(), EInitTestResults); err != nil {
		fmt.Printf("Error initializing session_data: %+v\n", err)
	}
	return nil
}

// DropTables is used for testing, drops all tables from database, returns error
func (dbHandler *DBHandler) DropTables() error {
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
