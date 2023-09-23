package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func OpenDBConnection() *pgx.Conn {
	connection_string := "postgres://postgres:password@localhost:5432/mathtestr"
	db, err := pgx.Connect(context.Background(), connection_string)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	return db
}

func FindByUsername(db *pgx.Conn, username string) (UserClientData, error) {
	var user UserClientData
	err := db.QueryRow(context.Background(), "SELECT * FROM user_info WHERE username=$1", username).Scan(&user.ID, &user.Username)
	if err != nil {
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

func GetAllUsers(db *pgx.Conn) ([]UserClientData, error) {
	var userList []UserClientData

	rows, err := db.Query(context.Background(), "SELECT user_id, username, uuid FROM user_info NATURAL JOIN session_data")

	for rows.Next() {
		var user UserClientData
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

func InsertUser(db *pgx.Conn, userInfo RegisterPayload) error {
	_, err := db.Exec(context.Background(), "INSERT INTO user_info (username, password, first_name, last_name) VALUES ($1, $2, $3, $4)", userInfo.Username, userInfo.Password, userInfo.FirstName, userInfo.LastName)
	if err != nil {
		log.Printf("Error inserting user: %+v\n", err)
		return err
	}
	return nil
}
