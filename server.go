package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type User struct {
	ID       int64  `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *pgx.Conn

func main() {
	// DB
	connection_string := "postgres://postgres:password@localhost:5432/testdb"
	var err error
	db, err = pgx.Connect(context.Background(), connection_string)
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close(context.Background())
	// findByUsername(db, "test")
	// addUser(db)
	getData(db)

	// Routing
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./build", true)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./build/index.html")
	})
	router.POST("/submit", getPost)

	router.Run(":5000")
}

func getPost(c *gin.Context) {
	var info User
	if err := c.BindJSON(&info); err != nil {
		fmt.Println("Error")
	}
	fmt.Printf("%+v\n", info)
	addUser(db, info)
}

func getData(db *pgx.Conn) error {
	var users []User
	var user User

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	rows, err := db.Query(context.Background(), "select * from account_info")

	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password)
		users = append(users, user)
		if err != nil {
			return err
		}
	}

	if rows.Err() != nil {
		return err
	}

	fmt.Printf("%+v\n", users)

	return nil
}

func addUser(db *pgx.Conn, user User) error {
	_, err := db.Exec(context.Background(), "INSERT INTO account_info (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	return nil
}

func findByUsername(db *pgx.Conn, name string) error {
	var user User
	err := db.QueryRow(context.Background(), "SELECT * FROM account_info WHERE username=$1", name).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", user)
	return nil
}

func deleteByUsername(db *pgx.Conn, name string) error {
	_, err := db.Exec(context.Background(), "DELETE FROM account_info WHERE username=$1", name)
	if err != nil {
		return err
	}
	fmt.Println("Deleted")
	return nil
}
