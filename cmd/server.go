package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mandrigin/gin-spa/spa"
	"mathtestr.com/server/internal/dbHandling"
	"mathtestr.com/server/internal/logger"
	"mathtestr.com/server/internal/routeHandling"
)

func main() {
	// ENV CONFIG
	if os.Getenv("MODE") != "PROD" {
		godotenv.Load("config.env")
	}
	logger.WriteLog("Hellos")

	// DB
	dbHandler := dbHandling.InitDBHandler(os.Getenv("DB_URL"))
	defer dbHandler.DB.Close(context.Background())

	user, err := dbHandler.GetUserByUsername("a")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", user)

	// ROUTING
	routeHandler := routeHandling.InitRouteHandler(dbHandler)
	router := gin.Default()
	router.POST("/register", routeHandler.Register)
	router.Use(spa.Middleware("/", "client"))
	router.Run(":5000")
}
