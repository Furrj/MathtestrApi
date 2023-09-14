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

	userList, _ := GetAllUsers(db)
	for _, user := range userList {
		fmt.Printf("%+v\n", user)
	}

	// Routing
	router := gin.Default()

	router.POST("/submit", UserLogin(db))
	router.Use(spa.Middleware("/", "../build"))

	router.Run(":5000")
}
