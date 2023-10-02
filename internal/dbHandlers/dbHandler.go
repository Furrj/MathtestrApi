package dbHandlers

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type dbHandler struct {
	db *pgx.Conn
}

func InitDBHandler() *dbHandler {
	var newDBHandler dbHandler
	connection_string := "postgres://postgres:password@localhost:5432/mathtestr"
	db, err := pgx.Connect(context.Background(), connection_string)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	newDBHandler.db = db
	return &newDBHandler
}

func (dbHandler *dbHandler) CheckIfUsernameExists(username string) (bool, error) {
	var returnedUsername string
	err := dbHandler.db.QueryRow(context.Background(), "SELECT username FROM user_info WHERE username=$1", username).Scan(&returnedUsername)
	if err == pgx.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
