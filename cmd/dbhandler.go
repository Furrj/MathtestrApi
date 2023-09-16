package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func OpenDBConnection() *pgx.Conn {
	connection_string := "postgres://postgres:password@localhost:5432/testdb"
	db, err := pgx.Connect(context.Background(), connection_string)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	return db
}

func FindByUsername(db *pgx.Conn, username string) (User, error) {
	var user User
	err := db.QueryRow(context.Background(), "SELECT * FROM account_info WHERE username=$1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
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

func AddUser(db *pgx.Conn, user User) error {
	_, err := db.Exec(context.Background(), "INSERT INTO account_info (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	return nil
}

func GetAllUsers(db *pgx.Conn) ([]User, error) {
	var userList []User

	rows, err := db.Query(context.Background(), "select * from account_info")

	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Username, &user.Password)
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
