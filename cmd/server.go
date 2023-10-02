package main

import (
	"context"
	"fmt"

	"mathtestr.com/server/internal/dbHandlers"
)

func main() {
	// DB
	dbHandler := dbHandlers.InitDBHandler()
	defer dbHandler.DB.Close(context.Background())

	user, err := dbHandler.GetUserByUsername("a")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", user)

	// userList, err := dbHandlers.GetAllUsers(db)
	// if err != nil {
	// 	fmt.Printf("%+v\n", err)
	// }
	// for _, user := range userList {
	// 	fmt.Printf("%+v\n", user)
	// }

	// // Routing
	// router := gin.Default()

	// router.POST("/login", routeHandlers.LoginPost(db))
	// router.POST("/register", routeHandlers.RegisterPost(db))
	// router.Use(spa.Middleware("/", "client"))

	// router.Run(":5000")
}
