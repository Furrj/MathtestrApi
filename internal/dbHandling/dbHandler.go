package dbHandling

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"mathtestr.com/server/internal/schemas"
)

type DBHandler struct {
	DB *pgx.Conn
}

func InitDBHandler() *DBHandler {
	var newDBHandler DBHandler
	//connection_string := "postgres://postgres:password@localhost:5432/mathtestr"
	connection_string := "postgres://postgres:password@host.docker.internal:5432/mathtestr"
	db, err := pgx.Connect(context.Background(), connection_string)
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
