package dbHandlers

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"mathtestr.com/server/internal/types"
)

func FindByUsername(db *pgx.Conn, username string) (types.UserClientData, error) {
	var user types.UserClientData
	err := db.QueryRow(context.Background(), "SELECT user_id, username FROM user_info WHERE username=$1", username).Scan(&user.ID, &user.Username)
	if err != nil {
		log.Printf("%+v\n", err)
		if err == pgx.ErrNoRows {
			user.ID = -1
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func DeleteByUsername(db *pgx.Conn, name string) error {
	_, err := db.Exec(context.Background(), "DELETE FROM account_info WHERE username=$1", name)
	if err != nil {
		return err
	}
	fmt.Println("Deleted")
	return nil
}

func GetAllUsers(db *pgx.Conn) ([]types.UserClientData, error) {
	var userList []types.UserClientData

	rows, err := db.Query(context.Background(), "SELECT user_id, username, uuid FROM user_info NATURAL JOIN session_data")

	for rows.Next() {
		var user types.UserClientData
		rows.Scan(&user.ID, &user.Username, &user.UUID)
		userList = append(userList, user)
		if err != nil {
			return userList, err
		}
	}

	if rows.Err() != nil {
		return userList, rows.Err()
	}

	return userList, nil
}

func InsertUser(db *pgx.Conn, userInfo types.RegisterPayload) error {
	_, err := db.Exec(context.Background(), "INSERT INTO user_info (username, password, first_name, last_name) VALUES ($1, $2, $3, $4)", userInfo.Username, userInfo.Password, userInfo.FirstName, userInfo.LastName)
	if err != nil {
		log.Printf("Error inserting user: %+v\n", err)
		return err
	}

	return nil
}

func InsertSessionData(db *pgx.Conn, sessionData types.SessionData) error {
	_, err := db.Exec(context.Background(), "Insert INTO session_data (user_id, uuid, expires) VALUES ($1, $2, $3)", sessionData.ID, sessionData.UUID, sessionData.Expires)
	log.Printf("Error inserting session data: %+v\n", err)
	return err
}
