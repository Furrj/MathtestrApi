package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mandrigin/gin-spa/spa"
	"mathtestr.com/server/internal/dbHandler"
	"mathtestr.com/server/internal/routeHandling"
)

func main() {
	// ENV CONFIG
	if os.Getenv("MODE") != "PROD" {
		godotenv.Load("config.env")
	}

	// Test backup
	// cmd := exec.Command("./backup.sh")
	// if err := cmd.Run(); err != nil {
	// 	log.Printf("Error backing up Postgres: %+v\n", err)
	// }

	// DB
	dbHandler := dbHandler.InitDBHandler(os.Getenv("DB_URL"))
	defer dbHandler.DB.Close(context.Background())

	// ROUTING
	routeHandler := routeHandling.InitRouteHandler(dbHandler)
	router := gin.Default()
	router.POST("/api/register", routeHandler.Register)
	router.POST("/api/validateSession", routeHandler.ValidateSession)
	router.POST("/api/login", routeHandler.Login)
	router.Use(spa.Middleware("/", "client"))
	log.Panic(router.Run(":5000"))
}
