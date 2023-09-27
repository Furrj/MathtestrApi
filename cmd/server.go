package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
	"mathtestr.com/server/internal/dbHandlers"
	"mathtestr.com/server/internal/routeHandlers"
)

func main() {
	// DB
	db := dbHandlers.OpenDBConnection()
	defer db.Close(context.Background())

	userList, err := dbHandlers.GetAllUsers(db)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	for _, user := range userList {
		fmt.Printf("%+v\n", user)
	}

	// Routing
	router := gin.Default()

	router.POST("/login", routeHandlers.LoginPost(db))
	router.POST("/register", routeHandlers.RegisterPost(db))
	router.Use(spa.Middleware("/", "../build"))

	router.Run(":5000")
}
