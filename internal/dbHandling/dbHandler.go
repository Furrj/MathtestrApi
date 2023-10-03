package dbHandling

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"mathtestr.com/server/internal/schemas"
)

type DBHandler struct {
	DB *pgx.Conn
}

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

func (dbHandler *DBHandler) CheckIfUsernameExists(username string) (bool, error) {
	var returnedUsername string
	err := dbHandler.DB.QueryRow(context.Background(), QCheckIfUsernameExists, username).Scan(&returnedUsername)
	if err == pgx.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dbHandler *DBHandler) GetUserByUsername(username string) (schemas.AllUserData, error) {
	var user schemas.AllUserData
	err := dbHandler.DB.QueryRow(context.Background(), QGetUserByUsername, username).Scan(&user.ID, &user.Username, &user.Password, &user.Firstname, &user.Lastname, &user.UUID, &user.Expires)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (dbHandler *DBHandler) GetUserIDByUsername(username string) (int, error) {
	var id int
	err := dbHandler.DB.QueryRow(context.Background(), QGetUserIDByUsername, username).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (dbHandler *DBHandler) InsertUser(r schemas.RegisterPayload) error {
	_, err := dbHandler.DB.Exec(context.Background(), EInsertUser, r.Username, r.Password, r.FirstName, r.LastName)
	if err != nil {
		return err
	}
	return nil
}

func (dbHandler *DBHandler) InsertSessionData(s schemas.SessionData) error {
	_, err := dbHandler.DB.Exec(context.Background(), EInsertSessionData, s.ID, s.UUID, s.Expires)
	if err != nil {
		log.Printf("Error inserting session data: %+v\n", err)
		return err
	}
	return nil
}

func (dbHandler *DBHandler) DropTables() error {
	if _, err := dbHandler.DB.Exec(context.Background(), EDeleteAllSessionData); err != nil {
		fmt.Printf("Error deleting session_data info: %+v\n", err)
		return err
	}
	if _, err := dbHandler.DB.Exec(context.Background(), EDeleteAllUserInfo); err != nil {
		fmt.Printf("Error deleting user_info info: %+v\n", err)
		return err
	}
	return nil
}

func (dbHandler *DBHandler) CreateTables() error {
	if _, err := dbHandler.DB.Exec(context.Background(), EInitUserInfo); err != nil {
		fmt.Printf("Error initializing user_info: %+v\n", err)
	}
	if _, err := dbHandler.DB.Exec(context.Background(), EInitSessionData); err != nil {
		fmt.Printf("Error initializing session_data: %+v\n", err)
	}
	return nil
}
