package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	dbHandlers "mathtestr.com/server/internal/dbHandling"
	"mathtestr.com/server/internal/routeHandling"
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

	routeHandler := routeHandling.InitRouteHandler(dbHandler)
	router := gin.Default()
	router.POST("/register", routeHandler.Register)
	router.Run(":5000")

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
