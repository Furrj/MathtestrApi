package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
)

func main() {
	// DB
	db := OpenDBConnection()
	defer db.Close(context.Background())

	userList, err := GetAllUsers(db)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	for _, user := range userList {
		fmt.Printf("%+v\n", user)
	}

	// Routing
	router := gin.Default()

	router.POST("/login", LoginPost(db))
	router.POST("/register", RegisterPost(db))
	router.Use(spa.Middleware("/", "../build"))

	router.Run(":5000")
}
